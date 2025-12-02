package task

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GetTasks() ([]Task, error) {
	// TODO: Get tasks from database
	return []Task{
		{ID: 1, Title: "Setup API", Description: "Create REST API", Status: "in_progress", AssigneeID: 1, TeamID: 1},
		{ID: 2, Title: "Design UI", Description: "Create user interface", Status: "todo", AssigneeID: 2, TeamID: 2},
	}, nil
}

func (s *Service) CreateTask(req CreateTaskRequest) (*Task, error) {
	// TODO: Save task to database
	task := &Task{
		ID:          3,
		Title:       req.Title,
		Description: req.Description,
		Status:      "todo",
		AssigneeID:  req.AssigneeID,
		TeamID:      req.TeamID,
	}
	return task, nil
}