package youtube

import (
	"context"
	"encoding/json"

	"github.com/CZnavody19/music-manager/src/db/config"
	"github.com/CZnavody19/music-manager/src/domain"
	"github.com/go-jet/jet/v2/qrm"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	youtubeApi "google.golang.org/api/youtube/v3"
)

type YouTube struct {
	enabled     bool
	configStore *config.ConfigStore
	yt          *youtubeApi.Service
}

func getYtService(ctx context.Context, cfg *domain.YouTubeConfig) (*youtubeApi.Service, error) {
	if cfg == nil {
		return nil, nil
	}

	config, err := google.ConfigFromJSON(cfg.OAuth, "https://www.googleapis.com/auth/youtube.readonly")
	if err != nil {
		return nil, err
	}

	t := &oauth2.Token{}
	err = json.Unmarshal(cfg.Token, t)
	if err != nil {
		return nil, err
	}

	client := config.Client(ctx, t)

	yt, err := youtubeApi.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return nil, err
	}

	return yt, nil
}

func NewYouTube(cs *config.ConfigStore) (*YouTube, error) {
	ctx := context.Background()

	config, err := cs.GetYoutubeConfig(ctx)
	if err != nil && err != qrm.ErrNoRows {
		return nil, err
	}

	service, err := getYtService(ctx, config)
	if err != nil {
		return nil, err
	}

	enabled := false
	if service != nil && config.Enabled {
		enabled = true
	}

	return &YouTube{
		enabled:     enabled,
		configStore: cs,
		yt:          service,
	}, nil
}

func (yt *YouTube) IsEnabled() bool {
	return yt.enabled
}

func (yt *YouTube) Enable(ctx context.Context) error {
	config, err := yt.configStore.GetYoutubeConfig(ctx)
	if err != nil {
		return err
	}

	service, err := getYtService(ctx, config)
	if err != nil {
		return err
	}

	err = yt.configStore.SetYoutubeEnabled(ctx, true)
	if err != nil {
		return err
	}

	yt.yt = service
	yt.enabled = true

	return nil
}

func (yt *YouTube) Disable(ctx context.Context) error {
	err := yt.configStore.SetYoutubeEnabled(ctx, false)
	if err != nil {
		return err
	}

	yt.yt = nil
	yt.enabled = false

	return nil
}
