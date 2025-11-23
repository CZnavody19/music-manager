-- +goose Up
-- +goose StatementBegin
ALTER TABLE config.tidal
ADD COLUMN file_permissions INTEGER NOT NULL CHECK (file_permissions BETWEEN 0 AND 511) DEFAULT 436, -- 0o664 in octal
ADD COLUMN directory_permissions INTEGER NOT NULL CHECK (directory_permissions BETWEEN 0 AND 511) DEFAULT 509, -- 0o775 in octal
ADD COLUMN ownr INTEGER NOT NULL DEFAULT 0,
ADD COLUMN grp INTEGER NOT NULL DEFAULT 0;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE config.tidal
DROP COLUMN file_permissions,
DROP COLUMN directory_permissions,
DROP COLUMN ownr,
DROP COLUMN grp;
-- +goose StatementEnd
