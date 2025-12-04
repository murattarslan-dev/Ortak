package repository

import (
	"ortak/internal/user"
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

func (r *RepositoryImpl) GetAll() []user.User {
	storageUsers := r.storage.GetAllUsers()
	users := make([]user.User, len(storageUsers))
	for i, u := range storageUsers {
		users[i] = user.User{
			ID:       u.ID,
			Username: u.Username,
			Email:    u.Email,
		}
	}
	return users
}

func (r *RepositoryImpl) Create(username, email, hashedPassword string) *user.User {
	storageUser := r.storage.CreateUser(username, email, hashedPassword)
	return &user.User{
		ID:       storageUser.ID,
		Username: storageUser.Username,
		Email:    storageUser.Email,
	}
}

func (r *RepositoryImpl) GetByID(id string) *user.User {
	storageUser := r.storage.GetUserByID(id)
	if storageUser == nil {
		return nil
	}
	return &user.User{
		ID:       storageUser.ID,
		Username: storageUser.Username,
		Email:    storageUser.Email,
	}
}

func (r *RepositoryImpl) Update(id, username, email, hashedPassword string) *user.User {
	storageUser := r.storage.UpdateUser(id, username, email, hashedPassword)
	if storageUser == nil {
		return nil
	}
	return &user.User{
		ID:       storageUser.ID,
		Username: storageUser.Username,
		Email:    storageUser.Email,
	}
}

func (r *RepositoryImpl) Delete(id string) error {
	return r.storage.DeleteUser(id)
}