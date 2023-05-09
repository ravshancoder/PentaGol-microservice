CREATE TABLE IF NOT EXISTS admins (
    id            SERIAL,
    name          VARCHAR(50) NOT NULL,
    email         VARCHAR(100) NOT NULL UNIQUE,
    password      TEXT NOT NULL,
    refresh_token TEXT,
    created_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at    TIME
);