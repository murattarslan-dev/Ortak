package task

type Task struct {
	ID           int           `json:"id"`
	Title        string        `json:"title"`
	Description  string        `json:"description"`
	Status       string        `json:"status"`
	AssigneeID   int           `json:"assignee_id"`
	TeamID       int           `json:"team_id"`
	Tags         []string      `json:"tags"`
	CommentCount int           `json:"comment_count,omitempty"`
	Comments     []TaskComment `json:"comments,omitempty"`
}

type TaskComment struct {
	ID        int        `json:"id"`
	TaskID    int        `json:"task_id"`
	Comment   string     `json:"comment"`
	CreatedAt string     `json:"created_at"`
	User      CommentUser `json:"user"`
}

type CommentUser struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type CreateTaskRequest struct {
	Title       string   `json:"title" binding:"required"`
	Description string   `json:"description"`
	AssigneeID  int      `json:"assignee_id"`
	TeamID      int      `json:"team_id" binding:"required"`
	Tags        []string `json:"tags"`
}

type UpdateTaskRequest struct {
	Title       string   `json:"title,omitempty"`
	Description string   `json:"description,omitempty"`
	Status      string   `json:"status,omitempty"`
	AssigneeID  int      `json:"assignee_id,omitempty"`
	Tags        []string `json:"tags,omitempty"`
}

type UpdateTaskStatusRequest struct {
	Status string `json:"status" binding:"required"`
}

type AddCommentRequest struct {
	Comment string `json:"comment" binding:"required"`
}