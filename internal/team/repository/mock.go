package repository

import (
	"fmt"
	"ortak/internal/team"
	"time"
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

func (m *MockRepository) Create(name, description string, ownerID string) *team.Team {
	now := time.Now()
	team := &team.Team{
		ID:          "2",
		Name:        name,
		Description: description,
		OwnerID:     ownerID,
		IsActive:    true,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	m.teams = append(m.teams, *team)
	return team
}

func (m *MockRepository) GetByID(id string) *team.Team {
	for i, t := range m.teams {
		if t.ID == id {
			// Add members to team
			teamCopy := t
			for _, member := range m.members {
				if member.TeamID == id && member.IsActive {
					teamCopy.Members = append(teamCopy.Members, member)
				}
			}
			m.teams[i] = teamCopy
			return &teamCopy
		}
	}
	return nil
}

func (m *MockRepository) Update(id, name, description string) *team.Team {
	for i, t := range m.teams {
		if t.ID == id {
			if name != "" {
				m.teams[i].Name = name
			}
			if description != "" {
				m.teams[i].Description = description
			}
			m.teams[i].UpdatedAt = time.Now()
			return &m.teams[i]
		}
	}
	return nil
}

func (m *MockRepository) Delete(id string) error {
	for i, t := range m.teams {
		if t.ID == id {
			m.teams[i].IsActive = false
			return nil
		}
	}
	return fmt.Errorf("team not found")
}

func (m *MockRepository) AddMember(teamID string, userID string, role string) (*team.TeamMember, error) {
	// Check if team exists
	teamExists := false
	for _, t := range m.teams {
		if t.ID == teamID {
			teamExists = true
			break
		}
	}
	if !teamExists {
		return nil, fmt.Errorf("team not found")
	}

	// Check if member already exists
	for _, member := range m.members {
		if member.TeamID == teamID && member.UserID == userID {
			return &member, nil
		}
	}

	member := team.TeamMember{
		UserID:   userID,
		TeamID:   teamID,
		Role:     role,
		IsActive: true,
		JoinedAt: time.Now(),
	}
	m.members = append(m.members, member)
	return &member, nil
}

func (m *MockRepository) RemoveMember(teamID, userID string) error {
	for i, member := range m.members {
		if member.TeamID == teamID && member.UserID == userID {
			now := time.Now()
			m.members[i].IsActive = false
			m.members[i].LeftAt = &now
			return nil
		}
	}
	return fmt.Errorf("member not found")
}

func (m *MockRepository) UpdateMemberRole(teamID, userID, role string) (*team.TeamMember, error) {
	for i, member := range m.members {
		if member.TeamID == teamID && member.UserID == userID {
			m.members[i].Role = role
			return &m.members[i], nil
		}
	}
	return nil, fmt.Errorf("member not found")
}

func (m *MockRepository) GetTeamMembers(teamID string) []team.TeamMember {
	members := make([]team.TeamMember, 0)
	for _, member := range m.members {
		if member.TeamID == teamID && member.IsActive {
			members = append(members, member)
		}
	}
	return members
}

func (m *MockRepository) GetMember(teamID, userID string) (*team.TeamMember, error) {
	for _, member := range m.members {
		if member.TeamID == teamID && member.UserID == userID {
			return &member, nil
		}
	}
	return nil, fmt.Errorf("member not found")
}
