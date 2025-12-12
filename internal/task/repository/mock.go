package repository

import (
	"fmt"
	"ortak/internal/task"
	"ortak/internal/team"
	"ortak/internal/user"
	"time"
)

type MockRepository struct {
	tasks       []task.Task
	comments    []task.TaskComment
	assignments []task.TaskAssignment
}

func NewMockRepository() Repository {
	return &MockRepository{
		tasks:       make([]task.Task, 0),
		comments:    make([]task.TaskComment, 0),
		assignments: make([]task.TaskAssignment, 0),
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

func (m *MockRepository) getCommentCount(taskID string) int {
	count := 0
	for _, c := range m.comments {
		if c.TaskID == taskID {
			count++
		}
	}
	return count
}

func (m *MockRepository) Create(title, description, createdBy string, tags []string, priority string, dueDate *time.Time) *task.Task {
	task := &task.Task{
		ID:           "1",
		Title:        title,
		Description:  description,
		Status:       "todo",
		Tags:         tags,
		Priority:     priority,
		DueDate:      dueDate,
		CreatedBy:    createdBy,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		CommentCount: 0,
	}
	m.tasks = append(m.tasks, *task)
	return task
}

func (m *MockRepository) GetByID(id string) *task.Task {
	for _, t := range m.tasks {
		if t.ID == id {
			commentCount := m.getCommentCount(t.ID)
			t.CommentCount = commentCount
			return &t
		}
	}
	return nil
}

func (m *MockRepository) GetByIDWithComments(id string) *task.Task {
	for _, t := range m.tasks {
		if t.ID == id {
			comments := make([]task.TaskComment, 0)
			for _, c := range m.comments {
				if c.TaskID == t.ID {
					comments = append(comments, c)
				}
			}
			assignments := make([]task.TaskAssignment, 0)
			for _, a := range m.assignments {
				if a.TaskID == t.ID {
					assignments = append(assignments, a)
				}
			}
			t.Comments = comments
			t.CommentCount = len(comments)
			t.Assignments = assignments
			return &t
		}
	}
	return nil
}

func (m *MockRepository) Update(id, title, description, status string, tags []string, priority string, dueDate *time.Time) *task.Task {
	for i, t := range m.tasks {
		if t.ID == id {
			if title != "" {
				m.tasks[i].Title = title
			}
			if description != "" {
				m.tasks[i].Description = description
			}
			if status != "" {
				m.tasks[i].Status = status
			}
			if len(tags) > 0 {
				m.tasks[i].Tags = tags
			}
			if priority != "" {
				m.tasks[i].Priority = priority
			}
			if dueDate != nil {
				m.tasks[i].DueDate = dueDate
			}
			m.tasks[i].UpdatedAt = time.Now()
			m.tasks[i].CommentCount = m.getCommentCount(t.ID)
			return &m.tasks[i]
		}
	}
	return nil
}

func (m *MockRepository) UpdateStatus(id, status string) *task.Task {
	for i, t := range m.tasks {
		if t.ID == id {
			m.tasks[i].Status = status
			m.tasks[i].UpdatedAt = time.Now()
			m.tasks[i].CommentCount = m.getCommentCount(t.ID)
			return &m.tasks[i]
		}
	}
	return nil
}

func (m *MockRepository) AddComment(taskID, userID, comment string) *task.TaskComment {
	commentObj := &task.TaskComment{
		ID:        "1",
		TaskID:    taskID,
		UserID:    userID,
		Comment:   comment,
		CreatedAt: time.Now(),
		User: &user.User{
			ID:        userID,
			Username:  fmt.Sprintf("user%s", userID),
			Email:     fmt.Sprintf("user%s@test.com", userID),
			FirstName: "Test",
			LastName:  "User",
		},
	}
	m.comments = append(m.comments, *commentObj)
	return commentObj
}

func (m *MockRepository) AddAssignment(taskID, assignType, assignID string) *task.TaskAssignment {
	assignment := &task.TaskAssignment{
		ID:         "1",
		TaskID:     taskID,
		AssignType: assignType,
		AssignID:   assignID,
		CreatedAt:  time.Now(),
	}

	if assignType == "user" {
		assignment.User = &user.User{
			ID:        assignID,
			Username:  fmt.Sprintf("user%s", assignID),
			Email:     fmt.Sprintf("user%s@test.com", assignID),
			FirstName: "Test",
			LastName:  "User",
		}
	} else if assignType == "team" {
		assignment.Team = &team.Team{
			ID:          assignID,
			Name:        fmt.Sprintf("Team %s", assignID),
			Description: "Test team",
		}
	}

	m.assignments = append(m.assignments, *assignment)
	return assignment
}

func (m *MockRepository) DeleteAssignment(assignmentID string) error {
	for i, a := range m.assignments {
		if a.ID == assignmentID {
			m.assignments = append(m.assignments[:i], m.assignments[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("assignment not found")
}

func (m *MockRepository) Delete(id string) error {
	for i, t := range m.tasks {
		if t.ID == id {
			m.tasks = append(m.tasks[:i], m.tasks[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("task not found")
}
