package service

import (
	"fmt"
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

func (s *Service) GetUserByID(id string) (*user.User, error) {
	user := s.repo.GetByID(id)
	if user == nil {
		return nil, fmt.Errorf("user not found")
	}
	return user, nil
}

func (s *Service) UpdateUser(id string, req user.UpdateUserRequest) (*user.User, error) {
	existingUser := s.repo.GetByID(id)
	if existingUser == nil {
		return nil, fmt.Errorf("user not found")
	}

	// Hash password if provided
	var hashedPassword string
	if req.Password != "" {
		hashed, err := utils.HashPassword(req.Password)
		if err != nil {
			return nil, err
		}
		hashedPassword = hashed
	}

	return s.repo.Update(id, req.Username, req.Email, hashedPassword), nil
}

func (s *Service) DeleteUser(id string) error {
	return s.repo.Delete(id)
}