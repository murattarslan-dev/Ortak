package repository

import "ortak/internal/user"

type Repository interface {
	GetAll() []user.User
	Create(username, email, hashedPassword string) *user.User
}