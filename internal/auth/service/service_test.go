package service

import (
	"ortak/internal/auth"
	"ortak/internal/auth/repository"
	"testing"
)

func TestService_Register(t *testing.T) {
	repo := repository.NewMockRepository()
	service := NewService(repo)

	req := auth.RegisterRequest{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "password123",
	}

	response, err := service.Register(req)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if response.User.Username != "testuser" {
		t.Errorf("Expected username testuser, got %s", response.User.Username)
	}

	if response.Token == "" {
		t.Error("Expected token to be generated")
	}
}

func TestService_Login(t *testing.T) {
	repo := repository.NewMockRepository()
	service := NewService(repo)

	// First register a user
	registerReq := auth.RegisterRequest{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "password123",
	}
	service.Register(registerReq)

	// Then login
	loginReq := auth.LoginRequest{
		Email:    "test@example.com",
		Password: "password123",
	}

	response, err := service.Login(loginReq)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if response.User.Email != "test@example.com" {
		t.Errorf("Expected email test@example.com, got %s", response.User.Email)
	}
}

func TestService_Login_InvalidCredentials(t *testing.T) {
	repo := repository.NewMockRepository()
	service := NewService(repo)

	loginReq := auth.LoginRequest{
		Email:    "nonexistent@example.com",
		Password: "wrongpassword",
	}

	_, err := service.Login(loginReq)
	if err == nil {
		t.Error("Expected error for invalid credentials")
	}
}