package repository

import (
	"fmt"
	"ortak/internal/user"
)

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

func (m *MockRepository) GetByID(id string) *user.User {
	for _, u := range m.users {
		if fmt.Sprintf("%d", u.ID) == id {
			return &u
		}
	}
	return nil
}

func (m *MockRepository) Update(id, username, email, hashedPassword string) *user.User {
	for i, u := range m.users {
		if fmt.Sprintf("%d", u.ID) == id {
			if username != "" {
				m.users[i].Username = username
			}
			if email != "" {
				m.users[i].Email = email
			}
			return &m.users[i]
		}
	}
	return nil
}

func (m *MockRepository) Delete(id string) error {
	for i, u := range m.users {
		if fmt.Sprintf("%d", u.ID) == id {
			m.users = append(m.users[:i], m.users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("user not found")
}