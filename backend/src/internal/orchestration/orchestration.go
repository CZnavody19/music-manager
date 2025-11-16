package orchestration

import (
	"context"

	"github.com/CZnavody19/music-manager/src/internal/musicbrainz"
	"github.com/CZnavody19/music-manager/src/internal/plex"
	"github.com/CZnavody19/music-manager/src/internal/tidal"
	"github.com/CZnavody19/music-manager/src/internal/youtube"
)

type Orchestrator struct {
	musicbrainz *musicbrainz.MusicBrainz
	plex        *plex.Plex
	youtube     *youtube.YouTube
	tidal       *tidal.Tidal
}

func NewOrchestrator(mb *musicbrainz.MusicBrainz, pl *plex.Plex, yt *youtube.YouTube, td *tidal.Tidal) (*Orchestrator, error) {
	return &Orchestrator{
		musicbrainz: mb,
		plex:        pl,
		youtube:     yt,
		tidal:       td,
	}, nil
}

// Gets called by a CRON job
func (o *Orchestrator) Refresh(ctx context.Context) error {
	err := o.plex.RefreshTracks(ctx)
	if err != nil {
		return err
	}

	err = o.youtube.RefreshPlaylist(ctx)
	if err != nil {
		return err
	}

	return nil
}

// Gets called by a CRON job
func (o *Orchestrator) Download(ctx context.Context) error {
	tracks, err := o.musicbrainz.GetTracks(ctx, true)
	if err != nil {
		return err
	}

	for _, track := range tracks {
		err = o.tidal.Download(ctx, track)
		if err != nil {
			return err
		}
	}

	return nil
}
