-- +goose Up
-- +goose StatementBegin
ALTER TABLE config.general
RENAME COLUMN enabled TO active;

CREATE TABLE config.auth (
    active BOOLEAN PRIMARY KEY CHECK (active = TRUE),
    username TEXT NOT NULL,
    password_hash TEXT NOT NULL
);

INSERT INTO config.auth (active, username, password_hash)
VALUES (TRUE, 'admin', '$2a$12$3JNJzgHS46ZEeaBTgPJLKOmcZDT.TmqS8I9iYwti.qyGpDu.HDvxO') ON CONFLICT DO NOTHING; -- password: meowmeow
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE config.auth;

ALTER TABLE config.general
RENAME COLUMN active TO enabled;
-- +goose StatementEnd
