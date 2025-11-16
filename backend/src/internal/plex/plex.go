package plex

import (
	"context"
	"strconv"
	"time"

	"github.com/CZnavody19/music-manager/plexapi"
	"github.com/CZnavody19/music-manager/plexapi/options"
	"github.com/CZnavody19/music-manager/src/db/config"
	"github.com/CZnavody19/music-manager/src/db/plex"
	"github.com/CZnavody19/music-manager/src/domain"
	"github.com/CZnavody19/music-manager/src/graph/model"
	"github.com/CZnavody19/music-manager/src/internal/musicbrainz"
	"github.com/CZnavody19/music-manager/src/internal/websockets"
	"github.com/go-jet/jet/v2/qrm"
	"go.uber.org/zap"
)

type Plex struct {
	enabled     bool
	configStore *config.ConfigStore
	config      *domain.PlexConfig
	client      *plexapi.Client
	websockets  *websockets.Websockets
	plexStore   *plex.PlexStore
	musicBrainz *musicbrainz.MusicBrainz
}

func getPlexAPI(cfg *domain.PlexConfig) *plexapi.Client {
	if cfg == nil {
		return nil
	}

	client := plexapi.NewClient(options.ClientOptions{
		Protocol: cfg.Protocol,
		Host:     cfg.Host,
		Port:     int(cfg.Port),
		Token:    cfg.Token,
	})

	return client
}

func NewPlex(cs *config.ConfigStore, ps *plex.PlexStore, mb *musicbrainz.MusicBrainz, ws *websockets.Websockets) (*Plex, error) {
	ctx := context.Background()

	config, err := cs.GetPlexConfig(ctx)
	if err != nil && err != qrm.ErrNoRows {
		return nil, err
	}

	enabled := false
	if config != nil && config.Enabled {
		enabled = true
	}

	return &Plex{
		enabled:     enabled,
		configStore: cs,
		config:      config,
		client:      getPlexAPI(config),
		websockets:  ws,
		plexStore:   ps,
		musicBrainz: mb,
	}, nil
}

func (p *Plex) IsEnabled() bool {
	return p.enabled
}

func (p *Plex) Enable(ctx context.Context) error {
	config, err := p.configStore.GetPlexConfig(ctx)
	if err != nil {
		return err
	}

	err = p.configStore.SetPlexEnabled(ctx, true)
	if err != nil {
		return err
	}

	p.client = getPlexAPI(config)
	p.config = config
	p.enabled = true
	return nil
}

func (p *Plex) Disable(ctx context.Context) error {
	err := p.configStore.SetPlexEnabled(ctx, false)
	if err != nil {
		return err
	}

	p.client = nil
	p.config = nil
	p.enabled = false
	return nil
}

func (p *Plex) GetTracks(ctx context.Context) ([]*domain.PlexTrack, error) {
	if !p.enabled {
		return nil, nil
	}

	tracks, err := p.plexStore.GetTracks(ctx, false)
	if err != nil {
		return nil, err
	}

	return tracks, nil
}

func (p *Plex) RefreshTracks(ctx context.Context) error {
	if !p.enabled {
		return nil
	}

	start := time.Now()
	p.websockets.SendTask(&model.Task{
		Title:     "Refreshing Plex library",
		StartedAt: start,
		Ended:     false,
	})

	zap.S().Info("Refreshing Plex tracks")

	secRes, err := p.client.Content.GetSectionLeaves(ctx, int(p.config.LibraryID))
	if err != nil {
		return err
	}

	var trackIds []string
	trackMap := make(map[string]domain.PlexTrack)
	for _, track := range secRes.MediaContainer.Metadata {
		trackIds = append(trackIds, track.RatingKey)

		id, err := strconv.ParseInt(track.RatingKey, 10, 64)
		if err != nil {
			return err
		}

		trackMap[track.RatingKey] = domain.PlexTrack{
			ID:       id,
			Title:    track.Title,
			Artist:   track.GrandparentTitle,
			Duration: int64(track.Duration),
		}
	}

	metRes, err := p.client.Content.GetMetadataItem(ctx, trackIds)
	if err != nil {
		return err
	}

	var tracks []domain.PlexTrack
	for _, item := range metRes.MediaContainer.Metadata {
		track := trackMap[item.RatingKey]

		if len(item.GUID) == 0 {
			zap.S().Warnf("No GUID found for track ID %d, probably not matched correctly", track.ID)
			track.Mbid = nil
		} else {
			id, err := mapMbid(item.GUID[0].ID)
			if err != nil {
				zap.S().Warnf("Failed to map MBID for track ID %d: %v", track.ID, err)
				track.Mbid = nil
			} else {
				track.Mbid = id
			}
		}

		tracks = append(tracks, track)
	}

	err = p.plexStore.StoreTracks(ctx, tracks)
	if err != nil {
		return err
	}

	unmatched, err := p.plexStore.GetTracks(ctx, true)
	if err != nil {
		return err
	}

	for _, track := range unmatched {
		if track.Mbid == nil {
			continue
		}

		p.musicBrainz.MatchQueue <- MatchRequest{
			track:     track,
			plexStore: p.plexStore,
		}
	}

	p.websockets.SendTask(&model.Task{
		Title:     "Refreshing Plex library",
		StartedAt: start,
		Ended:     true,
	})

	zap.S().Infof("Refreshed %d Plex tracks", len(tracks))

	return nil
}

func (p *Plex) RefreshLibrary(ctx context.Context) error {
	if !p.enabled {
		return nil
	}

	err := p.client.Library.RefreshSection(ctx, int(p.config.LibraryID))
	if err != nil {
		return err
	}

	return nil
}

func (p *Plex) DeleteTrack(ctx context.Context, id int64) error {
	if !p.enabled {
		return nil
	}

	err := p.plexStore.DeleteTrack(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
