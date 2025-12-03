package repository

import "ortak/internal/team"

type Repository interface {
	GetAll() []team.Team
	Create(name, description string, ownerID int) *team.Team
}