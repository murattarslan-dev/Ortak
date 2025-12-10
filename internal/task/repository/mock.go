package repository

import (
	"fmt"
	"ortak/internal/task"
)

type MockRepository struct {
	tasks    []task.Task
	comments []task.TaskComment
}

func NewMockRepository() Repository {
	return &MockRepository{
		tasks:    make([]task.Task, 0),
		comments: make([]task.TaskComment, 0),
	}
}

func (m *MockRepository) GetAll() []task.Task {
	tasks := make([]task.Task, len(m.tasks))
	for i, t := range m.tasks {
		commentCount := m.getCommentCount(t.ID)
		tasks[i] = t
		tasks[i].CommentCount = commentCount
	}
	return tasks
}

func (m *MockRepository) getCommentCount(taskID int) int {
	count := 0
	for _, c := range m.comments {
		if c.TaskID == taskID {
			count++
		}
	}
	return count
}

func (m *MockRepository) Create(title, description string, assigneeID, teamID int, tags []string) *task.Task {
	task := &task.Task{
		ID:           len(m.tasks) + 1,
		Title:        title,
		Description:  description,
		Status:       "todo",
		AssigneeID:   assigneeID,
		TeamID:       teamID,
		Tags:         tags,
		CommentCount: 0,
	}
	m.tasks = append(m.tasks, *task)
	return task
}

func (m *MockRepository) GetByID(id string) *task.Task {
	for _, t := range m.tasks {
		if fmt.Sprintf("%d", t.ID) == id {
			commentCount := m.getCommentCount(t.ID)
			t.CommentCount = commentCount
			return &t
		}
	}
	return nil
}

func (m *MockRepository) GetByIDWithComments(id string) *task.Task {
	for _, t := range m.tasks {
		if fmt.Sprintf("%d", t.ID) == id {
			comments := make([]task.TaskComment, 0)
			for _, c := range m.comments {
				if c.TaskID == t.ID {
					comments = append(comments, c)
				}
			}
			t.Comments = comments
			t.CommentCount = len(comments)
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
			m.tasks[i].CommentCount = m.getCommentCount(t.ID)
			return &m.tasks[i]
		}
	}
	return nil
}

func (m *MockRepository) UpdateStatus(id, status string) *task.Task {
	for i, t := range m.tasks {
		if fmt.Sprintf("%d", t.ID) == id {
			m.tasks[i].Status = status
			m.tasks[i].CommentCount = m.getCommentCount(t.ID)
			return &m.tasks[i]
		}
	}
	return nil
}

func (m *MockRepository) AddComment(taskID, userID int, comment, createdAt string) *task.TaskComment {
	commentObj := &task.TaskComment{
		ID:        len(m.comments) + 1,
		TaskID:    taskID,
		Comment:   comment,
		CreatedAt: createdAt,
		User: task.CommentUser{
			ID:       userID,
			Username: fmt.Sprintf("user%d", userID),
			Email:    fmt.Sprintf("user%d@test.com", userID),
		},
	}
	m.comments = append(m.comments, *commentObj)
	return commentObj
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
