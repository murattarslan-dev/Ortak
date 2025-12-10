package service

import (
	"fmt"
	"ortak/internal/task"
	"ortak/internal/task/repository"
	"time"
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

func (s *Service) CreateTask(req task.CreateTaskRequest) (*task.Task, error) {
	return s.repo.Create(req.Title, req.Description, req.AssigneeID, req.TeamID, req.Tags), nil
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
		validStatuses := []string{"todo", "in_progress", "done"}
		valid := false
		for _, status := range validStatuses {
			if req.Status == status {
				valid = true
				break
			}
		}
		if !valid {
			return nil, fmt.Errorf("invalid status: must be todo, in_progress, or done")
		}
	}

	return s.repo.Update(id, req.Title, req.Description, req.Status, req.AssigneeID, req.Tags), nil
}

func (s *Service) UpdateTaskStatus(id string, req task.UpdateTaskStatusRequest) (*task.Task, error) {
	existingTask := s.repo.GetByID(id)
	if existingTask == nil {
		return nil, fmt.Errorf("task not found")
	}

	// Validate status
	validStatuses := []string{"todo", "in_progress", "done"}
	valid := false
	for _, status := range validStatuses {
		if req.Status == status {
			valid = true
			break
		}
	}
	if !valid {
		return nil, fmt.Errorf("invalid status: must be todo, in_progress, or done")
	}

	return s.repo.UpdateStatus(id, req.Status), nil
}

func (s *Service) AddComment(taskID string, userID int, req task.AddCommentRequest) (*task.TaskComment, error) {
	existingTask := s.repo.GetByID(taskID)
	if existingTask == nil {
		return nil, fmt.Errorf("task not found")
	}

	// Convert taskID to int
	taskIDInt := 0
	for i := 1; i <= 1000; i++ {
		if fmt.Sprintf("%d", i) == taskID {
			taskIDInt = i
			break
		}
	}

	createdAt := time.Now().Format(time.RFC3339)
	return s.repo.AddComment(taskIDInt, userID, req.Comment, createdAt), nil
}

func (s *Service) DeleteTask(id string) error {
	existingTask := s.repo.GetByID(id)
	if existingTask == nil {
		return fmt.Errorf("task not found")
	}

	return s.repo.Delete(id)
}