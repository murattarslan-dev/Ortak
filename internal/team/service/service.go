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
	return s.repo.Create(req.Name, req.Description, ownerID), nil
}

func (s *Service) GetTeamByID(id string) (*team.Team, error) {
	team := s.repo.GetByID(id)
	if team == nil {
		return nil, fmt.Errorf("team not found")
	}
	return team, nil
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