-- +goose Up
-- +goose StatementBegin
INSERT INTO tracks (id, title, artist, length)
VALUES ('00000000-0000-4000-8000-000000000000', 'Blocked track', 'Blocked track', 0);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE FROM tracks WHERE id = '00000000-0000-4000-8000-000000000000';
-- +goose StatementEnd
