package config

import (
	"context"

	"github.com/CZnavody19/music-manager/src/db"
	"github.com/CZnavody19/music-manager/src/db/gen/musicdb/config/model"
	"github.com/CZnavody19/music-manager/src/db/gen/musicdb/config/table"
	"github.com/CZnavody19/music-manager/src/domain"
)

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
