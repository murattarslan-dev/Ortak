package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"ortak/internal/auth"
	"ortak/internal/auth/repository"
	"ortak/internal/auth/service"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestHandler_Register(t *testing.T) {
	gin.SetMode(gin.TestMode)
	
	repo := repository.NewMockRepository()
	svc := service.NewService(repo)
	handler := NewHandler(svc)

	req := auth.RegisterRequest{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "password123",
	}

	jsonData, _ := json.Marshal(req)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("POST", "/register", bytes.NewBuffer(jsonData))
	c.Request.Header.Set("Content-Type", "application/json")

	handler.Register(c)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status %d, got %d", http.StatusCreated, w.Code)
	}
}

func TestHandler_Login(t *testing.T) {
	gin.SetMode(gin.TestMode)
	
	repo := repository.NewMockRepository()
	svc := service.NewService(repo)
	handler := NewHandler(svc)

	// First register
	registerReq := auth.RegisterRequest{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "password123",
	}
	svc.Register(registerReq)

	// Then login
	loginReq := auth.LoginRequest{
		Email:    "test@example.com",
		Password: "password123",
	}

	jsonData, _ := json.Marshal(loginReq)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("POST", "/login", bytes.NewBuffer(jsonData))
	c.Request.Header.Set("Content-Type", "application/json")

	handler.Login(c)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}
}