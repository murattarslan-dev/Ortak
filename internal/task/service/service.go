package service

import (
	"fmt"
	"ortak/internal/task"
	"ortak/internal/task/repository"
)

type Service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GetTasks() ([]task.Task, error) {
	return s.repo.GetAll(), nil
}

func (s *Service) CreateTask(req task.CreateTaskRequest, createdBy string) (*task.Task, error) {
	priority := req.Priority
	if priority == "" {
		priority = "medium"
	}
	return s.repo.Create(req.Title, req.Description, createdBy, req.Tags, priority, req.DueDate), nil
}

func (s *Service) GetTaskByID(id string) (*task.Task, error) {
	task := s.repo.GetByIDWithComments(id)
	if task == nil {
		return nil, fmt.Errorf("task not found")
	}
	return task, nil
}

func (s *Service) UpdateTask(id string, req task.UpdateTaskRequest) (*task.Task, error) {
	existingTask := s.repo.GetByID(id)
	if existingTask == nil {
		return nil, fmt.Errorf("task not found")
	}

	// Validate status if provided
	if req.Status != "" {
		validStatuses := []string{"todo", "in_progress", "done", "cancelled"}
		valid := false
		for _, status := range validStatuses {
			if req.Status == status {
				valid = true
				break
			}
		}
		if !valid {
			return nil, fmt.Errorf("invalid status: must be todo, in_progress, done, or cancelled")
		}
	}

	// Validate priority if provided
	if req.Priority != "" {
		validPriorities := []string{"low", "medium", "high", "urgent"}
		valid := false
		for _, priority := range validPriorities {
			if req.Priority == priority {
				valid = true
				break
			}
		}
		if !valid {
			return nil, fmt.Errorf("invalid priority: must be low, medium, high, or urgent")
		}
	}

	return s.repo.Update(id, req.Title, req.Description, req.Status, req.Tags, req.Priority, req.DueDate), nil
}

func (s *Service) UpdateTaskStatus(id string, req task.UpdateTaskStatusRequest) (*task.Task, error) {
	existingTask := s.repo.GetByID(id)
	if existingTask == nil {
		return nil, fmt.Errorf("task not found")
	}

	// Validate status
	validStatuses := []string{"todo", "in_progress", "done", "cancelled"}
	valid := false
	for _, status := range validStatuses {
		if req.Status == status {
			valid = true
			break
		}
	}
	if !valid {
		return nil, fmt.Errorf("invalid status: must be todo, in_progress, done, or cancelled")
	}

	return s.repo.UpdateStatus(id, req.Status), nil
}

func (s *Service) AddComment(taskID, userID string, req task.AddCommentRequest) (*task.TaskComment, error) {
	existingTask := s.repo.GetByID(taskID)
	if existingTask == nil {
		return nil, fmt.Errorf("task not found")
	}

	return s.repo.AddComment(taskID, userID, req.Comment), nil
}

func (s *Service) AddAssignment(taskID string, req task.AddAssignmentRequest) (*task.TaskAssignment, error) {
	existingTask := s.repo.GetByID(taskID)
	if existingTask == nil {
		return nil, fmt.Errorf("task not found")
	}

	// Validate assign_type
	if req.AssignType != "user" && req.AssignType != "team" {
		return nil, fmt.Errorf("invalid assign_type: must be user or team")
	}

	return s.repo.AddAssignment(taskID, req.AssignType, req.AssignID), nil
}

func (s *Service) DeleteAssignment(assignmentID string) error {
	return s.repo.DeleteAssignment(assignmentID)
}

func (s *Service) DeleteTask(id string) error {
	existingTask := s.repo.GetByID(id)
	if existingTask == nil {
		return fmt.Errorf("task not found")
	}

	return s.repo.Delete(id)
}
