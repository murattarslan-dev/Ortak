package auth

import (
	"ortak/pkg/utils"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Register(req RegisterRequest) (*AuthResponse, error) {
	_, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	// TODO: Save user to database
	user := User{
		ID:       1,
		Username: req.Username,
		Email:    req.Email,
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return nil, err
	}

	return &AuthResponse{
		Token: token,
		User:  user,
	}, nil
}

func (s *Service) Login(req LoginRequest) (*AuthResponse, error) {
	// TODO: Get user from database and verify password
	user := User{
		ID:       1,
		Username: "testuser",
		Email:    req.Email,
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return nil, err
	}

	return &AuthResponse{
		Token: token,
		User:  user,
	}, nil
}