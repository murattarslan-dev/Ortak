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

func (h *Handler) GetUser(c *gin.Context) {
	id := c.Param("id")
	user, err := h.service.GetUserByID(id)
	if err != nil {
		response.SetError(c, 404, "User not found")
		return
	}
	response.SetSuccess(c, "User retrieved successfully", user)
}

func (h *Handler) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var req user.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.SetError(c, 400, "Invalid request format: "+err.Error())
		return
	}

	user, err := h.service.UpdateUser(id, req)
	if err != nil {
		if err.Error() == "user not found" {
			response.SetError(c, 404, "User not found")
		} else {
			response.SetError(c, 500, "Failed to update user")
		}
		return
	}

	response.SetSuccess(c, "User updated successfully", user)
}

func (h *Handler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	err := h.service.DeleteUser(id)
	if err != nil {
		if err.Error() == "user not found" {
			response.SetError(c, 404, "User not found")
		} else {
			response.SetError(c, 500, "Failed to delete user")
		}
		return
	}

	response.SetSuccess(c, "User deleted successfully", nil)
}