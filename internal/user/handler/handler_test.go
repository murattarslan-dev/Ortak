package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"ortak/internal/user"
	"ortak/internal/user/repository"
	"ortak/internal/user/service"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestHandler_GetUsers(t *testing.T) {
	gin.SetMode(gin.TestMode)
	
	repo := repository.NewMockRepository()
	svc := service.NewService(repo)
	handler := NewHandler(svc)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("GET", "/users", nil)

	handler.GetUsers(c)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}
}

func TestHandler_CreateUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	
	repo := repository.NewMockRepository()
	svc := service.NewService(repo)
	handler := NewHandler(svc)

	req := user.CreateUserRequest{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "password123",
	}

	jsonData, _ := json.Marshal(req)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("POST", "/users", bytes.NewBuffer(jsonData))
	c.Request.Header.Set("Content-Type", "application/json")

	handler.CreateUser(c)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status %d, got %d", http.StatusCreated, w.Code)
	}
}