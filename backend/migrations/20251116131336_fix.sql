-- +goose Up
-- +goose StatementBegin
ALTER TABLE config.tidal
ALTER COLUMN enabled SET DEFAULT true;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE config.tidal
ALTER COLUMN enabled DROP DEFAULT;
-- +goose StatementEnd
