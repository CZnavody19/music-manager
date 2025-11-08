package plex

import (
	"context"
	"fmt"

	"github.com/CZnavody19/music-manager/src/db/config"
	"github.com/CZnavody19/music-manager/src/domain"
	"github.com/LukeHagar/plexgo"
	"github.com/LukeHagar/plexgo/models/operations"
	"github.com/go-jet/jet/v2/qrm"
)

type Plex struct {
	enabled     bool
	configStore *config.ConfigStore
	config      *domain.PlexConfig
	plex        *plexgo.PlexAPI
}

func getPlexAPI(cfg *domain.PlexConfig) *plexgo.PlexAPI {
	if cfg == nil {
		return nil
	}

	plex := plexgo.New(
		plexgo.WithServerIndex(1),
		plexgo.WithProtocol(cfg.Protocol),
		plexgo.WithHost(cfg.Host),
		plexgo.WithPort(fmt.Sprint(cfg.Port)),
		plexgo.WithSecurity(cfg.Token),
	)

	return plex
}

func NewPlex(cs *config.ConfigStore) (*Plex, error) {
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
		plex:        getPlexAPI(config),
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

	p.plex = getPlexAPI(config)
	p.config = config
	p.enabled = true
	return nil
}

func (p *Plex) Disable(ctx context.Context) error {
	err := p.configStore.SetPlexEnabled(ctx, false)
	if err != nil {
		return err
	}

	p.plex = nil
	p.config = nil
	p.enabled = false
	return nil
}

func (p *Plex) RefreshLibrary(ctx context.Context) error {
	if !p.enabled {
		return nil
	}

	_, err := p.plex.Library.RefreshSection(ctx, operations.RefreshSectionRequest{
		SectionID: p.config.LibraryID,
	})

	return err
}
