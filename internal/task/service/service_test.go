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