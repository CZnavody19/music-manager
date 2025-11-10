package orchestration

import (
	"context"

	"github.com/CZnavody19/music-manager/src/internal/plex"
	"github.com/CZnavody19/music-manager/src/internal/youtube"
)

type Orchestrator struct {
	plex    *plex.Plex
	youtube *youtube.YouTube
}

func NewOrchestrator(pl *plex.Plex, yt *youtube.YouTube) (*Orchestrator, error) {
	return &Orchestrator{
		plex:    pl,
		youtube: yt,
	}, nil
}

// Gets called by a CRON job
func (o *Orchestrator) CRONjob(ctx context.Context) error {
	err := o.youtube.RefreshPlaylist(ctx)
	if err != nil {
		return err
	}

	err = o.plex.RefreshTracks(ctx)
	if err != nil {
		return err
	}

	return nil
}
