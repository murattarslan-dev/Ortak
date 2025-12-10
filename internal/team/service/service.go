package service

import (
	"fmt"
	"ortak/internal/team"
	"ortak/internal/team/repository"
)

type Service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) GetTeams() ([]team.Team, error) {
	return s.repo.GetAll(), nil
}

func (s *Service) CreateTeam(req team.CreateTeamRequest, ownerID int) (*team.Team, error) {
	// Create the team
	createdTeam := s.repo.Create(req.Name, req.Description, ownerID)
	
	// Add creator as owner member
	teamID := fmt.Sprintf("%d", createdTeam.ID)
	s.repo.AddMember(teamID, ownerID, "owner")
	
	return createdTeam, nil
}

func (s *Service) GetTeamByID(id string) (*team.Team, error) {
	team := s.repo.GetByID(id)
	if team == nil {
		return nil, fmt.Errorf("team not found")
	}
	return team, nil
}

func (s *Service) GetTeamWithMembers(id string) (*team.TeamWithMembers, error) {
	teamData := s.repo.GetByID(id)
	if teamData == nil {
		return nil, fmt.Errorf("team not found")
	}

	members := s.repo.GetTeamMembers(id)

	return &team.TeamWithMembers{
		ID:          teamData.ID,
		Name:        teamData.Name,
		Description: teamData.Description,
		OwnerID:     teamData.OwnerID,
		Members:     members,
	}, nil
}

func (s *Service) UpdateTeam(id string, req team.UpdateTeamRequest, userID int) (*team.Team, error) {
	existingTeam := s.repo.GetByID(id)
	if existingTeam == nil {
		return nil, fmt.Errorf("team not found")
	}

	// Check if user is the owner
	if existingTeam.OwnerID != userID {
		return nil, fmt.Errorf("only team owner can update team")
	}

	return s.repo.Update(id, req.Name, req.Description), nil
}

func (s *Service) DeleteTeam(id string, userID int) error {
	existingTeam := s.repo.GetByID(id)
	if existingTeam == nil {
		return fmt.Errorf("team not found")
	}

	// Check if user is the owner
	if existingTeam.OwnerID != userID {
		return fmt.Errorf("only team owner can delete team")
	}

	return s.repo.Delete(id)
}

func (s *Service) AddTeamMember(teamID string, memberUserID int, role string, requestUserID int) (*team.TeamMember, error) {
	// Check if team exists and user is owner
	existingTeam := s.repo.GetByID(teamID)
	if existingTeam == nil {
		return nil, fmt.Errorf("team not found")
	}

	if existingTeam.OwnerID != requestUserID {
		return nil, fmt.Errorf("only team owner can add members")
	}

	return s.repo.AddMember(teamID, memberUserID, role)
}

func (s *Service) RemoveTeamMember(teamID, memberUserID string, requestUserID int) error {
	// Check if team exists and user is owner
	existingTeam := s.repo.GetByID(teamID)
	if existingTeam == nil {
		return fmt.Errorf("team not found")
	}

	if existingTeam.OwnerID != requestUserID {
		return fmt.Errorf("only team owner can remove members")
	}

	return s.repo.RemoveMember(teamID, memberUserID)
}

func (s *Service) UpdateMemberRole(teamID, memberUserID, role string, requestUserID int) (*team.TeamMember, error) {
	// Check if team exists and user is owner
	existingTeam := s.repo.GetByID(teamID)
	if existingTeam == nil {
		return nil, fmt.Errorf("team not found")
	}

	if existingTeam.OwnerID != requestUserID {
		return nil, fmt.Errorf("only team owner can update member roles")
	}

	return s.repo.UpdateMemberRole(teamID, memberUserID, role)
}