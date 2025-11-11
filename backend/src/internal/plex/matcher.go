package plex

import (
	"context"

	"github.com/CZnavody19/music-manager/src/db/plex"
	"github.com/CZnavody19/music-manager/src/domain"
	"github.com/CZnavody19/music-manager/src/mq"
	"github.com/google/uuid"
)

type MatchRequest struct {
	track     *domain.PlexTrack
	plexStore *plex.PlexStore
	Mq        *mq.MessageQueue
}

func (r MatchRequest) GetTrackID() uuid.UUID {
	return *r.track.Mbid
}

func (r MatchRequest) LinkTrack(ctx context.Context, trackID uuid.UUID) error {
	return r.plexStore.LinkTrack(ctx, r.track.ID, trackID)
}

func (r MatchRequest) Done(ctx context.Context, track *domain.Track) error {
	return nil
}
