package handler

import (
	"ortak/internal/task"
	"ortak/internal/task/service"
	"ortak/pkg/response"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) GetTasks(c *gin.Context) {
	tasks, err := h.service.GetTasks()
	if err != nil {
		response.SetError(c, 500, "Failed to get tasks")
		return
	}
	response.SetSuccess(c, "Tasks retrieved successfully", tasks)
}

func (h *Handler) CreateTask(c *gin.Context) {
	var req task.CreateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.SetError(c, 400, "Invalid request format: "+err.Error())
		return
	}

	task, err := h.service.CreateTask(req)
	if err != nil {
		response.SetError(c, 500, "Failed to create task")
		return
	}

	response.SetCreated(c, "Task created successfully", task)
}

func (h *Handler) GetTask(c *gin.Context) {
	id := c.Param("id")
	task, err := h.service.GetTaskByID(id)
	if err != nil {
		response.SetError(c, 404, "Task not found")
		return
	}
	response.SetSuccess(c, "Task retrieved successfully", task)
}

func (h *Handler) UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var req task.UpdateTaskRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.SetError(c, 400, "Invalid request format: "+err.Error())
		return
	}

	task, err := h.service.UpdateTask(id, req)
	if err != nil {
		if err.Error() == "task not found" {
			response.SetError(c, 404, "Task not found")
		} else if err.Error() == "invalid status: must be todo, in_progress, or done" {
			response.SetError(c, 400, "Invalid status: must be todo, in_progress, or done")
		} else {
			response.SetError(c, 500, "Failed to update task")
		}
		return
	}

	response.SetSuccess(c, "Task updated successfully", task)
}

func (h *Handler) UpdateTaskStatus(c *gin.Context) {
	id := c.Param("id")
	var req task.UpdateTaskStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.SetError(c, 400, "Invalid request format: "+err.Error())
		return
	}

	task, err := h.service.UpdateTaskStatus(id, req)
	if err != nil {
		if err.Error() == "task not found" {
			response.SetError(c, 404, "Task not found")
		} else if err.Error() == "invalid status: must be todo, in_progress, or done" {
			response.SetError(c, 400, "Invalid status: must be todo, in_progress, or done")
		} else {
			response.SetError(c, 500, "Failed to update task status")
		}
		return
	}

	response.SetSuccess(c, "Task status updated successfully", task)
}

func (h *Handler) DeleteTask(c *gin.Context) {
	id := c.Param("id")
	err := h.service.DeleteTask(id)
	if err != nil {
		if err.Error() == "task not found" {
			response.SetError(c, 404, "Task not found")
		} else {
			response.SetError(c, 500, "Failed to delete task")
		}
		return
	}

	response.SetSuccess(c, "Task deleted successfully", nil)
}