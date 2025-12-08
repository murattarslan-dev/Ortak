package repository

import "ortak/internal/task"

type Repository interface {
	GetAll() []task.Task
	GetByID(id string) *task.Task
	Create(title, description string, assigneeID, teamID int) *task.Task
	Update(id, title, description, status string, assigneeID int) *task.Task
	Delete(id string) error
}