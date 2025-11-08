package plex

import (
	"context"
	"database/sql"

	"github.com/CZnavody19/music-manager/src/db"
	"github.com/CZnavody19/music-manager/src/db/gen/musicdb/public/table"
	"github.com/CZnavody19/music-manager/src/domain"
)

type PlexStore struct {
	DB     *sql.DB
	Mapper *db.Mapper
}

func NewPlexStore(db *sql.DB, mapper *db.Mapper) *PlexStore {
	return &PlexStore{
		DB:     db,
		Mapper: mapper,
	}
}

func (ps *PlexStore) StoreTracks(ctx context.Context, tracks []domain.PlexTrack) error {
	stmt := table.Plex.INSERT(table.Plex.AllColumns).MODELS(tracks)

	_, err := stmt.ExecContext(ctx, ps.DB)
	if err != nil {
		return err
	}

	return nil
}
