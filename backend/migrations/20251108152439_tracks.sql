-- +goose Up
-- +goose StatementBegin
CREATE TABLE tracks (
    id UUID PRIMARY KEY,
    title TEXT NOT NULL,
    artist TEXT NOT NULL,
    length INT NOT NULL
);

CREATE TABLE track_isrcs (
    track_id UUID REFERENCES tracks(id) ON DELETE CASCADE,
    isrc TEXT NOT NULL,
    PRIMARY KEY (track_id, isrc)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE track_isrcs;
DROP TABLE tracks;
-- +goose StatementEnd
