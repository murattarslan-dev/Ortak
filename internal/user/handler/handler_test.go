package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"ortak/internal/middleware"
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
	router := gin.New()
	router.Use(middleware.ErrorMiddleware())
	router.Use(middleware.FormatterMiddleware())
	router.GET("/users", handler.GetUsers)

	req := httptest.NewRequest("GET", "/users", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}
}

func TestHandler_CreateUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	
	repo := repository.NewMockRepository()
	svc := service.NewService(repo)
	handler := NewHandler(svc)

	reqData := user.CreateUserRequest{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "password123",
	}

	jsonData, _ := json.Marshal(reqData)
	w := httptest.NewRecorder()
	router := gin.New()
	router.Use(middleware.ErrorMiddleware())
	router.Use(middleware.FormatterMiddleware())
	router.POST("/users", handler.CreateUser)

	req := httptest.NewRequest("POST", "/users", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status %d, got %d", http.StatusCreated, w.Code)
	}
}

func TestHandler_GetUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	
	repo := repository.NewMockRepository()
	svc := service.NewService(repo)
	handler := NewHandler(svc)

	// Create a user first
	repo.Create("testuser", "test@example.com", "hashedpass")

	w := httptest.NewRecorder()
	router := gin.New()
	router.Use(middleware.ErrorMiddleware())
	router.Use(middleware.FormatterMiddleware())
	router.GET("/users/:id", handler.GetUser)

	req := httptest.NewRequest("GET", "/users/1", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}
}

func TestHandler_GetUser_NotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)
	
	repo := repository.NewMockRepository()
	svc := service.NewService(repo)
	handler := NewHandler(svc)

	w := httptest.NewRecorder()
	router := gin.New()
	router.Use(middleware.ErrorMiddleware())
	router.Use(middleware.FormatterMiddleware())
	router.GET("/users/:id", handler.GetUser)

	req := httptest.NewRequest("GET", "/users/999", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status %d, got %d", http.StatusNotFound, w.Code)
	}
}

func TestHandler_UpdateUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	
	repo := repository.NewMockRepository()
	svc := service.NewService(repo)
	handler := NewHandler(svc)

	// Create a user first
	repo.Create("testuser", "test@example.com", "hashedpass")

	reqData := user.UpdateUserRequest{
		Username: "updateduser",
		Email:    "updated@example.com",
	}

	jsonData, _ := json.Marshal(reqData)
	w := httptest.NewRecorder()
	router := gin.New()
	router.Use(middleware.ErrorMiddleware())
	router.Use(middleware.FormatterMiddleware())
	router.PUT("/users/:id", handler.UpdateUser)

	req := httptest.NewRequest("PUT", "/users/1", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}
}

func TestHandler_UpdateUser_NotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)
	
	repo := repository.NewMockRepository()
	svc := service.NewService(repo)
	handler := NewHandler(svc)

	reqData := user.UpdateUserRequest{
		Username: "updateduser",
	}

	jsonData, _ := json.Marshal(reqData)
	w := httptest.NewRecorder()
	router := gin.New()
	router.Use(middleware.ErrorMiddleware())
	router.Use(middleware.FormatterMiddleware())
	router.PUT("/users/:id", handler.UpdateUser)

	req := httptest.NewRequest("PUT", "/users/999", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("Expected status %d, got %d", http.StatusInternalServerError, w.Code)
	}
}

func TestHandler_DeleteUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	
	repo := repository.NewMockRepository()
	svc := service.NewService(repo)
	handler := NewHandler(svc)

	// Create a user first
	repo.Create("testuser", "test@example.com", "hashedpass")

	w := httptest.NewRecorder()
	router := gin.New()
	router.Use(middleware.ErrorMiddleware())
	router.Use(middleware.FormatterMiddleware())
	router.DELETE("/users/:id", handler.DeleteUser)

	req := httptest.NewRequest("DELETE", "/users/1", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}
}

func TestHandler_DeleteUser_NotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)
	
	repo := repository.NewMockRepository()
	svc := service.NewService(repo)
	handler := NewHandler(svc)

	w := httptest.NewRecorder()
	router := gin.New()
	router.Use(middleware.ErrorMiddleware())
	router.Use(middleware.FormatterMiddleware())
	router.DELETE("/users/:id", handler.DeleteUser)

	req := httptest.NewRequest("DELETE", "/users/999", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("Expected status %d, got %d", http.StatusInternalServerError, w.Code)
	}
}