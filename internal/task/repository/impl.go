package repository

import (
	"database/sql"
	"encoding/json"
	"ortak/internal/task"
	"ortak/internal/team"
	"ortak/internal/user"
	"time"
)

type RepositoryImpl struct {
	db *sql.DB
}

func NewRepositoryImpl(database *sql.DB) Repository {
	return &RepositoryImpl{
		db: database,
	}
}

func (r *RepositoryImpl) GetAll() ([]task.Task, error) {
	rows, err := r.db.Query("SELECT id, title, description, status, tags, priority, due_date, created_at, updated_at, created_by FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []task.Task
	for rows.Next() {
		var t task.Task
		var tagsJSON []byte
		err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.Status, &tagsJSON, &t.Priority, &t.DueDate, &t.CreatedAt, &t.UpdatedAt, &t.CreatedBy)
		if err != nil {
			return nil, err
		}

		if tagsJSON != nil {
			json.Unmarshal(tagsJSON, &t.Tags)
		}

		var commentCount int
		r.db.QueryRow("SELECT COUNT(*) FROM task_comments WHERE task_id = $1", t.ID).Scan(&commentCount)
		t.CommentCount = commentCount

		tasks = append(tasks, t)
	}
	return tasks, nil
}

func (r *RepositoryImpl) Create(title, description, createdBy string, tags []string, priority string, dueDate *time.Time) (*task.Task, error) {
	tagsJSON, _ := json.Marshal(tags)

	var id string
	err := r.db.QueryRow("INSERT INTO tasks (title, description, status, tags, priority, due_date, created_by) VALUES ($1, $2, 'todo', $3, $4, $5, $6) RETURNING id",
		title, description, tagsJSON, priority, dueDate, createdBy).Scan(&id)
	if err != nil {
		return nil, err
	}

	return r.GetByID(id)
}

func (r *RepositoryImpl) GetByID(id string) (*task.Task, error) {
	var t task.Task
	var tagsJSON []byte
	err := r.db.QueryRow("SELECT id, title, description, status, tags, priority, due_date, created_at, updated_at, created_by FROM tasks WHERE id = $1", id).
		Scan(&t.ID, &t.Title, &t.Description, &t.Status, &tagsJSON, &t.Priority, &t.DueDate, &t.CreatedAt, &t.UpdatedAt, &t.CreatedBy)
	if err != nil {
		return nil, err
	}

	if tagsJSON != nil {
		json.Unmarshal(tagsJSON, &t.Tags)
	}

	var commentCount int
	r.db.QueryRow("SELECT COUNT(*) FROM task_comments WHERE task_id = $1", t.ID).Scan(&commentCount)
	t.CommentCount = commentCount

	return &t, nil
}

func (r *RepositoryImpl) GetByIDWithComments(id string) (*task.Task, error) {
	t, err := r.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Get comments
	rows, _ := r.db.Query(`SELECT tc.id, tc.task_id, tc.user_id, tc.comment, tc.created_at, 
		u.id, u.username, u.email, u.first_name, u.last_name
		FROM task_comments tc JOIN users u ON tc.user_id = u.id WHERE tc.task_id = $1`, t.ID)
	defer rows.Close()

	var comments []task.TaskComment
	for rows.Next() {
		var c task.TaskComment
		c.User = &user.User{}
		rows.Scan(&c.ID, &c.TaskID, &c.UserID, &c.Comment, &c.CreatedAt,
			&c.User.ID, &c.User.Username, &c.User.Email, &c.User.FirstName, &c.User.LastName)
		comments = append(comments, c)
	}

	// Get assignments
	rows2, _ := r.db.Query("SELECT task_id, assign_type, assign_id, created_at FROM task_assignments WHERE task_id = $1", t.ID)
	defer rows2.Close()

	var assignments []task.TaskAssignment
	for rows2.Next() {
		var a task.TaskAssignment
		rows2.Scan(&a.TaskID, &a.AssignType, &a.AssignID, &a.CreatedAt)

		if a.AssignType == "user" {
			a.User = &user.User{}
			r.db.QueryRow("SELECT id, username, email, first_name, last_name FROM users WHERE id = $1", a.AssignID).
				Scan(&a.User.ID, &a.User.Username, &a.User.Email, &a.User.FirstName, &a.User.LastName)
		} else if a.AssignType == "team" {
			a.Team = &team.Team{}
			r.db.QueryRow("SELECT id, name, description FROM teams WHERE id = $1", a.AssignID).
				Scan(&a.Team.ID, &a.Team.Name, &a.Team.Description)
		}
		assignments = append(assignments, a)
	}

	t.Comments = comments
	t.Assignments = assignments
	t.CommentCount = len(comments)
	return t, nil
}

