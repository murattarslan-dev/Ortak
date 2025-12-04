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