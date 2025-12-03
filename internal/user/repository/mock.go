package repository

import "ortak/internal/user"

type MockRepository struct {
	users []user.User
}

func NewMockRepository() Repository {
	return &MockRepository{
		users: make([]user.User, 0),
	}
}

func (m *MockRepository) GetAll() []user.User {
	return m.users
}

func (m *MockRepository) Create(username, email, hashedPassword string) *user.User {
	user := &user.User{
		ID:       len(m.users) + 1,
		Username: username,
		Email:    email,
	}
	m.users = append(m.users, *user)
	return user
}