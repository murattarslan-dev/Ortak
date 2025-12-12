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

func (r *RepositoryImpl) GetAll() []user.User {
	rows, err := r.db.Query("SELECT id, username, email, first_name, last_name, phone, department, position, company, is_active FROM users WHERE is_active = true")
	if err != nil {
		return []user.User{}
	}
	defer rows.Close()

	var users []user.User
	for rows.Next() {
		var u user.User
		rows.Scan(&u.ID, &u.Username, &u.Email, &u.FirstName, &u.LastName, &u.Phone, &u.Department, &u.Position, &u.Company, &u.IsActive)
		users = append(users, u)
	}
	return users
}

func (r *RepositoryImpl) Create(username, email, hashedPassword string) *user.User {
	var u user.User
	err := r.db.QueryRow("INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id, username, email",
		username, email, hashedPassword).Scan(&u.ID, &u.Username, &u.Email)
	if err != nil {
		return nil
	}
	return &u
}

func (r *RepositoryImpl) GetByID(id string) *user.User {
	var u user.User
	err := r.db.QueryRow("SELECT id, username, email, first_name, last_name, phone, department, position, company, is_active FROM users WHERE id = $1", id).
		Scan(&u.ID, &u.Username, &u.Email, &u.FirstName, &u.LastName, &u.Phone, &u.Department, &u.Position, &u.Company, &u.IsActive)
	if err != nil {
		return nil
	}
	return &u
}

func (r *RepositoryImpl) Update(id, username, email, hashedPassword string) *user.User {
	_, err := r.db.Exec("UPDATE users SET username = $1, email = $2, password = $3 WHERE id = $4",
		username, email, hashedPassword, id)
	if err != nil {
		return nil
	}
	return r.GetByID(id)
}

func (r *RepositoryImpl) Delete(id string) error {
	_, err := r.db.Exec("UPDATE users SET is_active = false WHERE id = $1", id)
	return err
}
