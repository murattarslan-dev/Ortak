package service

import (
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