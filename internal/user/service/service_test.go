package service

import (
	"ortak/internal/user"
	"ortak/internal/user/repository"
	"testing"
)

func TestService_GetUsers(t *testing.T) {
	repo := repository.NewMockRepository()
	service := NewService(repo)

	users, err := service.GetUsers()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(users) != 0 {
		t.Errorf("Expected 0 users, got %d", len(users))
	}
}

func TestService_CreateUser(t *testing.T) {
	repo := repository.NewMockRepository()
	service := NewService(repo)

	req := user.CreateUserRequest{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "password123",
	}

	createdUser, err := service.CreateUser(req)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if createdUser.Username != "testuser" {
		t.Errorf("Expected username testuser, got %s", createdUser.Username)
	}

	if createdUser.Email != "test@example.com" {
		t.Errorf("Expected email test@example.com, got %s", createdUser.Email)
	}

	// Verify user was added
	users, _ := service.GetUsers()
	if len(users) != 1 {
		t.Errorf("Expected 1 user, got %d", len(users))
	}
}