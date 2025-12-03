package repository

import "ortak/pkg/utils"

type Repository interface {
	GetUserByEmail(email string) *utils.User
	CreateUser(username, email, hashedPassword string) *utils.User
	AddToken(token string, userID int)
	IsTokenValid(token string) (int, bool)
	RemoveToken(token string)
}