func (r *RepositoryImpl) Update(id, title, description, status string, tags []string, priority string, dueDate *time.Time) (*task.Task, error) {
	tagsJSON, _ := json.Marshal(tags)
	_, err := r.db.Exec("UPDATE tasks SET title = $1, description = $2, status = $3, tags = $4, priority = $5, due_date = $6, updated_at = CURRENT_TIMESTAMP WHERE id = $7",
		title, description, status, tagsJSON, priority, dueDate, id)
	if err != nil {
		return nil, err
	}
	return r.GetByID(id)
}

func (r *RepositoryImpl) UpdateStatus(id, status string) (*task.Task, error) {
	_, err := r.db.Exec("UPDATE tasks SET status = $1, updated_at = CURRENT_TIMESTAMP WHERE id = $2", status, id)
	if err != nil {
		return nil, err
	}
	return r.GetByID(id)
}

func (r *RepositoryImpl) AddComment(taskID, userID, comment string) (*task.TaskComment, error) {
	var id string
	err := r.db.QueryRow("INSERT INTO task_comments (task_id, user_id, comment) VALUES ($1, $2, $3) RETURNING id",
		taskID, userID, comment).Scan(&id)
	if err != nil {
		return nil, err
	}

	u := &user.User{}
	err = r.db.QueryRow("SELECT id, username, email, first_name, last_name FROM users WHERE id = $1", userID).
		Scan(&u.ID, &u.Username, &u.Email, &u.FirstName, &u.LastName)
	if err != nil {
		return nil, err
	}

	return &task.TaskComment{
		ID: id, TaskID: taskID, UserID: userID, Comment: comment,
		CreatedAt: time.Now(), User: u,
	}, nil
}

func (r *RepositoryImpl) AddAssignment(taskID, assignType, assignID string) (*task.TaskAssignment, error) {
	_, err := r.db.Exec("INSERT INTO task_assignments (task_id, assign_type, assign_id) VALUES ($1, $2, $3)",
		taskID, assignType, assignID)
	if err != nil {
		return nil, err
	}

	assignment := &task.TaskAssignment{
		TaskID: taskID, AssignType: assignType, AssignID: assignID, CreatedAt: time.Now(),
	}

	switch assignType {
	case "user":
		assignment.User = &user.User{}
		err = r.db.QueryRow("SELECT id, username, email, first_name, last_name FROM users WHERE id = $1", assignID).
			Scan(&assignment.User.ID, &assignment.User.Username, &assignment.User.Email, &assignment.User.FirstName, &assignment.User.LastName)
		if err != nil {
			return nil, err
		}
	case "team":
		assignment.Team = &team.Team{}
		err = r.db.QueryRow("SELECT id, name, description FROM teams WHERE id = $1", assignID).
			Scan(&assignment.Team.ID, &assignment.Team.Name, &assignment.Team.Description)
		if err != nil {
			return nil, err
		}
	}

	return assignment, nil
}

func (r *RepositoryImpl) DeleteAssignment(taskID, assignType, assignID string) error {
	_, err := r.db.Exec("DELETE FROM task_assignments WHERE task_id = $1 AND assign_type = $2 AND assign_id = $3", taskID, assignType, assignID)
	return err
}

func (r *RepositoryImpl) Delete(id string) error {
	_, err := r.db.Exec("DELETE FROM tasks WHERE id = $1", id)
	return err
}
