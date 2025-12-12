package service

import (
	"ortak/internal/task"
	"ortak/internal/task/repository"
	"testing"
	"time"
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
		Priority:    "high",
		DueDate:     &time.Time{},
		Tags:        []string{"backend", "api"},
	}

	createdTask, err := service.CreateTask(req, "user-123")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if createdTask.Title != "Test Task" {
		t.Errorf("Expected title Test Task, got %s", createdTask.Title)
	}

	if createdTask.Status != "todo" {
		t.Errorf("Expected status todo, got %s", createdTask.Status)
	}

	if createdTask.Priority != "high" {
		t.Errorf("Expected priority high, got %s", createdTask.Priority)
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
		Tags:        []string{"test"},
	}
	createdTask, _ := service.CreateTask(createReq, "user-123")

	// Test valid status update
	updateReq := task.UpdateTaskStatusRequest{
		Status: "in_progress",
	}

	updatedTask, err := service.UpdateTaskStatus(createdTask.ID, updateReq)
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

	_, err = service.UpdateTaskStatus(createdTask.ID, invalidReq)
	if err == nil {
		t.Error("Expected error for invalid status, got nil")
	}

	// Test non-existent task
	_, err = service.UpdateTaskStatus("non-existent-id", updateReq)
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
		Tags:        []string{"test"},
	}
	createdTask, _ := service.CreateTask(createReq, "user-123")

	// Test valid comment
	commentReq := task.AddCommentRequest{
		Comment: "This is a test comment",
	}

	comment, err := service.AddComment(createdTask.ID, "user-456", commentReq)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if comment.Comment != "This is a test comment" {
		t.Errorf("Expected comment 'This is a test comment', got %s", comment.Comment)
	}

	if comment.TaskID != createdTask.ID {
		t.Errorf("Expected task ID %s, got %s", createdTask.ID, comment.TaskID)
	}

	// Test non-existent task
	_, err = service.AddComment("non-existent-id", "user-456", commentReq)
	if err == nil {
		t.Error("Expected error for non-existent task, got nil")
	}
}

func TestService_AddAssignment(t *testing.T) {
	repo := repository.NewMockRepository()
	service := NewService(repo)

	// Create a task first
	createReq := task.CreateTaskRequest{
		Title:       "Test Task",
		Description: "Test Description",
		Tags:        []string{"test"},
	}
	createdTask, _ := service.CreateTask(createReq, "user-123")

	// Test user assignment
	userAssignReq := task.AddAssignmentRequest{
		AssignType: "user",
		AssignID:   "user-456",
	}

	assignment, err := service.AddAssignment(createdTask.ID, userAssignReq)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if assignment.AssignType != "user" {
		t.Errorf("Expected assign_type 'user', got %s", assignment.AssignType)
	}

	if assignment.AssignID != "user-456" {
		t.Errorf("Expected assign_id 'user-456', got %s", assignment.AssignID)
	}

	// Test team assignment
	teamAssignReq := task.AddAssignmentRequest{
		AssignType: "team",
		AssignID:   "team-789",
	}

	teamAssignment, err := service.AddAssignment(createdTask.ID, teamAssignReq)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if teamAssignment.AssignType != "team" {
		t.Errorf("Expected assign_type 'team', got %s", teamAssignment.AssignType)
	}

	// Test invalid assign_type
	invalidReq := task.AddAssignmentRequest{
		AssignType: "invalid",
		AssignID:   "some-id",
	}

	_, err = service.AddAssignment(createdTask.ID, invalidReq)
	if err == nil {
		t.Error("Expected error for invalid assign_type, got nil")
	}

	// Test non-existent task
	_, err = service.AddAssignment("non-existent-id", userAssignReq)
	if err == nil {
		t.Error("Expected error for non-existent task, got nil")
	}
}
