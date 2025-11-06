-- +goose Up
-- +goose StatementBegin
ALTER TABLE config.youtube
ADD COLUMN playlist_id TEXT NOT NULL DEFAULT '';

CREATE TABLE public.youtube (
    video_id TEXT PRIMARY KEY,
    title TEXT NOT NULL,
    channel_title TEXT NOT NULL,
    thumbnail_url TEXT,
    duration INTEGER,
    position INTEGER NOT NULL,
    next_page_token TEXT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE public.youtube;

ALTER TABLE config.youtube
DROP COLUMN playlist_id;
-- +goose StatementEnd
