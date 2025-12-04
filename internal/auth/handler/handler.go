package handler

import (
	"ortak/internal/auth"
	"ortak/internal/auth/service"
	"ortak/pkg/response"
	"strings"

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

func (h *Handler) Register(c *gin.Context) {
	var req auth.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.SetError(c, 400, "Invalid request format: "+err.Error())
		return
	}

	result, err := h.service.Register(req)
	if err != nil {
		response.SetError(c, 500, err.Error())
		return
	}

	response.SetCreated(c, "User registered successfully", result)
}

func (h *Handler) Login(c *gin.Context) {
	var req auth.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.SetError(c, 400, "Invalid request format: "+err.Error())
		return
	}

	result, err := h.service.Login(req)
	if err != nil {
		response.SetError(c, 401, "Invalid credentials")
		return
	}

	response.SetSuccess(c, "Login successful", result)
}

func (h *Handler) Logout(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		response.SetError(c, 400, "Authorization header required")
		return
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")
	err := h.service.Logout(token)
	if err != nil {
		response.SetError(c, 500, err.Error())
		return
	}

	response.SetSuccess(c, "Logged out successfully", nil)
}
