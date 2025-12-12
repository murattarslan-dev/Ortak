package repository

import (
	"database/sql"
	"ortak/internal/user"
)

type RepositoryImpl struct {
	db *sql.DB
}

func NewRepositoryImpl(database *sql.DB) Repository {
	return &RepositoryImpl{
		db: database,
	}
}

func (r *RepositoryImpl) GetUserByEmail(email string) *user.User {
	var user user.User
	err := r.db.QueryRow("SELECT id, username, email, password FROM users WHERE email = $1", email).
		Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return nil
	}
	return &user
}

func (r *RepositoryImpl) CreateUser(username, email, hashedPassword string) *user.User {
	var user user.User
	err := r.db.QueryRow("INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id, username, email",
		username, email, hashedPassword).Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		return nil
	}
	return &user
}

func (r *RepositoryImpl) AddToken(token string, userID string) {
	r.db.Exec("INSERT INTO tokens (token, user_id) VALUES ($1, $2)", token, userID)
}

func (r *RepositoryImpl) IsTokenValid(token string) (string, bool) {
	var userID string
	err := r.db.QueryRow("SELECT user_id FROM tokens WHERE token = $1", token).Scan(&userID)
	return userID, err == nil
}

func (r *RepositoryImpl) RemoveToken(token string) {
	r.db.Exec("DELETE FROM tokens WHERE token = $1", token)
}
