package repository

import "ortak/internal/user"

type Repository interface {
	GetAll() []user.User
	GetByID(id string) *user.User
	Create(username, email, hashedPassword string) *user.User
	Update(id, username, email, hashedPassword string) *user.User
	Delete(id string) error
}