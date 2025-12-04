package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"ortak/internal/auth"
	"ortak/internal/auth/repository"
	"ortak/internal/auth/service"
	"ortak/internal/middleware"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestHandler_Register(t *testing.T) {
	gin.SetMode(gin.TestMode)
	
	repo := repository.NewMockRepository()
	svc := service.NewService(repo)
	handler := NewHandler(svc)

	reqData := auth.RegisterRequest{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "password123",
	}

	jsonData, _ := json.Marshal(reqData)
	w := httptest.NewRecorder()
	router := gin.New()
	router.Use(middleware.ErrorMiddleware())
	router.Use(middleware.FormatterMiddleware())
	router.POST("/register", handler.Register)

	req := httptest.NewRequest("POST", "/register", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

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
	router := gin.New()
	router.Use(middleware.ErrorMiddleware())
	router.Use(middleware.FormatterMiddleware())
	router.POST("/login", handler.Login)

	req := httptest.NewRequest("POST", "/login", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}
}