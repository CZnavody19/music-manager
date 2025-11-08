package youtube

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/CZnavody19/music-manager/src/db"
	"github.com/CZnavody19/music-manager/src/db/gen/musicdb/public/model"
	"github.com/CZnavody19/music-manager/src/db/gen/musicdb/public/table"
	"github.com/CZnavody19/music-manager/src/domain"
	"github.com/go-jet/jet/v2/postgres"
	"github.com/google/uuid"
)

type YouTubeStore struct {
	DB     *sql.DB
	Mapper *db.Mapper
}

func NewYouTubeStore(db *sql.DB, m *db.Mapper) *YouTubeStore {
	return &YouTubeStore{
		DB:     db,
		Mapper: m,
	}
}

func (yts *YouTubeStore) GetLatestPageToken(ctx context.Context) (string, error) {
	stmt := table.Youtube.SELECT(table.Youtube.NextPageToken.AS("StringVal.Value")).ORDER_BY(table.Youtube.Position.DESC()).LIMIT(1)

	var token db.StringVal
	err := stmt.QueryContext(ctx, yts.DB, &token)
	if err != nil {
		return "", err
	}

	return token.Value, nil
}

func (yts *YouTubeStore) StoreVideos(ctx context.Context, videos []*domain.YouTubeVideo) error {
	stmt := table.Youtube.INSERT(table.Youtube.AllColumns).MODELS(videos).ON_CONFLICT(table.Youtube.VideoID).DO_NOTHING()

	res, err := stmt.ExecContext(ctx, yts.DB)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	fmt.Printf("Inserted %d new YouTube videos\n", affected)

	return nil
}

func (yts *YouTubeStore) GetVideos(ctx context.Context, unmatched bool) ([]*domain.YouTubeVideo, error) {
	stmt := table.Youtube.SELECT(table.Youtube.AllColumns).ORDER_BY(table.Youtube.Position.ASC())

	if unmatched {
		stmt = stmt.WHERE(table.Youtube.TrackID.IS_NULL())
	}

	var videos []*model.Youtube
	err := stmt.QueryContext(ctx, yts.DB, &videos)
	if err != nil {
		return nil, err
	}

	return yts.Mapper.MapYoutubeVideos(videos), nil
}

func (yts *YouTubeStore) LinkTrack(ctx context.Context, videoID string, trackID uuid.UUID) error {
	stmt := table.Youtube.UPDATE().SET(
		table.Youtube.TrackID.SET(postgres.UUID(trackID)),
	).WHERE(
		table.Youtube.VideoID.EQ(postgres.String(videoID)),
	)

	_, err := stmt.ExecContext(ctx, yts.DB)
	if err != nil {
		return err
	}

	return nil
}
