package team

type Team struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	OwnerID     int    `json:"owner_id"`
}

type CreateTeamRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

type UpdateTeamRequest struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type TeamMember struct {
	ID     int    `json:"id"`
	UserID int    `json:"user_id"`
	TeamID int    `json:"team_id"`
	Role   string `json:"role"`
}

type AddMemberRequest struct {
	UserID int    `json:"user_id" binding:"required"`
	Role   string `json:"role" binding:"required"`
}

type UpdateMemberRoleRequest struct {
	Role string `json:"role" binding:"required"`
}

type TeamWithMembers struct {
	ID          int          `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"description"`
	OwnerID     int          `json:"owner_id"`
	Members     []TeamMember `json:"members"`
}