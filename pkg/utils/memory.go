package utils

import (
	"sync"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

type Team struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	OwnerID     int    `json:"owner_id"`
}

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	AssigneeID  int    `json:"assignee_id"`
	TeamID      int    `json:"team_id"`
}

type MemoryStorage struct {
	users      map[int]*User
	teams      map[int]*Team
	tasks      map[int]*Task
	tokens     map[string]int
	userEmails map[string]int
	nextUserID int
	nextTeamID int
	nextTaskID int
	mu         sync.RWMutex
}

var instance *MemoryStorage
var once sync.Once

func GetMemoryStorage() *MemoryStorage {
	once.Do(func() {
		instance = &MemoryStorage{
			users:      make(map[int]*User),
			teams:      make(map[int]*Team),
			tasks:      make(map[int]*Task),
			tokens:     make(map[string]int),
			userEmails: make(map[string]int),
			nextUserID: 1,
			nextTeamID: 1,
			nextTaskID: 1,
		}
	})
	return instance
}

func (s *MemoryStorage) CreateUser(username, email, hashedPassword string) *User {
	s.mu.Lock()
	defer s.mu.Unlock()

	user := &User{
		ID:       s.nextUserID,
		Username: username,
		Email:    email,
		Password: hashedPassword,
	}
	s.users[s.nextUserID] = user
	s.userEmails[email] = s.nextUserID
	s.nextUserID++
	return user
}

func (s *MemoryStorage) GetUserByEmail(email string) *User {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if userID, exists := s.userEmails[email]; exists {
		return s.users[userID]
	}
	return nil
}

func (s *MemoryStorage) GetUserByID(id int) *User {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.users[id]
}

func (s *MemoryStorage) GetAllUsers() []*User {
	s.mu.RLock()
	defer s.mu.RUnlock()

	users := make([]*User, 0, len(s.users))
	for _, user := range s.users {
		users = append(users, user)
	}
	return users
}

func (s *MemoryStorage) AddToken(token string, userID int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.tokens[token] = userID
}

func (s *MemoryStorage) IsTokenValid(token string) (int, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	userID, exists := s.tokens[token]
	return userID, exists
}

func (s *MemoryStorage) RemoveToken(token string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.tokens, token)
}

func (s *MemoryStorage) CreateTeam(name, description string, ownerID int) *Team {
	s.mu.Lock()
	defer s.mu.Unlock()

	team := &Team{
		ID:          s.nextTeamID,
		Name:        name,
		Description: description,
		OwnerID:     ownerID,
	}
	s.teams[s.nextTeamID] = team
	s.nextTeamID++
	return team
}

func (s *MemoryStorage) GetAllTeams() []*Team {
	s.mu.RLock()
	defer s.mu.RUnlock()

	teams := make([]*Team, 0, len(s.teams))
	for _, team := range s.teams {
		teams = append(teams, team)
	}
	return teams
}

func (s *MemoryStorage) CreateTask(title, description string, assigneeID, teamID int) *Task {
	s.mu.Lock()
	defer s.mu.Unlock()

	task := &Task{
		ID:          s.nextTaskID,
		Title:       title,
		Description: description,
		Status:      "todo",
		AssigneeID:  assigneeID,
		TeamID:      teamID,
	}
	s.tasks[s.nextTaskID] = task
	s.nextTaskID++
	return task
}

func (s *MemoryStorage) GetAllTasks() []*Task {
	s.mu.RLock()
	defer s.mu.RUnlock()

	tasks := make([]*Task, 0, len(s.tasks))
	for _, task := range s.tasks {
		tasks = append(tasks, task)
	}
	return tasks
}