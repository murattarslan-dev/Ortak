package service

import (
	"ortak/internal/team"
	"ortak/internal/team/repository"
	"testing"
)

func TestService_GetTeams(t *testing.T) {
	repo := repository.NewMockRepository()
	service := NewService(repo)

	teams, err := service.GetTeams()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(teams) != 0 {
		t.Errorf("Expected 0 teams, got %d", len(teams))
	}
}

func TestService_CreateTeam(t *testing.T) {
	repo := repository.NewMockRepository()
	service := NewService(repo)

	req := team.CreateTeamRequest{
		Name:        "Test Team",
		Description: "Test Description",
	}

	createdTeam, err := service.CreateTeam(req, "1")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if createdTeam.Name != "Test Team" {
		t.Errorf("Expected name Test Team, got %s", createdTeam.Name)
	}

	if createdTeam.OwnerID != "1" {
		t.Errorf("Expected owner ID 1, got %s", createdTeam.OwnerID)
	}

	// Verify team was added
	teams, _ := service.GetTeams()
	if len(teams) != 1 {
		t.Errorf("Expected 1 team, got %d", len(teams))
	}
}
