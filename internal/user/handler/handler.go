package handler

import (
	"ortak/internal/user"
	"ortak/internal/user/service"
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

func (h *Handler) GetUsers(c *gin.Context) {
	users, err := h.service.GetUsers()
	if err != nil {
		response.SetError(c, 500, "Failed to get users")
		return
	}
	response.SetSuccess(c, "Users retrieved successfully", users)
}

func (h *Handler) CreateUser(c *gin.Context) {
	var req user.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.SetError(c, 400, "Invalid request format: "+err.Error())
		return
	}

	user, err := h.service.CreateUser(req)
	if err != nil {
		response.SetError(c, 500, "Failed to create user")
		return
	}

	response.SetCreated(c, "User created successfully", user)
}