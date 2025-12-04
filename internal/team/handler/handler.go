package handler

import (
	"ortak/internal/team"
	"ortak/internal/team/service"
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

func (h *Handler) GetTeams(c *gin.Context) {
	teams, err := h.service.GetTeams()
	if err != nil {
		response.SetError(c, 500, "Failed to get teams")
		return
	}
	response.SetSuccess(c, "Teams retrieved successfully", teams)
}

func (h *Handler) CreateTeam(c *gin.Context) {
	var req team.CreateTeamRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.SetError(c, 400, "Invalid request format: "+err.Error())
		return
	}

	userID := c.GetInt("user_id")
	team, err := h.service.CreateTeam(req, userID)
	if err != nil {
		response.SetError(c, 500, "Failed to create team")
		return
	}

	response.SetCreated(c, "Team created successfully", team)
}