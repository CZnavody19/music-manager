package config

import (
	"context"

	"github.com/CZnavody19/music-manager/src/db"
	"github.com/CZnavody19/music-manager/src/db/gen/musicdb/config/model"
	"github.com/CZnavody19/music-manager/src/db/gen/musicdb/config/table"
	"github.com/CZnavody19/music-manager/src/domain"
)

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
