-- Team Module Database Schema
-- Updated: December 2024
-- Compatible with PostgreSQL (pgcrypto extension required)

-- Enable pgcrypto extension (PostgreSQL)
CREATE EXTENSION IF NOT EXISTS pgcrypto;

-- Teams table
CREATE TABLE teams (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) NOT NULL UNIQUE,
    description TEXT,
    owner_id UUID NOT NULL,
    is_active BOOLEAN DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    
    -- Foreign key constraints
    CONSTRAINT fk_teams_owner FOREIGN KEY (owner_id) REFERENCES users(id) ON DELETE CASCADE,
    
    -- Check constraints
    CONSTRAINT chk_teams_name_length CHECK (LENGTH(name) >= 2)
);

-- Team members table (composite primary key)
CREATE TABLE team_members (
    user_id UUID NOT NULL,
    team_id UUID NOT NULL,
    role VARCHAR(20) NOT NULL DEFAULT 'member',
    is_active BOOLEAN DEFAULT TRUE,
    joined_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    left_at TIMESTAMP NULL,
    invited_by UUID,
    
    -- Composite primary key (prevents duplicate user-team combinations)
    PRIMARY KEY (user_id, team_id),
    
    -- Foreign key constraints
    CONSTRAINT fk_team_members_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT fk_team_members_team FOREIGN KEY (team_id) REFERENCES teams(id) ON DELETE CASCADE,
    CONSTRAINT fk_team_members_inviter FOREIGN KEY (invited_by) REFERENCES users(id) ON DELETE SET NULL,
    
    -- Check constraints
    -- To modify roles: ALTER TABLE team_members DROP CONSTRAINT chk_team_members_role, ADD CONSTRAINT chk_team_members_role CHECK (role IN ('new', 'values'));
    CONSTRAINT chk_team_members_role CHECK (role IN ('owner', 'admin', 'member', 'viewer')),
    CONSTRAINT chk_team_members_dates CHECK (left_at IS NULL OR left_at >= joined_at)
);



-- Performance indexes
CREATE INDEX idx_teams_owner_id ON teams(owner_id);
CREATE INDEX idx_teams_name ON teams(name);
CREATE INDEX idx_teams_active ON teams(is_active);

-- team_members already has composite PK (user_id, team_id)
-- Additional indexes for common queries
CREATE INDEX idx_team_members_team_id ON team_members(team_id);
CREATE INDEX idx_team_members_role ON team_members(role);
CREATE INDEX idx_team_members_active ON team_members(is_active);
CREATE INDEX idx_team_members_joined ON team_members(joined_at);



-- Triggers for updated_at (PostgreSQL)
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_teams_updated_at
    BEFORE UPDATE ON teams
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

-- Sample data (optional)
-- INSERT INTO teams (name, description, owner_id) VALUES 
-- ('Development Team', 'Main development team', 'uuid-here'),
-- ('QA Team', 'Quality assurance team', 'uuid-here');

-- Example queries:
-- Get user's teams: SELECT t.* FROM teams t JOIN team_members tm ON t.id = tm.team_id WHERE tm.user_id = 'user-uuid';
-- Get team members: SELECT u.* FROM users u JOIN team_members tm ON u.id = tm.user_id WHERE tm.team_id = 'team-uuid';