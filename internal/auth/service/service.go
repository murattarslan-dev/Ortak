package service

import (
	"errors"
	"ortak/internal/auth"
	"ortak/internal/auth/repository"
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

func (s *Service) Register(req auth.RegisterRequest) (*auth.AuthResponse, error) {
	if s.repo.GetUserByEmail(req.Email) != nil {
		return nil, errors.New("user already exists")
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := s.repo.CreateUser(req.Username, req.Email, hashedPassword)

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return nil, err
	}

	s.repo.AddToken(token, user.ID)

	return &auth.AuthResponse{
		Token: token,
		User: auth.User{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
		},
	}, nil
}

func (s *Service) Login(req auth.LoginRequest) (*auth.AuthResponse, error) {
	user := s.repo.GetUserByEmail(req.Email)
	if user == nil {
		return nil, errors.New("invalid credentials")
	}

	if !utils.CheckPasswordHash(req.Password, user.Password) {
		return nil, errors.New("invalid credentials")
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return nil, err
	}

	s.repo.AddToken(token, user.ID)

	return &auth.AuthResponse{
		Token: token,
		User: auth.User{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
		},
	}, nil
}