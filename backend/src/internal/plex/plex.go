package plex

import (
	"context"
	"fmt"
	"net/http"

	"github.com/CZnavody19/music-manager/src/db/config"
	"github.com/CZnavody19/music-manager/src/domain"
	"github.com/go-jet/jet/v2/qrm"
)

type Plex struct {
	enabled     bool
	configStore *config.ConfigStore
	config      *domain.PlexConfig
}

func NewPlex(cs *config.ConfigStore) (*Plex, error) {
	ctx := context.Background()

	config, err := cs.GetPlexConfig(ctx)
	if err != nil && err != qrm.ErrNoRows {
		return nil, err
	}

	enabled := false
	if config != nil {
		enabled = true
	}

	return &Plex{
		enabled:     enabled,
		configStore: cs,
		config:      config,
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

	p.config = config
	p.enabled = true
	return nil
}

func (p *Plex) Disable(ctx context.Context) error {
	p.config = nil
	p.enabled = false
	return nil
}

func (p *Plex) RefreshLibrary(ctx context.Context) error {
	if !p.enabled {
		return nil
	}

	resp, err := http.Get(fmt.Sprintf("%s://%s:%d/library/sections/%d/refresh?X-Plex-Token=%s", p.config.Protocol, p.config.Host, p.config.Port, p.config.LibraryID, p.config.Token))
	if err != nil {
		return err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("failed to refresh plex library, status code: %d", resp.StatusCode)
	}

	return nil
}
