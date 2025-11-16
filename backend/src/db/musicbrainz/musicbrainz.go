package musicbrainz

import (
	"context"
	"database/sql"

	"github.com/CZnavody19/music-manager/src/db"
	"github.com/CZnavody19/music-manager/src/db/gen/musicdb/public/table"
	"github.com/CZnavody19/music-manager/src/domain"
	"github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"
)

type MusicbrainzStore struct {
	DB     *sql.DB
	Mapper *db.Mapper
}

func NewMusicbrainzStore(db *sql.DB, m *db.Mapper) *MusicbrainzStore {
	return &MusicbrainzStore{
		DB:     db,
		Mapper: m,
	}
}

func (ms *MusicbrainzStore) StoreTrack(ctx context.Context, track domain.Track) error {
	tx, err := ms.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	trackStmt := table.Tracks.INSERT(table.Tracks.AllColumns).MODEL(track).ON_CONFLICT().DO_NOTHING()

	_, err = trackStmt.ExecContext(ctx, tx)
	if err != nil {
		return err
	}

	if len(track.ISRCs) > 0 {
		isrcStmt := table.TrackIsrcs.INSERT(table.TrackIsrcs.AllColumns).MODELS(mapISRCs(track)).ON_CONFLICT().DO_NOTHING()
		_, err = isrcStmt.ExecContext(ctx, tx)
		if err != nil {
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (ms *MusicbrainzStore) GetTracks(ctx context.Context, notDownloaded bool) ([]*domain.Track, error) {
	ytStmt := postgres.EXISTS(table.Youtube.SELECT(table.Youtube.VideoID).WHERE(table.Youtube.TrackID.EQ(table.Tracks.ID)))
	plexStmt := postgres.EXISTS(table.Plex.SELECT(table.Plex.ID).WHERE(table.Plex.TrackID.EQ(table.Tracks.ID)))

	stmt := postgres.SELECT(table.Tracks.AllColumns, table.TrackIsrcs.AllColumns, ytStmt.AS("trackwithisrcs.youtube_exists"), plexStmt.AS("trackwithisrcs.plex_exists")).FROM(table.Tracks.
		LEFT_JOIN(table.TrackIsrcs, table.TrackIsrcs.TrackID.EQ(table.Tracks.ID)))

	if notDownloaded {
		stmt = stmt.WHERE(plexStmt.IS_FALSE())
	}

	var dest []db.TrackWithISRCs
	err := stmt.QueryContext(ctx, ms.DB, &dest)
	if err != nil {
		return nil, err
	}

	return ms.Mapper.MapTracksWithISRCs(dest), nil
}

func (ms *MusicbrainzStore) DeleteTrack(ctx context.Context, id uuid.UUID) error {
	stmt := table.Tracks.DELETE().WHERE(table.Tracks.ID.EQ(postgres.UUID(id)))

	_, err := stmt.ExecContext(ctx, ms.DB)
	if err != nil {
		return err
	}

	return nil
}
