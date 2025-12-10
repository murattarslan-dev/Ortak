package repository

import (
	"fmt"
	"ortak/internal/task"
)

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

func (m *MockRepository) Create(title, description string, assigneeID, teamID int, tags []string) *task.Task {
	task := &task.Task{
		ID:          len(m.tasks) + 1,
		Title:       title,
		Description: description,
		Status:      "todo",
		AssigneeID:  assigneeID,
		TeamID:      teamID,
		Tags:        tags,
	}
	m.tasks = append(m.tasks, *task)
	return task
}

func (m *MockRepository) GetByID(id string) *task.Task {
	for _, t := range m.tasks {
		if fmt.Sprintf("%d", t.ID) == id {
			return &t
		}
	}
	return nil
}

func (m *MockRepository) Update(id, title, description, status string, assigneeID int, tags []string) *task.Task {
	for i, t := range m.tasks {
		if fmt.Sprintf("%d", t.ID) == id {
			if title != "" {
				m.tasks[i].Title = title
			}
			if description != "" {
				m.tasks[i].Description = description
			}
			if status != "" {
				m.tasks[i].Status = status
			}
			if assigneeID != 0 {
				m.tasks[i].AssigneeID = assigneeID
			}
			if len(tags) > 0 {
				m.tasks[i].Tags = tags
			}
			return &m.tasks[i]
		}
	}
	return nil
}

func (m *MockRepository) UpdateStatus(id, status string) *task.Task {
	for i, t := range m.tasks {
		if fmt.Sprintf("%d", t.ID) == id {
			m.tasks[i].Status = status
			return &m.tasks[i]
		}
	}
	return nil
}

func (m *MockRepository) Delete(id string) error {
	for i, t := range m.tasks {
		if fmt.Sprintf("%d", t.ID) == id {
			m.tasks = append(m.tasks[:i], m.tasks[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("task not found")
}