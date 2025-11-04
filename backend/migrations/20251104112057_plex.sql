-- +goose Up
-- +goose StatementBegin
CREATE TABLE config.plex (
    enabled BOOLEAN PRIMARY KEY CHECK (enabled = TRUE),
    protocol TEXT NOT NULL CHECK (protocol IN ('http', 'https')),
    host TEXT NOT NULL,
    port INTEGER NOT NULL,
    token TEXT NOT NULL,
    library_id INTEGER NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE config.plex;
-- +goose StatementEnd
