package user

import "time"

type User struct {
	ID         string    `json:"id"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	Phone      string    `json:"phone"`
	Role       string    `json:"role"`
	Password   string    `json:"-"`
	AvatarURL  string    `json:"avatar_url"`
	Department string    `json:"department"`
	Position   string    `json:"position"`
	Company    string    `json:"company"`
	IsActive   bool      `json:"is_active"`
	LastLogin  time.Time `json:"last_login"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	CreatedBy  string    `json:"created_by"`
	UpdatedBy  string    `json:"updated_by"`
}

type CreateUserRequest struct {
	Username   string `json:"username" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required,min=6"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Phone      string `json:"phone"`
	Department string `json:"department"`
	Position   string `json:"position"`
	Company    string `json:"company"`
}

type UpdateUserRequest struct {
	Username   string `json:"username,omitempty"`
	Email      string `json:"email,omitempty" binding:"omitempty,email"`
	Password   string `json:"password,omitempty" binding:"omitempty,min=6"`
	FirstName  string `json:"first_name,omitempty"`
	LastName   string `json:"last_name,omitempty"`
	Phone      string `json:"phone,omitempty"`
	AvatarURL  string `json:"avatar_url,omitempty"`
	Department string `json:"department,omitempty"`
	Position   string `json:"position,omitempty"`
	Company    string `json:"company,omitempty"`
}

type UserTeam struct {
	TeamID string `json:"team_id"`
	Role   string `json:"role"`
}

type UserWithTeams struct {
	ID         string     `json:"id"`
	Username   string     `json:"username"`
	Email      string     `json:"email"`
	FirstName  string     `json:"first_name"`
	LastName   string     `json:"last_name"`
	Department string     `json:"department"`
	Position   string     `json:"position"`
	Company    string     `json:"company"`
	Teams      []UserTeam `json:"teams"`
}
