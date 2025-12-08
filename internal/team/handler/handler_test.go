package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"ortak/internal/middleware"
	"ortak/internal/team"
	"ortak/internal/team/repository"
	"ortak/internal/team/service"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestHandler_GetTeams(t *testing.T) {
	gin.SetMode(gin.TestMode)
	
	repo := repository.NewMockRepository()
	svc := service.NewService(repo)
	handler := NewHandler(svc)

	w := httptest.NewRecorder()
	router := gin.New()
	router.Use(middleware.ErrorMiddleware())
	router.Use(middleware.FormatterMiddleware())
	router.GET("/teams", handler.GetTeams)

	req := httptest.NewRequest("GET", "/teams", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}
}

func TestHandler_CreateTeam(t *testing.T) {
	gin.SetMode(gin.TestMode)
	
	repo := repository.NewMockRepository()
	svc := service.NewService(repo)
	handler := NewHandler(svc)

	reqData := team.CreateTeamRequest{
		Name:        "Test Team",
		Description: "Test Description",
	}

	jsonData, _ := json.Marshal(reqData)
	w := httptest.NewRecorder()
	router := gin.New()
	router.Use(middleware.ErrorMiddleware())
	router.Use(middleware.FormatterMiddleware())
	router.POST("/teams", func(c *gin.Context) {
		c.Set("user_id", 1)
		handler.CreateTeam(c)
	})

	req := httptest.NewRequest("POST", "/teams", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status %d, got %d", http.StatusCreated, w.Code)
	}
}

func TestHandler_GetTeam(t *testing.T) {
	gin.SetMode(gin.TestMode)
	
	repo := repository.NewMockRepository()
	svc := service.NewService(repo)
	handler := NewHandler(svc)

	// Create a team first
	repo.Create("Test Team", "Test Description", 1)

	w := httptest.NewRecorder()
	router := gin.New()
	router.Use(middleware.ErrorMiddleware())
	router.Use(middleware.FormatterMiddleware())
	router.GET("/teams/:id", handler.GetTeam)

	req := httptest.NewRequest("GET", "/teams/1", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}
}

func TestHandler_GetTeam_NotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)
	
	repo := repository.NewMockRepository()
	svc := service.NewService(repo)
	handler := NewHandler(svc)

	w := httptest.NewRecorder()
	router := gin.New()
	router.Use(middleware.ErrorMiddleware())
	router.Use(middleware.FormatterMiddleware())
	router.GET("/teams/:id", handler.GetTeam)

	req := httptest.NewRequest("GET", "/teams/999", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status %d, got %d", http.StatusNotFound, w.Code)
	}
}

func TestHandler_UpdateTeam(t *testing.T) {
	gin.SetMode(gin.TestMode)
	
	repo := repository.NewMockRepository()
	svc := service.NewService(repo)
	handler := NewHandler(svc)

	// Create a team first
	repo.Create("Test Team", "Test Description", 1)

	reqData := team.UpdateTeamRequest{
		Name:        "Updated Team",
		Description: "Updated Description",
	}

	jsonData, _ := json.Marshal(reqData)
	w := httptest.NewRecorder()
	router := gin.New()
	router.Use(middleware.ErrorMiddleware())
	router.Use(middleware.FormatterMiddleware())
	router.PUT("/teams/:id", func(c *gin.Context) {
		c.Set("user_id", 1)
		handler.UpdateTeam(c)
	})

	req := httptest.NewRequest("PUT", "/teams/1", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}
}

func TestHandler_UpdateTeam_NotOwner(t *testing.T) {
	gin.SetMode(gin.TestMode)
	
	repo := repository.NewMockRepository()
	svc := service.NewService(repo)
	handler := NewHandler(svc)

	// Create a team with owner ID 1
	repo.Create("Test Team", "Test Description", 1)

	reqData := team.UpdateTeamRequest{
		Name: "Updated Team",
	}

	jsonData, _ := json.Marshal(reqData)
	w := httptest.NewRecorder()
	router := gin.New()
	router.Use(middleware.ErrorMiddleware())
	router.Use(middleware.FormatterMiddleware())
	router.PUT("/teams/:id", func(c *gin.Context) {
		c.Set("user_id", 2) // Different user
		handler.UpdateTeam(c)
	})

	req := httptest.NewRequest("PUT", "/teams/1", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("Expected status %d, got %d", http.StatusInternalServerError, w.Code)
	}
}

func TestHandler_DeleteTeam(t *testing.T) {
	gin.SetMode(gin.TestMode)
	
	repo := repository.NewMockRepository()
	svc := service.NewService(repo)
	handler := NewHandler(svc)

	// Create a team first
	repo.Create("Test Team", "Test Description", 1)

	w := httptest.NewRecorder()
	router := gin.New()
	router.Use(middleware.ErrorMiddleware())
	router.Use(middleware.FormatterMiddleware())
	router.DELETE("/teams/:id", func(c *gin.Context) {
		c.Set("user_id", 1)
		handler.DeleteTeam(c)
	})

	req := httptest.NewRequest("DELETE", "/teams/1", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}
}

func TestHandler_DeleteTeam_NotOwner(t *testing.T) {
	gin.SetMode(gin.TestMode)
	
	repo := repository.NewMockRepository()
	svc := service.NewService(repo)
	handler := NewHandler(svc)

	// Create a team with owner ID 1
	repo.Create("Test Team", "Test Description", 1)

	w := httptest.NewRecorder()
	router := gin.New()
	router.Use(middleware.ErrorMiddleware())
	router.Use(middleware.FormatterMiddleware())
	router.DELETE("/teams/:id", func(c *gin.Context) {
		c.Set("user_id", 2) // Different user
		handler.DeleteTeam(c)
	})

	req := httptest.NewRequest("DELETE", "/teams/1", nil)
	router.ServeHTTP(w, req)

	if w.Code != http.StatusInternalServerError {
		t.Errorf("Expected status %d, got %d", http.StatusInternalServerError, w.Code)
	}
}