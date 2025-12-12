package repository

import (
	"ortak/internal/task"
	"time"
)

type Repository interface {
	GetAll() ([]task.Task, error)
	GetByID(id string) (*task.Task, error)
	GetByIDWithComments(id string) (*task.Task, error)
	Create(title, description, createdBy string, tags []string, priority string, dueDate *time.Time) (*task.Task, error)
	Update(id, title, description, status string, tags []string, priority string, dueDate *time.Time) (*task.Task, error)
	UpdateStatus(id, status string) (*task.Task, error)
	Delete(id string) error
	AddComment(taskID, userID, comment string) (*task.TaskComment, error)
	AddAssignment(taskID, assignType, assignID string) (*task.TaskAssignment, error)
	DeleteAssignment(taskID, assignType, assignID string) error
}
