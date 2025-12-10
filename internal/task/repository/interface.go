package repository

import "ortak/internal/task"

type Repository interface {
	GetAll() []task.Task
	GetByID(id string) *task.Task
	GetByIDWithComments(id string) *task.Task
	Create(title, description string, assigneeID, teamID int, tags []string) *task.Task
	Update(id, title, description, status string, assigneeID int, tags []string) *task.Task
	UpdateStatus(id, status string) *task.Task
	Delete(id string) error
	AddComment(taskID, userID int, comment, createdAt string) *task.TaskComment
}