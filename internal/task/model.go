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

type UpdateTaskRequest struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Status      string `json:"status,omitempty"`
	AssigneeID  int    `json:"assignee_id,omitempty"`
}

type UpdateTaskStatusRequest struct {
	Status string `json:"status" binding:"required"`
}