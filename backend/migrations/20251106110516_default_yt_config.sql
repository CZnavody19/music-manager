-- +goose Up
-- +goose StatementBegin
ALTER TABLE config.youtube
ALTER COLUMN oauth SET DEFAULT ''::bytea,
ALTER COLUMN token SET DEFAULT ''::bytea;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE config.youtube
ALTER COLUMN oauth DROP DEFAULT,
ALTER COLUMN token DROP DEFAULT;
-- +goose StatementEnd
