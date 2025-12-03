package service

import (
	"ortak/internal/user"
	"ortak/internal/user/repository"
	"ortak/pkg/utils"
)

type Service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GetUsers() ([]user.User, error) {
	return s.repo.GetAll(), nil
}

func (s *Service) CreateUser(req user.CreateUserRequest) (*user.User, error) {
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	return s.repo.Create(req.Username, req.Email, hashedPassword), nil
}