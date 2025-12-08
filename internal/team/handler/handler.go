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

func (h *Handler) GetTeam(c *gin.Context) {
	id := c.Param("id")
	team, err := h.service.GetTeamByID(id)
	if err != nil {
		response.SetError(c, 404, "Team not found")
		return
	}
	response.SetSuccess(c, "Team retrieved successfully", team)
}

func (h *Handler) UpdateTeam(c *gin.Context) {
	id := c.Param("id")
	var req team.UpdateTeamRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.SetError(c, 400, "Invalid request format: "+err.Error())
		return
	}

	userID := c.GetInt("user_id")
	team, err := h.service.UpdateTeam(id, req, userID)
	if err != nil {
		response.SetError(c, 500, "Failed to update team")
		return
	}

	response.SetSuccess(c, "Team updated successfully", team)
}

func (h *Handler) DeleteTeam(c *gin.Context) {
	id := c.Param("id")
	userID := c.GetInt("user_id")
	err := h.service.DeleteTeam(id, userID)
	if err != nil {
		response.SetError(c, 500, "Failed to delete team")
		return
	}

	response.SetSuccess(c, "Team deleted successfully", nil)
}