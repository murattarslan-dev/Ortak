package repository

import "ortak/internal/task"

type Repository interface {
	GetAll() []task.Task
	Create(title, description string, assigneeID, teamID int) *task.Task
}