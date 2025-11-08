-- +goose Up
-- +goose StatementBegin
CREATE TABLE public.plex (
    id INT PRIMARY KEY,
    title TEXT NOT NULL,
    artist TEXT NOT NULL,
    duration INT NOT NULL,
    mbid UUID,
    track_id UUID REFERENCES public.tracks(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE public.plex;
-- +goose StatementEnd
