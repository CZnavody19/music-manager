package plex

import (
	"context"
	"database/sql"

	"github.com/CZnavody19/music-manager/src/db"
	"github.com/CZnavody19/music-manager/src/db/gen/musicdb/public/model"
	"github.com/CZnavody19/music-manager/src/db/gen/musicdb/public/table"
	"github.com/CZnavody19/music-manager/src/domain"
	"github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"
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

func (ps *PlexStore) GetTracks(ctx context.Context, unmatched bool) ([]*domain.PlexTrack, error) {
	stmt := table.Plex.SELECT(table.Plex.AllColumns)

	if unmatched {
		stmt = stmt.WHERE(table.Plex.TrackID.IS_NULL().AND(table.Plex.Mbid.IS_NOT_NULL())).LIMIT(100)
	}

	var tracks []*model.Plex
	err := stmt.QueryContext(ctx, ps.DB, &tracks)
	if err != nil {
		return nil, err
	}

	return ps.Mapper.MapPlexTracks(tracks), nil
}

func (ps *PlexStore) LinkTrack(ctx context.Context, plexID int64, trackID uuid.UUID) error {
	stmt := table.Plex.UPDATE().SET(
		table.Plex.TrackID.SET(postgres.UUID(trackID)),
	).WHERE(
		table.Plex.ID.EQ(postgres.Int64(plexID)),
	)

	_, err := stmt.ExecContext(ctx, ps.DB)
	return err
}
