-- +goose Up
-- +goose StatementBegin
CREATE TABLE config.tidal (
    active BOOLEAN PRIMARY KEY CHECK (active = TRUE),
    auth_token_type TEXT NOT NULL,
    auth_access_token TEXT NOT NULL,
    auth_refresh_token TEXT NOT NULL,
    auth_expires_at TIMESTAMPTZ NOT NULL,
    auth_client_id TEXT NOT NULL,
    auth_client_secret TEXT NOT NULL,
    download_timeout INTEGER NOT NULL,
    download_retries INTEGER NOT NULL,
    download_threads INTEGER NOT NULL,
    audio_quality TEXT NOT NULL,
    enabled BOOLEAN NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE config.tidal;
-- +goose StatementEnd
