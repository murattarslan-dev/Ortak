package service

import (
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

func (s *Service) CreateTask(req task.CreateTaskRequest) (*task.Task, error) {
	return s.repo.Create(req.Title, req.Description, req.AssigneeID, req.TeamID), nil
}