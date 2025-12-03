package repository

import "ortak/pkg/utils"

type RepositoryImpl struct {
	storage *utils.MemoryStorage
}

func NewRepositoryImpl() Repository {
	return &RepositoryImpl{
		storage: utils.GetMemoryStorage(),
	}
}

func (r *RepositoryImpl) GetUserByEmail(email string) *utils.User {
	return r.storage.GetUserByEmail(email)
}

func (r *RepositoryImpl) CreateUser(username, email, hashedPassword string) *utils.User {
	return r.storage.CreateUser(username, email, hashedPassword)
}

func (r *RepositoryImpl) AddToken(token string, userID int) {
	r.storage.AddToken(token, userID)
}

func (r *RepositoryImpl) IsTokenValid(token string) (int, bool) {
	return r.storage.IsTokenValid(token)
}