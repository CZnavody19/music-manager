package musicbrainz

import (
	"context"
	"database/sql"

	"github.com/CZnavody19/music-manager/src/db"
	"github.com/CZnavody19/music-manager/src/db/gen/musicdb/public/table"
	"github.com/CZnavody19/music-manager/src/domain"
	"github.com/go-jet/jet/v2/postgres"
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

func (ms *MusicbrainzStore) GetTracks(ctx context.Context) ([]*domain.Track, error) {
	stmt := postgres.SELECT(table.Tracks.AllColumns, table.TrackIsrcs.AllColumns, table.Youtube.AllColumns, table.Plex.AllColumns).FROM(table.Tracks.
		LEFT_JOIN(table.TrackIsrcs, table.TrackIsrcs.TrackID.EQ(table.Tracks.ID)).
		LEFT_JOIN(table.Youtube, table.Youtube.TrackID.EQ(table.Tracks.ID)).
		LEFT_JOIN(table.Plex, table.Plex.TrackID.EQ(table.Tracks.ID)))

	var dest []db.TrackWithISRCs
	err := stmt.QueryContext(ctx, ms.DB, &dest)
	if err != nil {
		return nil, err
	}

	return ms.Mapper.MapTracksWithISRCs(dest), nil
}
