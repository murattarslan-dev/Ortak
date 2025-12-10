package service

import (
	"ortak/internal/task"
	"ortak/internal/task/repository"
	"testing"
)

func TestService_GetTasks(t *testing.T) {
	repo := repository.NewMockRepository()
	service := NewService(repo)

	tasks, err := service.GetTasks()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if len(tasks) != 0 {
		t.Errorf("Expected 0 tasks, got %d", len(tasks))
	}
}

func TestService_CreateTask(t *testing.T) {
	repo := repository.NewMockRepository()
	service := NewService(repo)

	req := task.CreateTaskRequest{
		Title:       "Test Task",
		Description: "Test Description",
		AssigneeID:  1,
		TeamID:      1,
		Tags:        []string{"backend", "api"},
	}

	createdTask, err := service.CreateTask(req)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if createdTask.Title != "Test Task" {
		t.Errorf("Expected title Test Task, got %s", createdTask.Title)
	}

	if createdTask.Status != "todo" {
		t.Errorf("Expected status todo, got %s", createdTask.Status)
	}

	if createdTask.AssigneeID != 1 {
		t.Errorf("Expected assignee ID 1, got %d", createdTask.AssigneeID)
	}

	// Verify task was added
	tasks, _ := service.GetTasks()
	if len(tasks) != 1 {
		t.Errorf("Expected 1 task, got %d", len(tasks))
	}
}

func TestService_UpdateTaskStatus(t *testing.T) {
	repo := repository.NewMockRepository()
	service := NewService(repo)

	// Create a task first
	createReq := task.CreateTaskRequest{
		Title:       "Test Task",
		Description: "Test Description",
		AssigneeID:  1,
		TeamID:      1,
		Tags:        []string{"test"},
	}
	_, _ = service.CreateTask(createReq)

	// Test valid status update
	updateReq := task.UpdateTaskStatusRequest{
		Status: "in_progress",
	}

	updatedTask, err := service.UpdateTaskStatus("1", updateReq)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if updatedTask.Status != "in_progress" {
		t.Errorf("Expected status in_progress, got %s", updatedTask.Status)
	}

	// Test invalid status
	invalidReq := task.UpdateTaskStatusRequest{
		Status: "invalid_status",
	}

	_, err = service.UpdateTaskStatus("1", invalidReq)
	if err == nil {
		t.Error("Expected error for invalid status, got nil")
	}

	// Test non-existent task
	_, err = service.UpdateTaskStatus("999", updateReq)
	if err == nil {
		t.Error("Expected error for non-existent task, got nil")
	}
}

func TestService_AddComment(t *testing.T) {
	repo := repository.NewMockRepository()
	service := NewService(repo)

	// Create a task first
	createReq := task.CreateTaskRequest{
		Title:       "Test Task",
		Description: "Test Description",
		AssigneeID:  1,
		TeamID:      1,
		Tags:        []string{"test"},
	}
	service.CreateTask(createReq)

	// Test valid comment
	commentReq := task.AddCommentRequest{
		Comment: "This is a test comment",
	}

	comment, err := service.AddComment("1", 1, commentReq)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if comment.Comment != "This is a test comment" {
		t.Errorf("Expected comment 'This is a test comment', got %s", comment.Comment)
	}

	if comment.TaskID != 1 {
		t.Errorf("Expected task ID 1, got %d", comment.TaskID)
	}

	// Test non-existent task
	_, err = service.AddComment("999", 1, commentReq)
	if err == nil {
		t.Error("Expected error for non-existent task, got nil")
	}
}
