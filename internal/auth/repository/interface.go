package repository

import "ortak/internal/user"

type Repository interface {
	GetUserByEmail(email string) *user.User
	CreateUser(username, email, hashedPassword string) *user.User
	AddToken(token string, userID string)
	IsTokenValid(token string) (string, bool)
	RemoveToken(token string)
}
