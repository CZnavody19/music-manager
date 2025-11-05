package discord

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/CZnavody19/music-manager/src/db/config"
	"github.com/CZnavody19/music-manager/src/domain"
	"github.com/go-jet/jet/v2/qrm"
)

type Discord struct {
	enabled     bool
	configStore *config.ConfigStore
	webhookURL  *string
}

func NewDiscord(cs *config.ConfigStore) (*Discord, error) {
	ctx := context.Background()

	config, err := cs.GetDiscordConfig(ctx)
	if err != nil && err != qrm.ErrNoRows {
		return nil, err
	}

	var url *string
	enabled := false
	if config != nil && config.Enabled {
		url = &config.WebhookURL
		enabled = true
	}

	return &Discord{
		enabled:     enabled,
		configStore: cs,
		webhookURL:  url,
	}, nil
}

func (d *Discord) IsEnabled() bool {
	return d.enabled
}

func (d *Discord) Enable(ctx context.Context) error {
	config, err := d.configStore.GetDiscordConfig(ctx)
	if err != nil {
		return err
	}

	err = d.configStore.SetDiscordEnabled(ctx, true)
	if err != nil {
		return err
	}

	d.webhookURL = &config.WebhookURL
	d.enabled = true
	return nil
}

func (d *Discord) Disable(ctx context.Context) error {
	err := d.configStore.SetDiscordEnabled(ctx, false)
	if err != nil {
		return err
	}

	d.webhookURL = nil
	d.enabled = false
	return nil
}

func (d *Discord) SendMessage(ctx context.Context, message *domain.DiscordMessage) error {
	if !d.enabled {
		return nil
	}

	data, err := json.Marshal(domain.DiscordMessageSchema{
		Embeds: []domain.DiscordMessage{*message},
	})
	if err != nil {
		return err
	}

	resp, err := http.Post(*d.webhookURL, "application/json", bytes.NewBuffer(data))
	if err != nil {
		return err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("failed to send discord message, status code: %d", resp.StatusCode)
	}

	return nil
}
