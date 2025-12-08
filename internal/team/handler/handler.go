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
	teamWithMembers, err := h.service.GetTeamWithMembers(id)
	if err != nil {
		response.SetError(c, 404, "Team not found")
		return
	}
	response.SetSuccess(c, "Team retrieved successfully", teamWithMembers)
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

func (h *Handler) AddTeamMember(c *gin.Context) {
	teamID := c.Param("id")
	var req team.AddMemberRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.SetError(c, 400, "Invalid request format: "+err.Error())
		return
	}

	userID := c.GetInt("user_id")
	member, err := h.service.AddTeamMember(teamID, req.UserID, req.Role, userID)
	if err != nil {
		response.SetError(c, 500, "Failed to add team member")
		return
	}

	response.SetCreated(c, "Team member added successfully", member)
}

func (h *Handler) RemoveTeamMember(c *gin.Context) {
	teamID := c.Param("id")
	memberUserID := c.Param("userId")
	userID := c.GetInt("user_id")
	
	err := h.service.RemoveTeamMember(teamID, memberUserID, userID)
	if err != nil {
		response.SetError(c, 500, "Failed to remove team member")
		return
	}

	response.SetSuccess(c, "Team member removed successfully", nil)
}

func (h *Handler) UpdateMemberRole(c *gin.Context) {
	teamID := c.Param("id")
	memberUserID := c.Param("userId")
	var req team.UpdateMemberRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.SetError(c, 400, "Invalid request format: "+err.Error())
		return
	}

	userID := c.GetInt("user_id")
	member, err := h.service.UpdateMemberRole(teamID, memberUserID, req.Role, userID)
	if err != nil {
		response.SetError(c, 500, "Failed to update member role")
		return
	}

	response.SetSuccess(c, "Member role updated successfully", member)
}