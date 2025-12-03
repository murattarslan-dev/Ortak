package repository

import "ortak/pkg/utils"

type MockRepository struct {
	users  map[string]*utils.User
	tokens map[string]int
}

func NewMockRepository() Repository {
	return &MockRepository{
		users:  make(map[string]*utils.User),
		tokens: make(map[string]int),
	}
}

func (m *MockRepository) GetUserByEmail(email string) *utils.User {
	return m.users[email]
}

func (m *MockRepository) CreateUser(username, email, hashedPassword string) *utils.User {
	user := &utils.User{
		ID:       len(m.users) + 1,
		Username: username,
		Email:    email,
		Password: hashedPassword,
	}
	m.users[email] = user
	return user
}

func (m *MockRepository) AddToken(token string, userID int) {
	m.tokens[token] = userID
}

func (m *MockRepository) IsTokenValid(token string) (int, bool) {
	userID, exists := m.tokens[token]
	return userID, exists
}

func (m *MockRepository) RemoveToken(token string) {
	delete(m.tokens, token)
}