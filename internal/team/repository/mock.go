package repository

import (
	"fmt"
	"ortak/internal/team"
)

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

func (m *MockRepository) GetByID(id string) *team.Team {
	for _, t := range m.teams {
		if fmt.Sprintf("%d", t.ID) == id {
			return &t
		}
	}
	return nil
}

func (m *MockRepository) Update(id, name, description string) *team.Team {
	for i, t := range m.teams {
		if fmt.Sprintf("%d", t.ID) == id {
			if name != "" {
				m.teams[i].Name = name
			}
			if description != "" {
				m.teams[i].Description = description
			}
			return &m.teams[i]
		}
	}
	return nil
}

func (m *MockRepository) Delete(id string) error {
	for i, t := range m.teams {
		if fmt.Sprintf("%d", t.ID) == id {
			m.teams = append(m.teams[:i], m.teams[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("team not found")
}