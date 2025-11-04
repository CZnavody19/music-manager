-- +goose Up
-- +goose StatementBegin
CREATE TABLE config.discord (
    enabled BOOLEAN PRIMARY KEY CHECK (enabled = TRUE),
    webhook_url TEXT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE config.discord;
-- +goose StatementEnd
