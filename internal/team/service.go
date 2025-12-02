package team

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GetTeams() ([]Team, error) {
	// TODO: Get teams from database
	return []Team{
		{ID: 1, Name: "Development Team", Description: "Backend development", OwnerID: 1},
		{ID: 2, Name: "Design Team", Description: "UI/UX design", OwnerID: 2},
	}, nil
}

func (s *Service) CreateTeam(req CreateTeamRequest, ownerID int) (*Team, error) {
	// TODO: Save team to database
	team := &Team{
		ID:          3,
		Name:        req.Name,
		Description: req.Description,
		OwnerID:     ownerID,
	}
	return team, nil
}