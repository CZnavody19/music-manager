package config

import (
	"context"
	"database/sql"

	"github.com/CZnavody19/music-manager/src/db"
	"github.com/CZnavody19/music-manager/src/db/gen/musicdb/config/model"
	"github.com/CZnavody19/music-manager/src/db/gen/musicdb/config/table"
	"github.com/CZnavody19/music-manager/src/domain"
	"github.com/go-jet/jet/v2/postgres"
)

type ConfigStore struct {
	DB     *sql.DB
	Mapper *db.Mapper
}

func NewConfigStore(db *sql.DB, m *db.Mapper) *ConfigStore {
	return &ConfigStore{
		DB:     db,
		Mapper: m,
	}
}

func (cs *ConfigStore) GetYoutubeConfig(ctx context.Context) (*domain.YouTubeConfig, error) {
	stmt := table.Youtube.SELECT(table.Youtube.AllColumns).WHERE(table.Youtube.Enabled.IS_TRUE())

	var dest model.Youtube
	err := stmt.QueryContext(ctx, cs.DB, &dest)
	if err != nil {
		return nil, err
	}

	return cs.Mapper.MapYouTubeConfig(&dest), nil
}

func (cs *ConfigStore) SaveYoutubeFiles(ctx context.Context, oauthData, tokenData []byte) error {
	stmt := table.Youtube.INSERT(table.Youtube.OAuth, table.Youtube.Token, table.Youtube.Enabled).VALUES(oauthData, tokenData, true).ON_CONFLICT(table.Youtube.Enabled).DO_UPDATE(
		postgres.SET(
			table.Youtube.OAuth.SET(table.Youtube.EXCLUDED.OAuth),
			table.Youtube.Token.SET(table.Youtube.EXCLUDED.Token),
		),
	)

	_, err := stmt.ExecContext(ctx, cs.DB)

	return err
}

func (cs *ConfigStore) GetDiscordConfig(ctx context.Context) (*domain.DiscordConfig, error) {
	stmt := table.Discord.SELECT(table.Discord.AllColumns).WHERE(table.Discord.Enabled.IS_TRUE())

	var dest model.Discord
	err := stmt.QueryContext(ctx, cs.DB, &dest)
	if err != nil {
		return nil, err
	}

	return cs.Mapper.MapDiscordConfig(&dest), nil
}

func (cs *ConfigStore) SaveDiscordConfig(ctx context.Context, config *domain.DiscordConfig) error {
	stmt := table.Discord.INSERT(table.Discord.WebhookURL, table.Discord.Enabled).VALUES(config.WebhookURL, true).ON_CONFLICT(table.Discord.Enabled).DO_UPDATE(
		postgres.SET(
			table.Discord.WebhookURL.SET(table.Discord.EXCLUDED.WebhookURL),
		),
	)

	_, err := stmt.ExecContext(ctx, cs.DB)

	return err
}
