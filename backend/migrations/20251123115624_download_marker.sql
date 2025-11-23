-- +goose Up
-- +goose StatementBegin
ALTER TABLE tracks
ADD COLUMN downloaded BOOLEAN;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE tracks
DROP COLUMN downloaded;
-- +goose StatementEnd
