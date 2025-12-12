package repository

import "ortak/internal/team"

type Repository interface {
	GetAll() []team.Team
	GetByID(id string) *team.Team
	Create(name, description string, ownerID string) *team.Team
	Update(id, name, description string) *team.Team
	Delete(id string) error
	AddMember(teamID string, userID string, role string) (*team.TeamMember, error)
	RemoveMember(teamID, userID string) error
	UpdateMemberRole(teamID, userID, role string) (*team.TeamMember, error)
	GetTeamMembers(teamID string) []team.TeamMember
	GetMember(teamID, userID string) (*team.TeamMember, error)
}