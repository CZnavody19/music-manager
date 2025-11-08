package musicbrainz

import (
	"context"
	"database/sql"

	"github.com/CZnavody19/music-manager/src/db"
	"github.com/CZnavody19/music-manager/src/db/gen/musicdb/public/table"
	"github.com/CZnavody19/music-manager/src/domain"
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
