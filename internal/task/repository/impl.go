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
		commentCount := r.storage.GetTaskCommentCount(t.ID)
		tasks[i] = task.Task{
			ID:           t.ID,
			Title:        t.Title,
			Description:  t.Description,
			Status:       t.Status,
			AssigneeID:   t.AssigneeID,
			TeamID:       t.TeamID,
			Tags:         r.stringToTags(t.Tags),
			CommentCount: commentCount,
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
	commentCount := r.storage.GetTaskCommentCount(storageTask.ID)
	return &task.Task{
		ID:           storageTask.ID,
		Title:        storageTask.Title,
		Description:  storageTask.Description,
		Status:       storageTask.Status,
		AssigneeID:   storageTask.AssigneeID,
		TeamID:       storageTask.TeamID,
		Tags:         r.stringToTags(storageTask.Tags),
		CommentCount: commentCount,
	}
}

func (r *RepositoryImpl) GetByID(id string) *task.Task {
	storageTask := r.storage.GetTaskByID(id)
	if storageTask == nil {
		return nil
	}
	commentCount := r.storage.GetTaskCommentCount(storageTask.ID)
	return &task.Task{
		ID:           storageTask.ID,
		Title:        storageTask.Title,
		Description:  storageTask.Description,
		Status:       storageTask.Status,
		AssigneeID:   storageTask.AssigneeID,
		TeamID:       storageTask.TeamID,
		Tags:         r.stringToTags(storageTask.Tags),
		CommentCount: commentCount,
	}
}

func (r *RepositoryImpl) GetByIDWithComments(id string) *task.Task {
	storageTask := r.storage.GetTaskByID(id)
	if storageTask == nil {
		return nil
	}

	storageComments := r.storage.GetTaskComments(storageTask.ID)
	comments := make([]task.TaskComment, len(storageComments))
	for i, c := range storageComments {
		user := r.storage.GetUserByIDInt(c.UserID)
		comments[i] = task.TaskComment{
			ID:        c.ID,
			TaskID:    c.TaskID,
			Comment:   c.Comment,
			CreatedAt: c.CreatedAt,
			User: task.CommentUser{
				ID:       user.ID,
				Username: user.Username,
				Email:    user.Email,
			},
		}
	}

	return &task.Task{
		ID:           storageTask.ID,
		Title:        storageTask.Title,
		Description:  storageTask.Description,
		Status:       storageTask.Status,
		AssigneeID:   storageTask.AssigneeID,
		TeamID:       storageTask.TeamID,
		Tags:         r.stringToTags(storageTask.Tags),
		CommentCount: len(comments),
		Comments:     comments,
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
	commentCount := r.storage.GetTaskCommentCount(storageTask.ID)
	return &task.Task{
		ID:           storageTask.ID,
		Title:        storageTask.Title,
		Description:  storageTask.Description,
		Status:       storageTask.Status,
		AssigneeID:   storageTask.AssigneeID,
		TeamID:       storageTask.TeamID,
		Tags:         r.stringToTags(storageTask.Tags),
		CommentCount: commentCount,
	}
}

func (r *RepositoryImpl) UpdateStatus(id, status string) *task.Task {
	storageTask := r.storage.UpdateTask(id, "", "", status, "", 0)
	if storageTask == nil {
		return nil
	}
	commentCount := r.storage.GetTaskCommentCount(storageTask.ID)
	return &task.Task{
		ID:           storageTask.ID,
		Title:        storageTask.Title,
		Description:  storageTask.Description,
		Status:       storageTask.Status,
		AssigneeID:   storageTask.AssigneeID,
		TeamID:       storageTask.TeamID,
		Tags:         r.stringToTags(storageTask.Tags),
		CommentCount: commentCount,
	}
}

func (r *RepositoryImpl) AddComment(taskID, userID int, comment, createdAt string) *task.TaskComment {
	storageComment := r.storage.AddTaskComment(taskID, userID, comment, createdAt)
	user := r.storage.GetUserByIDInt(userID)
	return &task.TaskComment{
		ID:        storageComment.ID,
		TaskID:    storageComment.TaskID,
		Comment:   storageComment.Comment,
		CreatedAt: storageComment.CreatedAt,
		User: task.CommentUser{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
		},
	}
}

func (r *RepositoryImpl) Delete(id string) error {
	return r.storage.DeleteTask(id)
}
