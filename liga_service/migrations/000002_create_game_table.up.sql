CREATE TABLE IF NOT EXISTS games (
    id SERIAL PRIMARY KEY,
    time VARCHAR(10) NOT NULL,
    condtion BOOLEAN NOT NULL,
    first_team_id VARCHAR NOT NULL,
    second_team_id VARCHAR NOT NULL,
    result_first_team INT NOT NULL,
    result_second_team INT NOT NULL,
    first_team_point INT NOT NULL,
    second_team_point INT NOT NULL,
    liga_id INT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIME
);