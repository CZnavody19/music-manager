-- +goose Up
-- +goose StatementBegin
CREATE SCHEMA IF NOT EXISTS config;

CREATE TABLE config.general (
    enabled BOOLEAN PRIMARY KEY CHECK (enabled = TRUE),
    download_path TEXT NOT NULL,
    temp_path TEXT NOT NULL
);

CREATE TABLE config.youtube (
    enabled BOOLEAN PRIMARY KEY CHECK (enabled = TRUE),
    oauth bytea NOT NULL,
    token bytea NOT NULL
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS config.general;
DROP TABLE IF EXISTS config.youtube;
DROP SCHEMA IF EXISTS config;
-- +goose StatementEnd