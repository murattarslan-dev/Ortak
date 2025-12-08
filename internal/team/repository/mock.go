package repository

import (
	"fmt"
	"ortak/internal/team"
)

type MockRepository struct {
	teams   []team.Team
	members []team.TeamMember
}

func NewMockRepository() Repository {
	return &MockRepository{
		teams:   make([]team.Team, 0),
		members: make([]team.TeamMember, 0),
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

func (m *MockRepository) AddMember(teamID string, userID int, role string) (*team.TeamMember, error) {
	// Check if team exists
	teamExists := false
	for _, t := range m.teams {
		if fmt.Sprintf("%d", t.ID) == teamID {
			teamExists = true
			break
		}
	}
	if !teamExists {
		return nil, fmt.Errorf("team not found")
	}

	// Check if member already exists
	for _, member := range m.members {
		if fmt.Sprintf("%d", member.TeamID) == teamID && member.UserID == userID {
			return &member, nil // Already a member
		}
	}

	// Convert teamID to int
	teamIDInt := 0
	for _, t := range m.teams {
		if fmt.Sprintf("%d", t.ID) == teamID {
			teamIDInt = t.ID
			break
		}
	}

	member := team.TeamMember{
		ID:     len(m.members) + 1,
		UserID: userID,
		TeamID: teamIDInt,
		Role:   role,
	}
	m.members = append(m.members, member)
	return &member, nil
}

func (m *MockRepository) RemoveMember(teamID, userID string) error {
	for i, member := range m.members {
		if fmt.Sprintf("%d", member.TeamID) == teamID && fmt.Sprintf("%d", member.UserID) == userID {
			m.members = append(m.members[:i], m.members[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("member not found")
}

func (m *MockRepository) UpdateMemberRole(teamID, userID, role string) (*team.TeamMember, error) {
	for i, member := range m.members {
		if fmt.Sprintf("%d", member.TeamID) == teamID && fmt.Sprintf("%d", member.UserID) == userID {
			m.members[i].Role = role
			return &m.members[i], nil
		}
	}
	return nil, fmt.Errorf("member not found")
}

func (m *MockRepository) GetTeamMembers(teamID string) []team.TeamMember {
	members := make([]team.TeamMember, 0)
	for _, member := range m.members {
		if fmt.Sprintf("%d", member.TeamID) == teamID {
			members = append(members, member)
		}
	}
	return members
}