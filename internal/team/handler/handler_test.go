package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
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
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("GET", "/teams", nil)

	handler.GetTeams(c)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}
}

func TestHandler_CreateTeam(t *testing.T) {
	gin.SetMode(gin.TestMode)
	
	repo := repository.NewMockRepository()
	svc := service.NewService(repo)
	handler := NewHandler(svc)

	req := team.CreateTeamRequest{
		Name:        "Test Team",
		Description: "Test Description",
	}

	jsonData, _ := json.Marshal(req)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("POST", "/teams", bytes.NewBuffer(jsonData))
	c.Request.Header.Set("Content-Type", "application/json")
	c.Set("user_id", 1)

	handler.CreateTeam(c)

	if w.Code != http.StatusCreated {
		t.Errorf("Expected status %d, got %d", http.StatusCreated, w.Code)
	}
}