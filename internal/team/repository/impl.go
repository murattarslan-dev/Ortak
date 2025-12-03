package repository

import (
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