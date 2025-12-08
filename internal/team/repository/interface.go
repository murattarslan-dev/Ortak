package repository

import "ortak/internal/team"

type Repository interface {
	GetAll() []team.Team
	GetByID(id string) *team.Team
	Create(name, description string, ownerID int) *team.Team
	Update(id, name, description string) *team.Team
	Delete(id string) error
}