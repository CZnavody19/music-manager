package config

import (
	"context"
	"database/sql"

	"github.com/CZnavody19/music-manager/src/db"
	"github.com/CZnavody19/music-manager/src/db/gen/musicdb/config/model"
	"github.com/CZnavody19/music-manager/src/db/gen/musicdb/config/table"
	"github.com/CZnavody19/music-manager/src/domain"
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
	stmt := table.Youtube.SELECT(table.Youtube.AllColumns).WHERE(table.Youtube.Active.IS_TRUE())

	var dest model.Youtube
	err := stmt.QueryContext(ctx, cs.DB, &dest)
	if err != nil {
		return nil, err
	}

	return cs.Mapper.MapYouTubeConfig(&dest), nil
}

func (cs *ConfigStore) SaveYoutubeFiles(ctx context.Context, oauthData, tokenData []byte) error {
	stmt := table.Youtube.INSERT(table.Youtube.OAuth, table.Youtube.Token, table.Youtube.Active).VALUES(oauthData, tokenData, true)

	stmt = db.DoUpsert(stmt, table.Youtube.Active, table.Youtube.MutableColumns, table.Youtube.EXCLUDED.MutableColumns)

	_, err := stmt.ExecContext(ctx, cs.DB)

	return err
}

func (cs *ConfigStore) SetYoutubeEnabled(ctx context.Context, enabled bool) error {
	stmt := table.Youtube.UPDATE(table.Youtube.Enabled).SET(enabled).WHERE(table.Youtube.Active.IS_TRUE())

	_, err := stmt.ExecContext(ctx, cs.DB)

	return err
}

func (cs *ConfigStore) GetDiscordConfig(ctx context.Context) (*domain.DiscordConfig, error) {
	stmt := table.Discord.SELECT(table.Discord.AllColumns).WHERE(table.Discord.Active.IS_TRUE())

	var dest model.Discord
	err := stmt.QueryContext(ctx, cs.DB, &dest)
	if err != nil {
		return nil, err
	}

	return cs.Mapper.MapDiscordConfig(&dest), nil
}

func (cs *ConfigStore) SaveDiscordConfig(ctx context.Context, config *domain.DiscordConfig) error {
	stmt := table.Discord.INSERT(table.Discord.WebhookURL, table.Discord.Active).VALUES(config.WebhookURL, true)

	stmt = db.DoUpsert(stmt, table.Discord.Active, table.Discord.MutableColumns, table.Discord.EXCLUDED.MutableColumns)

	_, err := stmt.ExecContext(ctx, cs.DB)

	return err
}

func (cs *ConfigStore) SetDiscordEnabled(ctx context.Context, enabled bool) error {
	stmt := table.Discord.UPDATE(table.Discord.Enabled).SET(enabled).WHERE(table.Discord.Active.IS_TRUE())

	_, err := stmt.ExecContext(ctx, cs.DB)

	return err
}

func (cs *ConfigStore) GetPlexConfig(ctx context.Context) (*domain.PlexConfig, error) {
	stmt := table.Plex.SELECT(table.Plex.AllColumns).WHERE(table.Plex.Active.IS_TRUE())

	var dest model.Plex
	err := stmt.QueryContext(ctx, cs.DB, &dest)
	if err != nil {
		return nil, err
	}

	return cs.Mapper.MapPlexConfig(&dest), nil
}

func (cs *ConfigStore) SavePlexConfig(ctx context.Context, config *domain.PlexConfig) error {
	stmt := table.Plex.INSERT(table.Plex.AllColumns.Except(table.Plex.Enabled)).MODEL(struct {
		Active bool
		domain.PlexConfig
	}{
		Active:     true,
		PlexConfig: *config,
	})

	stmt = db.DoUpsert(stmt, table.Plex.Active, table.Plex.MutableColumns, table.Plex.EXCLUDED.MutableColumns)

	_, err := stmt.ExecContext(ctx, cs.DB)

	return err
}

func (cs *ConfigStore) SetPlexEnabled(ctx context.Context, enabled bool) error {
	stmt := table.Plex.UPDATE(table.Plex.Enabled).SET(enabled).WHERE(table.Plex.Active.IS_TRUE())

	_, err := stmt.ExecContext(ctx, cs.DB)

	return err
}
