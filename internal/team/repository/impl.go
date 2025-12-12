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
	rows, err := r.db.Query("SELECT id, name, description, owner_id FROM teams")
	if err != nil {
		return []team.Team{}
	}
	defer rows.Close()

	var teams []team.Team
	for rows.Next() {
		var t team.Team
		rows.Scan(&t.ID, &t.Name, &t.Description, &t.OwnerID)
		teams = append(teams, t)
	}
	return teams
}

func (r *RepositoryImpl) Create(name, description string, ownerID int) *team.Team {
	result, err := r.db.Exec("INSERT INTO teams (name, description, owner_id) VALUES (?, ?, ?)", name, description, ownerID)
	if err != nil {
		return nil
	}
	id, _ := result.LastInsertId()
	return &team.Team{ID: int(id), Name: name, Description: description, OwnerID: ownerID}
}

func (r *RepositoryImpl) GetByID(id string) *team.Team {
	var t team.Team
	err := r.db.QueryRow("SELECT id, name, description, owner_id FROM teams WHERE id = ?", id).
		Scan(&t.ID, &t.Name, &t.Description, &t.OwnerID)
	if err != nil {
		return nil
	}
	return &t
}

func (r *RepositoryImpl) Update(id, name, description string) *team.Team {
	_, err := r.db.Exec("UPDATE teams SET name = ?, description = ? WHERE id = ?", name, description, id)
	if err != nil {
		return nil
	}
	return r.GetByID(id)
}

func (r *RepositoryImpl) Delete(id string) error {
	_, err := r.db.Exec("DELETE FROM teams WHERE id = ?", id)
	return err
}

func (r *RepositoryImpl) AddMember(teamID string, userID int, role string) (*team.TeamMember, error) {
	result, err := r.db.Exec("INSERT INTO team_members (user_id, team_id, role) VALUES (?, ?, ?)", userID, teamID, role)
	if err != nil {
		return nil, err
	}
	id, _ := result.LastInsertId()
	teamIDInt := 0
	r.db.QueryRow("SELECT id FROM teams WHERE id = ?", teamID).Scan(&teamIDInt)
	return &team.TeamMember{ID: int(id), UserID: userID, TeamID: teamIDInt, Role: role}, nil
}

func (r *RepositoryImpl) RemoveMember(teamID, userID string) error {
	_, err := r.db.Exec("DELETE FROM team_members WHERE team_id = ? AND user_id = ?", teamID, userID)
	return err
}

func (r *RepositoryImpl) UpdateMemberRole(teamID, userID, role string) (*team.TeamMember, error) {
	_, err := r.db.Exec("UPDATE team_members SET role = ? WHERE team_id = ? AND user_id = ?", role, teamID, userID)
	if err != nil {
		return nil, err
	}

	var member team.TeamMember
	err = r.db.QueryRow("SELECT id, user_id, team_id, role FROM team_members WHERE team_id = ? AND user_id = ?", teamID, userID).
		Scan(&member.ID, &member.UserID, &member.TeamID, &member.Role)
	if err != nil {
		return nil, err
	}
	return &member, nil
}

func (r *RepositoryImpl) GetTeamMembers(teamID string) []team.TeamMember {
	rows, err := r.db.Query("SELECT id, user_id, team_id, role FROM team_members WHERE team_id = ?", teamID)
	if err != nil {
		return []team.TeamMember{}
	}
	defer rows.Close()

	var members []team.TeamMember
	for rows.Next() {
		var m team.TeamMember
		rows.Scan(&m.ID, &m.UserID, &m.TeamID, &m.Role)
		members = append(members, m)
	}
	return members
}
