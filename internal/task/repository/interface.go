package repository

import (
	"ortak/internal/task"
	"time"
)

type Repository interface {
	GetAll() []task.Task
	GetByID(id string) *task.Task
	GetByIDWithComments(id string) *task.Task
	Create(title, description, createdBy string, tags []string, priority string, dueDate *time.Time) *task.Task
	Update(id, title, description, status string, tags []string, priority string, dueDate *time.Time) *task.Task
	UpdateStatus(id, status string) *task.Task
	Delete(id string) error
	AddComment(taskID, userID, comment string) *task.TaskComment
	AddAssignment(taskID, assignType, assignID string) *task.TaskAssignment
	DeleteAssignment(assignmentID string) error
}