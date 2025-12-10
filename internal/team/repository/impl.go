package repository

import (
	"fmt"
	"ortak/internal/team"
	"ortak/pkg/utils"
)

type RepositoryImpl struct {
	storage *utils.MemoryStorage
}

func NewRepositoryImpl() Repository {
	return &RepositoryImpl{
		storage: utils.GetMemoryStorage(),
	}
}

func (r *RepositoryImpl) GetAll() []team.Team {
	storageTeams := r.storage.GetAllTeams()
	teams := make([]team.Team, len(storageTeams))
	for i, t := range storageTeams {
		teams[i] = team.Team{
			ID:          t.ID,
			Name:        t.Name,
			Description: t.Description,
			OwnerID:     t.OwnerID,
		}
	}
	return teams
}

func (r *RepositoryImpl) Create(name, description string, ownerID int) *team.Team {
	storageTeam := r.storage.CreateTeam(name, description, ownerID)
	return &team.Team{
		ID:          storageTeam.ID,
		Name:        storageTeam.Name,
		Description: storageTeam.Description,
		OwnerID:     storageTeam.OwnerID,
	}
}

func (r *RepositoryImpl) GetByID(id string) *team.Team {
	storageTeam := r.storage.GetTeamByID(id)
	if storageTeam == nil {
		return nil
	}
	return &team.Team{
		ID:          storageTeam.ID,
		Name:        storageTeam.Name,
		Description: storageTeam.Description,
		OwnerID:     storageTeam.OwnerID,
	}
}

func (r *RepositoryImpl) Update(id, name, description string) *team.Team {
	storageTeam := r.storage.UpdateTeam(id, name, description)
	if storageTeam == nil {
		return nil
	}
	return &team.Team{
		ID:          storageTeam.ID,
		Name:        storageTeam.Name,
		Description: storageTeam.Description,
		OwnerID:     storageTeam.OwnerID,
	}
}

func (r *RepositoryImpl) Delete(id string) error {
	return r.storage.DeleteTeam(id)
}

func (r *RepositoryImpl) AddMember(teamID string, userID int, role string) (*team.TeamMember, error) {
	// Convert teamID to int
	teamIDInt := 0
	for i := range r.storage.GetAllTeams() {
		if fmt.Sprintf("%d", i+1) == teamID {
			teamIDInt = i + 1
			break
		}
	}
	if teamIDInt == 0 {
		return nil, fmt.Errorf("team not found")
	}

	storageMember := r.storage.AddTeamMember(userID, teamIDInt, role)
	return &team.TeamMember{
		ID:     storageMember.ID,
		UserID: storageMember.UserID,
		TeamID: storageMember.TeamID,
		Role:   storageMember.Role,
	}, nil
}

func (r *RepositoryImpl) RemoveMember(teamID, userID string) error {
	// Convert IDs to int
	teamIDInt := 0
	userIDInt := 0
	
	for i := range r.storage.GetAllTeams() {
		if fmt.Sprintf("%d", i+1) == teamID {
			teamIDInt = i + 1
			break
		}
	}
	
	for i := range r.storage.GetAllUsers() {
		if fmt.Sprintf("%d", i+1) == userID {
			userIDInt = i + 1
			break
		}
	}

	return r.storage.RemoveTeamMember(userIDInt, teamIDInt)
}

func (r *RepositoryImpl) UpdateMemberRole(teamID, userID, role string) (*team.TeamMember, error) {
	// Convert IDs to int
	teamIDInt := 0
	userIDInt := 0
	
	for i := range r.storage.GetAllTeams() {
		if fmt.Sprintf("%d", i+1) == teamID {
			teamIDInt = i + 1
			break
		}
	}
	
	for i := range r.storage.GetAllUsers() {
		if fmt.Sprintf("%d", i+1) == userID {
			userIDInt = i + 1
			break
		}
	}

	storageMember := r.storage.UpdateMemberRole(userIDInt, teamIDInt, role)
	if storageMember == nil {
		return nil, fmt.Errorf("member not found")
	}
	
	return &team.TeamMember{
		ID:     storageMember.ID,
		UserID: storageMember.UserID,
		TeamID: storageMember.TeamID,
		Role:   storageMember.Role,
	}, nil
}

func (r *RepositoryImpl) GetTeamMembers(teamID string) []team.TeamMember {
	// Convert teamID to int
	teamIDInt := 0
	for i := range r.storage.GetAllTeams() {
		if fmt.Sprintf("%d", i+1) == teamID {
			teamIDInt = i + 1
			break
		}
	}

	storageMembers := r.storage.GetTeamMembers(teamIDInt)
	members := make([]team.TeamMember, len(storageMembers))
	for i, m := range storageMembers {
		members[i] = team.TeamMember{
			ID:     m.ID,
			UserID: m.UserID,
			TeamID: m.TeamID,
			Role:   m.Role,
		}
	}
	return members
}