package task

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	AssigneeID  int    `json:"assignee_id"`
	TeamID      int    `json:"team_id"`
}

type CreateTaskRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	AssigneeID  int    `json:"assignee_id"`
	TeamID      int    `json:"team_id" binding:"required"`
}