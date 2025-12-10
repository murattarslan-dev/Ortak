package repository

import (
	"ortak/internal/task"
	"ortak/pkg/utils"
	"strings"
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
			Tags:        r.stringToTags(t.Tags),
		}
	}
	return tasks
}

func (r *RepositoryImpl) stringToTags(tagStr string) []string {
	if tagStr == "" {
		return []string{}
	}
	return strings.Split(tagStr, "{$^#}")
}

func (r *RepositoryImpl) tagsToString(tags []string) string {
	return strings.Join(tags, "{$^#}")
}

func (r *RepositoryImpl) Create(title, description string, assigneeID, teamID int, tags []string) *task.Task {
	storageTask := r.storage.CreateTask(title, description, assigneeID, teamID)
	if len(tags) > 0 {
		storageTask.Tags = r.tagsToString(tags)
	}
	return &task.Task{
		ID:          storageTask.ID,
		Title:       storageTask.Title,
		Description: storageTask.Description,
		Status:      storageTask.Status,
		AssigneeID:  storageTask.AssigneeID,
		TeamID:      storageTask.TeamID,
		Tags:        r.stringToTags(storageTask.Tags),
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
		Tags:        r.stringToTags(storageTask.Tags),
	}
}

func (r *RepositoryImpl) Update(id, title, description, status string, assigneeID int, tags []string) *task.Task {
	tagsStr := ""
	if len(tags) > 0 {
		tagsStr = r.tagsToString(tags)
	}
	storageTask := r.storage.UpdateTask(id, title, description, status, tagsStr, assigneeID)
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
		Tags:        r.stringToTags(storageTask.Tags),
	}
}

func (r *RepositoryImpl) UpdateStatus(id, status string) *task.Task {
	storageTask := r.storage.UpdateTask(id, "", "", status, "", 0)
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
		Tags:        r.stringToTags(storageTask.Tags),
	}
}

func (r *RepositoryImpl) Delete(id string) error {
	return r.storage.DeleteTask(id)
}
