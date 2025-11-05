-- +goose Up
-- +goose StatementBegin
ALTER TABLE config.youtube
RENAME COLUMN enabled TO active;

ALTER TABLE config.discord
RENAME COLUMN enabled TO active;

ALTER TABLE config.plex
RENAME COLUMN enabled TO active;

ALTER TABLE config.youtube
ADD COLUMN enabled BOOLEAN NOT NULL DEFAULT TRUE;

ALTER TABLE config.discord
ADD COLUMN enabled BOOLEAN NOT NULL DEFAULT TRUE;

ALTER TABLE config.plex
ADD COLUMN enabled BOOLEAN NOT NULL DEFAULT TRUE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE config.youtube
DROP COLUMN enabled;

ALTER TABLE config.discord
DROP COLUMN enabled;

ALTER TABLE config.plex
DROP COLUMN enabled;

ALTER TABLE config.youtube
RENAME COLUMN active TO enabled;

ALTER TABLE config.discord
RENAME COLUMN active TO enabled;

ALTER TABLE config.plex
RENAME COLUMN active TO enabled;
-- +goose StatementEnd
