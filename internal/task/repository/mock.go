package repository

import "ortak/internal/task"

type MockRepository struct {
	tasks []task.Task
}

func NewMockRepository() Repository {
	return &MockRepository{
		tasks: make([]task.Task, 0),
	}
}

func (m *MockRepository) GetAll() []task.Task {
	return m.tasks
}

func (m *MockRepository) Create(title, description string, assigneeID, teamID int) *task.Task {
	task := &task.Task{
		ID:          len(m.tasks) + 1,
		Title:       title,
		Description: description,
		Status:      "todo",
		AssigneeID:  assigneeID,
		TeamID:      teamID,
	}
	m.tasks = append(m.tasks, *task)
	return task
}