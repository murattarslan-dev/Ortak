package repository

import "ortak/internal/team"

type MockRepository struct {
	teams []team.Team
}

func NewMockRepository() Repository {
	return &MockRepository{
		teams: make([]team.Team, 0),
	}
}

func (m *MockRepository) GetAll() []team.Team {
	return m.teams
}

func (m *MockRepository) Create(name, description string, ownerID int) *team.Team {
	team := &team.Team{
		ID:          len(m.teams) + 1,
		Name:        name,
		Description: description,
		OwnerID:     ownerID,
	}
	m.teams = append(m.teams, *team)
	return team
}