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
	return s.repo.GetAll()
}

func (s *Service) CreateTask(req task.CreateTaskRequest, createdBy string) (*task.Task, error) {
	priority := req.Priority
	if priority == "" {
		priority = "medium"
	}
	return s.repo.Create(req.Title, req.Description, createdBy, req.Tags, priority, req.DueDate)
}

func (s *Service) GetTaskByID(id string) (*task.Task, error) {
	return s.repo.GetByIDWithComments(id)
}

func (s *Service) UpdateTask(id string, req task.UpdateTaskRequest) (*task.Task, error) {
	_, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
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

	return s.repo.Update(id, req.Title, req.Description, req.Status, req.Tags, req.Priority, req.DueDate)
}

func (s *Service) UpdateTaskStatus(id string, req task.UpdateTaskStatusRequest) (*task.Task, error) {
	_, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
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

	return s.repo.UpdateStatus(id, req.Status)
}

func (s *Service) AddComment(taskID, userID string, req task.AddCommentRequest) (*task.TaskComment, error) {
	_, err := s.repo.GetByID(taskID)
	if err != nil {
		return nil, err
	}

	return s.repo.AddComment(taskID, userID, req.Comment)
}

func (s *Service) AddAssignment(taskID string, req task.AddAssignmentRequest) (*task.TaskAssignment, error) {
	_, err := s.repo.GetByID(taskID)
	if err != nil {
		return nil, err
	}

	// Validate assign_type
	if req.AssignType != "user" && req.AssignType != "team" {
		return nil, fmt.Errorf("invalid assign_type: must be user or team")
	}

	return s.repo.AddAssignment(taskID, req.AssignType, req.AssignID)
}

func (s *Service) DeleteAssignment(taskID, assignType, assignID string) error {
	return s.repo.DeleteAssignment(taskID, assignType, assignID)
}

func (s *Service) DeleteTask(id string) error {
	_, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	return s.repo.Delete(id)
}
