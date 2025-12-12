package repository

import (
	"database/sql"
	"ortak/internal/team"
)

type RepositoryImpl struct {
	db *sql.DB
}

func NewRepositoryImpl(database *sql.DB) Repository {
	return &RepositoryImpl{
		db: database,
	}
}

func (r *RepositoryImpl) GetAll() []team.Team {
	rows, err := r.db.Query("SELECT id, name, description, owner_id, is_active, created_at, updated_at FROM teams WHERE is_active = true")
	if err != nil {
		return []team.Team{}
	}
	defer rows.Close()

	var teams []team.Team
	for rows.Next() {
		var t team.Team
		rows.Scan(&t.ID, &t.Name, &t.Description, &t.OwnerID, &t.IsActive, &t.CreatedAt, &t.UpdatedAt)
		teams = append(teams, t)
	}
	return teams
}

func (r *RepositoryImpl) Create(name, description string, ownerID string) *team.Team {
	var teamID string
	err := r.db.QueryRow("INSERT INTO teams (name, description, owner_id) VALUES ($1, $2, $3) RETURNING id, created_at, updated_at", name, description, ownerID).Scan(&teamID, new(interface{}), new(interface{}))
	if err != nil {
		return nil
	}
	return r.GetByID(teamID)
}

func (r *RepositoryImpl) GetByID(id string) *team.Team {
	rows, err := r.db.Query(`
		SELECT 
			t.id, t.name, t.description, t.owner_id, t.is_active, t.created_at, t.updated_at,
			tm.user_id, tm.role, tm.is_active, tm.joined_at, tm.left_at, tm.invited_by
		FROM teams t
		LEFT JOIN team_members tm ON t.id = tm.team_id AND tm.is_active = true
		WHERE t.id = $1
		ORDER BY tm.joined_at
	`, id)
	if err != nil {
		return nil
	}
	defer rows.Close()

	var t *team.Team
	for rows.Next() {
		var (
			userID, role sql.NullString
			memberActive sql.NullBool
			joinedAt     sql.NullTime
			leftAt       sql.NullTime
			invitedBy    sql.NullString
		)

		if t == nil {
			t = &team.Team{}
			rows.Scan(&t.ID, &t.Name, &t.Description, &t.OwnerID, &t.IsActive, &t.CreatedAt, &t.UpdatedAt,
				&userID, &role, &memberActive, &joinedAt, &leftAt, &invitedBy)
		} else {
			rows.Scan(new(interface{}), new(interface{}), new(interface{}), new(interface{}), new(interface{}), new(interface{}), new(interface{}),
				&userID, &role, &memberActive, &joinedAt, &leftAt, &invitedBy)
		}

		// Add member if exists
		if userID.Valid {
			member := team.TeamMember{
				UserID:   userID.String,
				TeamID:   t.ID,
				Role:     role.String,
				IsActive: memberActive.Bool,
				JoinedAt: joinedAt.Time,
			}
			if leftAt.Valid {
				member.LeftAt = &leftAt.Time
			}
			if invitedBy.Valid {
				member.InvitedBy = &invitedBy.String
			}
			t.Members = append(t.Members, member)
		}
	}

	return t
}

func (r *RepositoryImpl) Update(id, name, description string) *team.Team {
	_, err := r.db.Exec("UPDATE teams SET name = $1, description = $2 WHERE id = $3", name, description, id)
	if err != nil {
		return nil
	}
	return r.GetByID(id)
}

func (r *RepositoryImpl) Delete(id string) error {
	_, err := r.db.Exec("UPDATE teams SET is_active = false WHERE id = $1", id)
	return err
}

func (r *RepositoryImpl) AddMember(teamID string, userID string, role string) (*team.TeamMember, error) {
	_, err := r.db.Exec("INSERT INTO team_members (user_id, team_id, role) VALUES ($1, $2, $3)", userID, teamID, role)
	if err != nil {
		return nil, err
	}
	return r.GetMember(teamID, userID)
}

func (r *RepositoryImpl) RemoveMember(teamID, userID string) error {
	_, err := r.db.Exec("UPDATE team_members SET is_active = false, left_at = CURRENT_TIMESTAMP WHERE team_id = $1 AND user_id = $2", teamID, userID)
	return err
}

func (r *RepositoryImpl) UpdateMemberRole(teamID, userID, role string) (*team.TeamMember, error) {
	_, err := r.db.Exec("UPDATE team_members SET role = $1 WHERE team_id = $2 AND user_id = $3", role, teamID, userID)
	if err != nil {
		return nil, err
	}
	return r.GetMember(teamID, userID)
}

func (r *RepositoryImpl) GetTeamMembers(teamID string) []team.TeamMember {
	rows, err := r.db.Query("SELECT user_id, team_id, role, is_active, joined_at, left_at, invited_by FROM team_members WHERE team_id = $1 AND is_active = true", teamID)
	if err != nil {
		return []team.TeamMember{}
	}
	defer rows.Close()

	var members []team.TeamMember
	for rows.Next() {
		var m team.TeamMember
		rows.Scan(&m.UserID, &m.TeamID, &m.Role, &m.IsActive, &m.JoinedAt, &m.LeftAt, &m.InvitedBy)
		members = append(members, m)
	}
	return members
}

func (r *RepositoryImpl) GetMember(teamID, userID string) (*team.TeamMember, error) {
	var m team.TeamMember
	err := r.db.QueryRow("SELECT user_id, team_id, role, is_active, joined_at, left_at, invited_by FROM team_members WHERE team_id = $1 AND user_id = $2", teamID, userID).
		Scan(&m.UserID, &m.TeamID, &m.Role, &m.IsActive, &m.JoinedAt, &m.LeftAt, &m.InvitedBy)
	if err != nil {
		return nil, err
	}
	return &m, nil
}
