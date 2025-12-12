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

func (r *RepositoryImpl) GetAll() []task.Task {
	rows, err := r.db.Query("SELECT id, title, description, status, tags, priority, due_date, created_at, updated_at, created_by FROM tasks")
	if err != nil {
		return []task.Task{}
	}
	defer rows.Close()

	var tasks []task.Task
	for rows.Next() {
		var t task.Task
		var tagsJSON []byte
		rows.Scan(&t.ID, &t.Title, &t.Description, &t.Status, &tagsJSON, &t.Priority, &t.DueDate, &t.CreatedAt, &t.UpdatedAt, &t.CreatedBy)

		if tagsJSON != nil {
			json.Unmarshal(tagsJSON, &t.Tags)
		}

		var commentCount int
		r.db.QueryRow("SELECT COUNT(*) FROM task_comments WHERE task_id = $1", t.ID).Scan(&commentCount)
		t.CommentCount = commentCount

		tasks = append(tasks, t)
	}
	return tasks
}

func (r *RepositoryImpl) Create(title, description, createdBy string, tags []string, priority string, dueDate *time.Time) *task.Task {
	tagsJSON, _ := json.Marshal(tags)

	var id string
	err := r.db.QueryRow("INSERT INTO tasks (title, description, status, tags, priority, due_date, created_by) VALUES ($1, $2, 'todo', $3, $4, $5, $6) RETURNING id",
		title, description, tagsJSON, priority, dueDate, createdBy).Scan(&id)
	if err != nil {
		return nil
	}

	return r.GetByID(id)
}

func (r *RepositoryImpl) GetByID(id string) *task.Task {
	var t task.Task
	var tagsJSON []byte
	err := r.db.QueryRow("SELECT id, title, description, status, tags, priority, due_date, created_at, updated_at, created_by FROM tasks WHERE id = $1", id).
		Scan(&t.ID, &t.Title, &t.Description, &t.Status, &tagsJSON, &t.Priority, &t.DueDate, &t.CreatedAt, &t.UpdatedAt, &t.CreatedBy)
	if err != nil {
		return nil
	}

	if tagsJSON != nil {
		json.Unmarshal(tagsJSON, &t.Tags)
	}

	var commentCount int
	r.db.QueryRow("SELECT COUNT(*) FROM task_comments WHERE task_id = $1", t.ID).Scan(&commentCount)
	t.CommentCount = commentCount

	return &t
}

func (r *RepositoryImpl) GetByIDWithComments(id string) *task.Task {
	t := r.GetByID(id)
	if t == nil {
		return nil
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
	rows2, _ := r.db.Query("SELECT id, task_id, assign_type, assign_id, created_at FROM task_assignments WHERE task_id = $1", t.ID)
	defer rows2.Close()

	var assignments []task.TaskAssignment
	for rows2.Next() {
		var a task.TaskAssignment
		rows2.Scan(&a.ID, &a.TaskID, &a.AssignType, &a.AssignID, &a.CreatedAt)

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
	return t
}

func (r *RepositoryImpl) Update(id, title, description, status string, tags []string, priority string, dueDate *time.Time) *task.Task {
	tagsJSON, _ := json.Marshal(tags)
	_, err := r.db.Exec("UPDATE tasks SET title = $1, description = $2, status = $3, tags = $4, priority = $5, due_date = $6, updated_at = CURRENT_TIMESTAMP WHERE id = $7",
		title, description, status, tagsJSON, priority, dueDate, id)
	if err != nil {
		return nil
	}
	return r.GetByID(id)
}

func (r *RepositoryImpl) UpdateStatus(id, status string) *task.Task {
	_, err := r.db.Exec("UPDATE tasks SET status = $1, updated_at = CURRENT_TIMESTAMP WHERE id = $2", status, id)
	if err != nil {
		return nil
	}
	return r.GetByID(id)
}

func (r *RepositoryImpl) AddComment(taskID, userID, comment string) *task.TaskComment {
	var id string
	err := r.db.QueryRow("INSERT INTO task_comments (task_id, user_id, comment) VALUES ($1, $2, $3) RETURNING id",
		taskID, userID, comment).Scan(&id)
	if err != nil {
		return nil
	}

	u := &user.User{}
	r.db.QueryRow("SELECT id, username, email, first_name, last_name FROM users WHERE id = $1", userID).
		Scan(&u.ID, &u.Username, &u.Email, &u.FirstName, &u.LastName)

	return &task.TaskComment{
		ID: id, TaskID: taskID, UserID: userID, Comment: comment,
		CreatedAt: time.Now(), User: u,
	}
}

func (r *RepositoryImpl) AddAssignment(taskID, assignType, assignID string) *task.TaskAssignment {
	var id string
	err := r.db.QueryRow("INSERT INTO task_assignments (task_id, assign_type, assign_id) VALUES ($1, $2, $3) RETURNING id",
		taskID, assignType, assignID).Scan(&id)
	if err != nil {
		return nil
	}

	assignment := &task.TaskAssignment{
		ID: id, TaskID: taskID, AssignType: assignType, AssignID: assignID, CreatedAt: time.Now(),
	}

	if assignType == "user" {
		assignment.User = &user.User{}
		r.db.QueryRow("SELECT id, username, email, first_name, last_name FROM users WHERE id = $1", assignID).
			Scan(&assignment.User.ID, &assignment.User.Username, &assignment.User.Email, &assignment.User.FirstName, &assignment.User.LastName)
	} else if assignType == "team" {
		assignment.Team = &team.Team{}
		r.db.QueryRow("SELECT id, name, description FROM teams WHERE id = $1", assignID).
			Scan(&assignment.Team.ID, &assignment.Team.Name, &assignment.Team.Description)
	}

	return assignment
}

func (r *RepositoryImpl) DeleteAssignment(assignmentID string) error {
	_, err := r.db.Exec("DELETE FROM task_assignments WHERE id = $1", assignmentID)
	return err
}

func (r *RepositoryImpl) Delete(id string) error {
	_, err := r.db.Exec("DELETE FROM tasks WHERE id = $1", id)
	return err
}
