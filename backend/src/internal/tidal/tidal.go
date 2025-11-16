package tidal

import (
	"context"
	"time"

	"github.com/CZnavody19/music-manager/src/db/config"
	"github.com/CZnavody19/music-manager/src/domain"
	"github.com/CZnavody19/music-manager/src/graph/model"
	"github.com/CZnavody19/music-manager/src/internal/websockets"
	"github.com/CZnavody19/music-manager/src/mq"
	"github.com/go-jet/jet/v2/qrm"
	"go.uber.org/zap"
)

type Tidal struct {
	enabled     bool
	configStore *config.ConfigStore
	websockets  *websockets.Websockets
	mq          *mq.MessageQueue
}

func NewTidal(cs *config.ConfigStore, ws *websockets.Websockets, mq *mq.MessageQueue) (*Tidal, error) {
	ctx := context.Background()

	config, err := cs.GetTidalConfig(ctx)
	if err != nil && err != qrm.ErrNoRows {
		return nil, err
	}

	enabled := false
	if config != nil && config.Enabled {
		enabled = true
	}

	return &Tidal{
		enabled:     enabled,
		configStore: cs,
		websockets:  ws,
		mq:          mq,
	}, nil
}

func (t *Tidal) IsEnabled() bool {
	return t.enabled
}

func (t *Tidal) Enable(ctx context.Context) error {
	err := t.configStore.SetTidalEnabled(ctx, true)
	if err != nil {
		return err
	}

	err = t.mq.Reload(ctx, "tidal")
	if err != nil {
		return err
	}

	t.enabled = true

	return nil
}

func (t *Tidal) Disable(ctx context.Context) error {
	err := t.configStore.SetTidalEnabled(ctx, false)
	if err != nil {
		return err
	}

	err = t.mq.Reload(ctx, "tidal")
	if err != nil {
		return err
	}

	t.enabled = false

	return nil
}

func (t *Tidal) Download(ctx context.Context, track *domain.Track) error {
	if !t.enabled {
		return nil
	}

	start := time.Now()
	t.websockets.SendTask(&model.Task{
		Title:     "Downloading from Tidal",
		StartedAt: start,
		Ended:     false,
	})

	zap.S().Info("Downloading from Tidal")

	err := t.mq.Download(ctx, track, "tidal")
	if err != nil {
		return err
	}

	t.websockets.SendTask(&model.Task{
		Title:     "Downloading from Tidal",
		StartedAt: start,
		Ended:     true,
	})

	zap.S().Info("Finished downloading from Tidal")

	return nil
}
