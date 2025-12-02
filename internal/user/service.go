package user

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GetUsers() ([]User, error) {
	// TODO: Get users from database
	return []User{
		{ID: 1, Username: "user1", Email: "user1@example.com"},
		{ID: 2, Username: "user2", Email: "user2@example.com"},
	}, nil
}

func (s *Service) CreateUser(req CreateUserRequest) (*User, error) {
	// TODO: Save user to database
	user := &User{
		ID:       3,
		Username: req.Username,
		Email:    req.Email,
	}
	return user, nil
}