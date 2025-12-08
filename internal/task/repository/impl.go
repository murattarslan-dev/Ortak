package repository

import (
	"ortak/internal/task"
	"ortak/pkg/utils"
)

type RepositoryImpl struct {
	storage *utils.MemoryStorage
}

func NewRepositoryImpl() Repository {
	return &RepositoryImpl{
		storage: utils.GetMemoryStorage(),
	}
}

func (r *RepositoryImpl) GetAll() []task.Task {
	storageTasks := r.storage.GetAllTasks()
	tasks := make([]task.Task, len(storageTasks))
	for i, t := range storageTasks {
		tasks[i] = task.Task{
			ID:          t.ID,
			Title:       t.Title,
			Description: t.Description,
			Status:      t.Status,
			AssigneeID:  t.AssigneeID,
			TeamID:      t.TeamID,
		}
	}
	return tasks
}

func (r *RepositoryImpl) Create(title, description string, assigneeID, teamID int) *task.Task {
	storageTask := r.storage.CreateTask(title, description, assigneeID, teamID)
	return &task.Task{
		ID:          storageTask.ID,
		Title:       storageTask.Title,
		Description: storageTask.Description,
		Status:      storageTask.Status,
		AssigneeID:  storageTask.AssigneeID,
		TeamID:      storageTask.TeamID,
	}
}

func (r *RepositoryImpl) GetByID(id string) *task.Task {
	storageTask := r.storage.GetTaskByID(id)
	if storageTask == nil {
		return nil
	}
	return &task.Task{
		ID:          storageTask.ID,
		Title:       storageTask.Title,
		Description: storageTask.Description,
		Status:      storageTask.Status,
		AssigneeID:  storageTask.AssigneeID,
		TeamID:      storageTask.TeamID,
	}
}

func (r *RepositoryImpl) Update(id, title, description, status string, assigneeID int) *task.Task {
	storageTask := r.storage.UpdateTask(id, title, description, status, assigneeID)
	if storageTask == nil {
		return nil
	}
	return &task.Task{
		ID:          storageTask.ID,
		Title:       storageTask.Title,
		Description: storageTask.Description,
		Status:      storageTask.Status,
		AssigneeID:  storageTask.AssigneeID,
		TeamID:      storageTask.TeamID,
	}
}

func (r *RepositoryImpl) Delete(id string) error {
	return r.storage.DeleteTask(id)
}