package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"ortak/internal/middleware"
	"ortak/internal/task"
	"ortak/internal/task/repository"
	"ortak/internal/task/service"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

func TestHandler_GetTasks(t *testing.T) {
	gin.SetMode(gin.TestMode)

	repo := repository.NewMockRepository()
	svc := service.NewService(repo)
	handler := NewHandler(svc)

	w := httptest.NewRecorder()
	router := gin.New()
	router.Use(middleware.ErrorMiddleware())
	router.Use(middleware.FormatterMiddleware())
	router.GET("/tasks", handler.GetTasks)

	req := httptest.NewRequest("GET", "/tasks", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}
}

func TestHandler_CreateTask(t *testing.T) {
	gin.SetMode(gin.TestMode)

	repo := repository.NewMockRepository()
	svc := service.NewService(repo)
	handler := NewHandler(svc)

	reqData := task.CreateTaskRequest{
		Title:       "Test Task",
		Description: "Test Description",
		Priority:    "high",
		DueDate:     &time.Time{},
		Tags:        []string{"test", "api"},
	}

	jsonData, _ := json.Marshal(reqData)
	w := httptest.NewRecorder()
	router := gin.New()
	router.Use(middleware.ErrorMiddleware())
	router.Use(middleware.FormatterMiddleware())
	router.POST("/tasks", handler.CreateTask)

	req := httptest.NewRequest("POST", "/tasks", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status %d, got %d", http.StatusCreated, w.Code)
	}
}

func TestHandler_GetTask(t *testing.T) {
	gin.SetMode(gin.TestMode)

	repo := repository.NewMockRepository()
	svc := service.NewService(repo)
	handler := NewHandler(svc)

	// Create a task first
	repo.Create("Test Task", "Test Description", "1", []string{"test"}, "high", &time.Time{})

	w := httptest.NewRecorder()
	router := gin.New()
	router.Use(middleware.ErrorMiddleware())
	router.Use(middleware.FormatterMiddleware())
	router.GET("/tasks/:id", handler.GetTask)

	req := httptest.NewRequest("GET", "/tasks/1", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}
}

func TestHandler_GetTask_NotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)

	repo := repository.NewMockRepository()
	svc := service.NewService(repo)
	handler := NewHandler(svc)

	w := httptest.NewRecorder()
	router := gin.New()
	router.Use(middleware.ErrorMiddleware())
	router.Use(middleware.FormatterMiddleware())
	router.GET("/tasks/:id", handler.GetTask)

	req := httptest.NewRequest("GET", "/tasks/999", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status %d, got %d", http.StatusNotFound, w.Code)
	}
}

func TestHandler_UpdateTask(t *testing.T) {
	gin.SetMode(gin.TestMode)

	repo := repository.NewMockRepository()
	svc := service.NewService(repo)
	handler := NewHandler(svc)

	// Create a task first
	repo.Create("Test Task", "Test Description", "1", []string{"test"}, "high", &time.Time{})

	reqData := task.UpdateTaskRequest{
		Title:       "Updated Task",
		Description: "Updated Description",
		Status:      "in_progress",
	}

	jsonData, _ := json.Marshal(reqData)
	w := httptest.NewRecorder()
	router := gin.New()
	router.Use(middleware.ErrorMiddleware())
	router.Use(middleware.FormatterMiddleware())
	router.PUT("/tasks/:id", handler.UpdateTask)

	req := httptest.NewRequest("PUT", "/tasks/1", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}
}

