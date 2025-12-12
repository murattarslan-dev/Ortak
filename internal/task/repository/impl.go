package repository

import (
	"database/sql"
	"ortak/internal/task"
	"strings"
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
	rows, err := r.db.Query("SELECT id, title, description, status, assignee_id, team_id, tags FROM tasks")
	if err != nil {
		return []task.Task{}
	}
	defer rows.Close()

	var tasks []task.Task
	for rows.Next() {
		var t task.Task
		var tagsStr string
		rows.Scan(&t.ID, &t.Title, &t.Description, &t.Status, &t.AssigneeID, &t.TeamID, &tagsStr)
		t.Tags = r.stringToTags(tagsStr)

		var commentCount int
		r.db.QueryRow("SELECT COUNT(*) FROM task_comments WHERE task_id = ?", t.ID).Scan(&commentCount)
		t.CommentCount = commentCount

		tasks = append(tasks, t)
	}
	return tasks
}

func (r *RepositoryImpl) stringToTags(tagStr string) []string {
	if tagStr == "" {
		return []string{}
	}
	return strings.Split(tagStr, "{$^#}")
}

func (r *RepositoryImpl) tagsToString(tags []string) string {
	return strings.Join(tags, "{$^#}")
}

func (r *RepositoryImpl) Create(title, description string, assigneeID, teamID int, tags []string) *task.Task {
	tagsStr := r.tagsToString(tags)
	result, err := r.db.Exec("INSERT INTO tasks (title, description, status, assignee_id, team_id, tags) VALUES (?, ?, 'todo', ?, ?, ?)",
		title, description, assigneeID, teamID, tagsStr)
	if err != nil {
		return nil
	}
	id, _ := result.LastInsertId()
	return &task.Task{
		ID: int(id), Title: title, Description: description, Status: "todo",
		AssigneeID: assigneeID, TeamID: teamID, Tags: tags, CommentCount: 0,
	}
}

func (r *RepositoryImpl) GetByID(id string) *task.Task {
	var t task.Task
	var tagsStr string
	err := r.db.QueryRow("SELECT id, title, description, status, assignee_id, team_id, tags FROM tasks WHERE id = ?", id).
		Scan(&t.ID, &t.Title, &t.Description, &t.Status, &t.AssigneeID, &t.TeamID, &tagsStr)
	if err != nil {
		return nil
	}
	t.Tags = r.stringToTags(tagsStr)

	var commentCount int
	r.db.QueryRow("SELECT COUNT(*) FROM task_comments WHERE task_id = ?", t.ID).Scan(&commentCount)
	t.CommentCount = commentCount

	return &t
}

func (r *RepositoryImpl) GetByIDWithComments(id string) *task.Task {
	t := r.GetByID(id)
	if t == nil {
		return nil
	}

	// Get comments
	rows, _ := r.db.Query(`SELECT tc.id, tc.task_id, tc.comment, tc.created_at, u.id, u.username, u.email 
		FROM task_comments tc JOIN users u ON tc.user_id = u.id WHERE tc.task_id = ?`, t.ID)
	defer rows.Close()

	var comments []task.TaskComment
	for rows.Next() {
		var c task.TaskComment
		rows.Scan(&c.ID, &c.TaskID, &c.Comment, &c.CreatedAt, &c.User.ID, &c.User.Username, &c.User.Email)
		comments = append(comments, c)
	}

	// Get assignments
	rows2, _ := r.db.Query("SELECT id, task_id, assign_type, assign_id, created_at FROM task_assignments WHERE task_id = ?", t.ID)
	defer rows2.Close()

	var assignments []task.TaskAssignment
	for rows2.Next() {
		var a task.TaskAssignment
		rows2.Scan(&a.ID, &a.TaskID, &a.AssignType, &a.AssignID, &a.CreatedAt)

		if a.AssignType == "user" {
			var user task.CommentUser
			r.db.QueryRow("SELECT id, username, email FROM users WHERE id = ?", a.AssignID).
				Scan(&user.ID, &user.Username, &user.Email)
			a.User = &user
		} else if a.AssignType == "team" {
			var team task.AssignTeam
			r.db.QueryRow("SELECT id, name FROM teams WHERE id = ?", a.AssignID).
				Scan(&team.ID, &team.Name)
			a.Team = &team
		}
		assignments = append(assignments, a)
	}

	t.Comments = comments
	t.Assignments = assignments
	t.CommentCount = len(comments)
	return t
}

func (r *RepositoryImpl) Update(id, title, description, status string, assigneeID int, tags []string) *task.Task {
	tagsStr := r.tagsToString(tags)
	_, err := r.db.Exec("UPDATE tasks SET title = ?, description = ?, status = ?, assignee_id = ?, tags = ? WHERE id = ?",
		title, description, status, assigneeID, tagsStr, id)
	if err != nil {
		return nil
	}
	return r.GetByID(id)
}

func (r *RepositoryImpl) UpdateStatus(id, status string) *task.Task {
	_, err := r.db.Exec("UPDATE tasks SET status = ? WHERE id = ?", status, id)
	if err != nil {
		return nil
	}
	return r.GetByID(id)
}

func (r *RepositoryImpl) AddComment(taskID, userID int, comment, createdAt string) *task.TaskComment {
	result, err := r.db.Exec("INSERT INTO task_comments (task_id, user_id, comment, created_at) VALUES (?, ?, ?, ?)",
		taskID, userID, comment, createdAt)
	if err != nil {
		return nil
	}
	id, _ := result.LastInsertId()

	var user task.CommentUser
	r.db.QueryRow("SELECT id, username, email FROM users WHERE id = ?", userID).
		Scan(&user.ID, &user.Username, &user.Email)

	return &task.TaskComment{
		ID: int(id), TaskID: taskID, Comment: comment, CreatedAt: createdAt, User: user,
	}
}

func (r *RepositoryImpl) AddAssignment(taskID int, assignType string, assignID int, createdAt string) *task.TaskAssignment {
	result, err := r.db.Exec("INSERT INTO task_assignments (task_id, assign_type, assign_id, created_at) VALUES (?, ?, ?, ?)",
		taskID, assignType, assignID, createdAt)
	if err != nil {
		return nil
	}
	id, _ := result.LastInsertId()

	assignment := &task.TaskAssignment{
		ID: int(id), TaskID: taskID, AssignType: assignType, AssignID: assignID, CreatedAt: createdAt,
	}

	if assignType == "user" {
		var user task.CommentUser
		r.db.QueryRow("SELECT id, username, email FROM users WHERE id = ?", assignID).
			Scan(&user.ID, &user.Username, &user.Email)
		assignment.User = &user
	} else if assignType == "team" {
		var team task.AssignTeam
		r.db.QueryRow("SELECT id, name FROM teams WHERE id = ?", assignID).
			Scan(&team.ID, &team.Name)
		assignment.Team = &team
	}

	return assignment
}

func (r *RepositoryImpl) DeleteAssignment(assignmentID int) error {
	_, err := r.db.Exec("DELETE FROM task_assignments WHERE id = ?", assignmentID)
	return err
}

func (r *RepositoryImpl) Delete(id string) error {
	_, err := r.db.Exec("DELETE FROM tasks WHERE id = ?", id)
	return err
}
