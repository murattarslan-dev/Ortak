package repository

import "ortak/internal/user"

type MockRepository struct {
	users  map[string]*user.User
	tokens map[string]string
}

func NewMockRepository() Repository {
	return &MockRepository{
		users:  make(map[string]*user.User),
		tokens: make(map[string]string),
	}
}

func (m *MockRepository) GetUserByEmail(email string) *user.User {
	return m.users[email]
}

func (m *MockRepository) CreateUser(username, email, hashedPassword string) *user.User {
	user := &user.User{
		ID:       "2",
		Username: username,
		Email:    email,
		Password: hashedPassword,
	}
	m.users[email] = user
	return user
}

func (m *MockRepository) AddToken(token string, userID string) {
	m.tokens[token] = userID
}

func (m *MockRepository) IsTokenValid(token string) (string, bool) {
	userID, exists := m.tokens[token]
	return userID, exists
}

func (m *MockRepository) RemoveToken(token string) {
	delete(m.tokens, token)
}