func TestHandler_UpdateTask_InvalidStatus(t *testing.T) {
	gin.SetMode(gin.TestMode)

	repo := repository.NewMockRepository()
	svc := service.NewService(repo)
	handler := NewHandler(svc)

	// Create a task first
	repo.Create("Test Task", "Test Description", "1", []string{"test"}, "high", &time.Time{})

	// Test different invalid statuses
	invalidStatuses := []string{"invalid_status", "pending", "completed", "cancelled", ""}

	for _, invalidStatus := range invalidStatuses {
		reqData := task.UpdateTaskRequest{
			Status: invalidStatus,
		}

		jsonData, _ := json.Marshal(reqData)
		w := httptest.NewRecorder()
		router := gin.New()
		router.Use(middleware.ErrorMiddleware())
		router.Use(middleware.FormatterMiddleware())
		router.PUT("/tasks/:id", handler.UpdateTask)

		req := httptest.NewRequest("PUT", "/tasks/1", bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")
		router.ServeHTTP(w, req)

		if invalidStatus != "" && w.Code != http.StatusBadRequest {
			t.Errorf("Expected status %d for invalid status '%s', got %d", http.StatusBadRequest, invalidStatus, w.Code)
		}
	}
}

func TestHandler_DeleteTask(t *testing.T) {
	gin.SetMode(gin.TestMode)

	repo := repository.NewMockRepository()
	svc := service.NewService(repo)
	handler := NewHandler(svc)

	// Create a task first
	repo.Create("Test Task", "Test Description", "1", []string{"test"}, "high", &time.Time{})

	w := httptest.NewRecorder()
	router := gin.New()
	router.Use(middleware.ErrorMiddleware())
	router.Use(middleware.FormatterMiddleware())
	router.DELETE("/tasks/:id", handler.DeleteTask)

	req := httptest.NewRequest("DELETE", "/tasks/1", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}
}

func TestHandler_DeleteTask_NotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)

	repo := repository.NewMockRepository()
	svc := service.NewService(repo)
	handler := NewHandler(svc)

	w := httptest.NewRecorder()
	router := gin.New()
	router.Use(middleware.ErrorMiddleware())
	router.Use(middleware.FormatterMiddleware())
	router.DELETE("/tasks/:id", handler.DeleteTask)

	req := httptest.NewRequest("DELETE", "/tasks/999", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status %d, got %d", http.StatusNotFound, w.Code)
	}
}

func TestHandler_UpdateTaskStatus(t *testing.T) {
	gin.SetMode(gin.TestMode)

	repo := repository.NewMockRepository()
	svc := service.NewService(repo)
	handler := NewHandler(svc)

	// Create a task first
	repo.Create("Test Task", "Test Description", "1", []string{"test"}, "high", &time.Time{})

	reqData := task.UpdateTaskStatusRequest{
		Status: "in_progress",
	}

	jsonData, _ := json.Marshal(reqData)
	w := httptest.NewRecorder()
	router := gin.New()
	router.Use(middleware.ErrorMiddleware())
	router.Use(middleware.FormatterMiddleware())
	router.PATCH("/tasks/:id/status", handler.UpdateTaskStatus)

	req := httptest.NewRequest("PATCH", "/tasks/1/status", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}
}

func TestHandler_UpdateTaskStatus_InvalidStatus(t *testing.T) {
	gin.SetMode(gin.TestMode)

	repo := repository.NewMockRepository()
	svc := service.NewService(repo)
	handler := NewHandler(svc)

	// Create a task first
	repo.Create("Test Task", "Test Description", "1", []string{"test"}, "high", &time.Time{})

	reqData := task.UpdateTaskStatusRequest{
		Status: "invalid_status",
	}

	jsonData, _ := json.Marshal(reqData)
	w := httptest.NewRecorder()
	router := gin.New()
	router.Use(middleware.ErrorMiddleware())
	router.Use(middleware.FormatterMiddleware())
	router.PATCH("/tasks/:id/status", handler.UpdateTaskStatus)

	req := httptest.NewRequest("PATCH", "/tasks/1/status", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status %d, got %d", http.StatusBadRequest, w.Code)
	}
}

func TestHandler_UpdateTaskStatus_NotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)

	repo := repository.NewMockRepository()
	svc := service.NewService(repo)
	handler := NewHandler(svc)

	reqData := task.UpdateTaskStatusRequest{
		Status: "in_progress",
	}

	jsonData, _ := json.Marshal(reqData)
	w := httptest.NewRecorder()
	router := gin.New()
	router.Use(middleware.ErrorMiddleware())
	router.Use(middleware.FormatterMiddleware())
	router.PATCH("/tasks/:id/status", handler.UpdateTaskStatus)

	req := httptest.NewRequest("PATCH", "/tasks/999/status", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status %d, got %d", http.StatusNotFound, w.Code)
	}
}

func TestHandler_AddComment(t *testing.T) {
	gin.SetMode(gin.TestMode)

	repo := repository.NewMockRepository()
	svc := service.NewService(repo)
	handler := NewHandler(svc)

	// Create a task first
	repo.Create("Test Task", "Test Description", "1", []string{"test"}, "high", &time.Time{})

	reqData := task.AddCommentRequest{
		Comment: "This is a test comment",
	}

	jsonData, _ := json.Marshal(reqData)
	w := httptest.NewRecorder()
	router := gin.New()
	router.Use(middleware.ErrorMiddleware())
	router.Use(middleware.FormatterMiddleware())
	router.Use(func(c *gin.Context) {
		c.Set("userID", 1)
		c.Next()
	})
	router.POST("/tasks/:id/comments", handler.AddComment)

	req := httptest.NewRequest("POST", "/tasks/1/comments", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status %d, got %d", http.StatusCreated, w.Code)
	}
}

