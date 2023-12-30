-- +goose Up
CREATE EXTENSION IF NOT EXISTS pg_trgm;

CREATE TABLE IF NOT EXISTS players (
    id UUID PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    initials VARCHAR(3),
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(32) NOT NULL,
    created_at TIMESTAMPTZ DEFAULT (CURRENT_TIMESTAMP AT TIME ZONE 'UTC'),
    updated_at TIMESTAMPTZ DEFAULT (CURRENT_TIMESTAMP AT TIME ZONE 'UTC')
);
CREATE INDEX idx_players_first_name_trgm ON players USING GIN (first_name gin_trgm_ops);
CREATE INDEX idx_players_last_name_trgm ON players USING GIN (last_name gin_trgm_ops);
CREATE INDEX idx_players_email_trgm ON players USING GIN (email gin_trgm_ops);

-- +goose Down
DROP TABLE IF EXISTS players;
DROP INDEX IF EXISTS idx_players_first_name_trgm;
DROP INDEX IF EXISTS idx_players_last_name_trgm;
DROP INDEX IF EXISTS idx_players_email_trgm;

