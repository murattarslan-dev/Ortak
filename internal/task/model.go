package task

import (
	"ortak/internal/team"
	"ortak/internal/user"
	"time"
)

type Task struct {
	ID           string           `json:"id"`
	Title        string           `json:"title"`
	Description  string           `json:"description"`
	Status       string           `json:"status"`
	Tags         []string         `json:"tags"`
	Priority     string           `json:"priority"`
	DueDate      *time.Time       `json:"due_date,omitempty"`
	CreatedAt    time.Time        `json:"created_at"`
	UpdatedAt    time.Time        `json:"updated_at"`
	CreatedBy    string           `json:"created_by"`
	CommentCount int              `json:"comment_count,omitempty"`
	Comments     []TaskComment    `json:"comments,omitempty"`
	Assignments  []TaskAssignment `json:"assignments,omitempty"`
}

type TaskComment struct {
	ID        string     `json:"id"`
	TaskID    string     `json:"task_id"`
	UserID    string     `json:"user_id"`
	Comment   string     `json:"comment"`
	CreatedAt time.Time  `json:"created_at"`
	User      *user.User `json:"user,omitempty"`
}

type TaskAssignment struct {
	TaskID     string     `json:"task_id"`
	AssignType string     `json:"assign_type"`
	AssignID   string     `json:"assign_id"`
	CreatedAt  time.Time  `json:"created_at"`
	User       *user.User `json:"user,omitempty"`
	Team       *team.Team `json:"team,omitempty"`
}

type AddAssignmentRequest struct {
	AssignType string `json:"assign_type" binding:"required"`
	AssignID   string `json:"assign_id" binding:"required"`
}

type CreateTaskRequest struct {
	Title       string     `json:"title" binding:"required"`
	Description string     `json:"description"`
	Tags        []string   `json:"tags"`
	Priority    string     `json:"priority"`
	DueDate     *time.Time `json:"due_date"`
}

type UpdateTaskRequest struct {
	Title       string     `json:"title,omitempty"`
	Description string     `json:"description,omitempty"`
	Status      string     `json:"status,omitempty"`
	Tags        []string   `json:"tags,omitempty"`
	Priority    string     `json:"priority,omitempty"`
	DueDate     *time.Time `json:"due_date,omitempty"`
}

type UpdateTaskStatusRequest struct {
	Status string `json:"status" binding:"required"`
}

type AddCommentRequest struct {
	Comment string `json:"comment" binding:"required"`
}