func TestHandler_AddComment_TaskNotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)

	repo := repository.NewMockRepository()
	svc := service.NewService(repo)
	handler := NewHandler(svc)

	reqData := task.AddCommentRequest{
		Comment: "This is a test comment",
	}

	jsonData, _ := json.Marshal(reqData)
	w := httptest.NewRecorder()
	router := gin.New()
	router.Use(middleware.ErrorMiddleware())
	router.Use(middleware.FormatterMiddleware())
	router.Use(func(c *gin.Context) {
		c.Set("userID", 1)
		c.Next()
	})
	router.POST("/tasks/:id/comments", handler.AddComment)

	req := httptest.NewRequest("POST", "/tasks/999/comments", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status %d, got %d", http.StatusNotFound, w.Code)
	}
}

func TestHandler_AddComment_Unauthorized(t *testing.T) {
	gin.SetMode(gin.TestMode)

	repo := repository.NewMockRepository()
	svc := service.NewService(repo)
	handler := NewHandler(svc)

	reqData := task.AddCommentRequest{
		Comment: "This is a test comment",
	}

	jsonData, _ := json.Marshal(reqData)
	w := httptest.NewRecorder()
	router := gin.New()
	router.Use(middleware.ErrorMiddleware())
	router.Use(middleware.FormatterMiddleware())
	router.POST("/tasks/:id/comments", handler.AddComment)

	req := httptest.NewRequest("POST", "/tasks/1/comments", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized {
		t.Errorf("Expected status %d, got %d", http.StatusUnauthorized, w.Code)
	}
}

func TestHandler_AddAssignment(t *testing.T) {
	gin.SetMode(gin.TestMode)

	repo := repository.NewMockRepository()
	svc := service.NewService(repo)
	handler := NewHandler(svc)

	// Create a task first
	repo.Create("Test Task", "Test Description", "1", []string{"test"}, "high", &time.Time{})

	reqData := task.AddAssignmentRequest{
		AssignType: "user",
		AssignID:   "5",
	}

	jsonData, _ := json.Marshal(reqData)
	w := httptest.NewRecorder()
	router := gin.New()
	router.Use(middleware.ErrorMiddleware())
	router.Use(middleware.FormatterMiddleware())
	router.POST("/tasks/:id/assignments", handler.AddAssignment)

	req := httptest.NewRequest("POST", "/tasks/1/assignments", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status %d, got %d", http.StatusCreated, w.Code)
	}
}

func TestHandler_AddAssignment_InvalidType(t *testing.T) {
	gin.SetMode(gin.TestMode)

	repo := repository.NewMockRepository()
	svc := service.NewService(repo)
	handler := NewHandler(svc)

	// Create a task first
	repo.Create("Test Task", "Test Description", "1", []string{"test"}, "high", &time.Time{})

	reqData := task.AddAssignmentRequest{
		AssignType: "invalid",
		AssignID:   "5",
	}

	jsonData, _ := json.Marshal(reqData)
	w := httptest.NewRecorder()
	router := gin.New()
	router.Use(middleware.ErrorMiddleware())
	router.Use(middleware.FormatterMiddleware())
	router.POST("/tasks/:id/assignments", handler.AddAssignment)

	req := httptest.NewRequest("POST", "/tasks/1/assignments", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status %d, got %d", http.StatusBadRequest, w.Code)
	}
}

func TestHandler_DeleteAssignment(t *testing.T) {
	gin.SetMode(gin.TestMode)

	repo := repository.NewMockRepository()
	svc := service.NewService(repo)
	handler := NewHandler(svc)

	// Create a task and assignment first
	repo.Create("Test Task", "Test Description", "1", []string{"test"}, "high", &time.Time{})
	repo.AddAssignment("1", "user", "5")

	w := httptest.NewRecorder()
	router := gin.New()
	router.Use(middleware.ErrorMiddleware())
	router.Use(middleware.FormatterMiddleware())
	router.DELETE("/tasks/:id/assignments/:assignmentId", handler.DeleteAssignment)

	req := httptest.NewRequest("DELETE", "/tasks/1/assignments/1", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}
}
