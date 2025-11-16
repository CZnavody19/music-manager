package config

import (
	"context"
	"time"

	"github.com/CZnavody19/music-manager/src/db"
	"github.com/CZnavody19/music-manager/src/db/gen/musicdb/config/model"
	"github.com/CZnavody19/music-manager/src/db/gen/musicdb/config/table"
	"github.com/CZnavody19/music-manager/src/domain"
	"github.com/go-jet/jet/v2/qrm"
)

func (cs *ConfigStore) GetTidalConfig(ctx context.Context) (*domain.TidalConfig, error) {
	stmt := table.Tidal.SELECT(table.Tidal.AllColumns).WHERE(table.Tidal.Active.IS_TRUE())

	var dest model.Tidal
	err := stmt.QueryContext(ctx, cs.DB, &dest)
	if err == qrm.ErrNoRows {
		return &domain.TidalConfig{
			Enabled:       false,
			AuthExpiresAt: time.Unix(0, 0),
		}, nil
	}
	if err != nil {
		return nil, err
	}

	return cs.Mapper.MapTidalConfig(&dest), nil
}

func (cs *ConfigStore) SaveTidalConfig(ctx context.Context, config *domain.TidalConfig) error {
	stmt := table.Tidal.INSERT(table.Tidal.AllColumns.Except(table.Tidal.Enabled)).MODEL(struct {
		Active bool
		domain.TidalConfig
	}{
		Active:      true,
		TidalConfig: *config,
	})

	stmt = db.DoUpsert(stmt, table.Tidal.Active, table.Tidal.MutableColumns.Except(table.Tidal.Enabled), table.Tidal.EXCLUDED.MutableColumns.Except(table.Tidal.Enabled))

	_, err := stmt.ExecContext(ctx, cs.DB)

	return err
}

func (cs *ConfigStore) SetTidalEnabled(ctx context.Context, enabled bool) error {
	stmt := table.Tidal.UPDATE(table.Tidal.Enabled).SET(enabled).WHERE(table.Tidal.Active.IS_TRUE())

	_, err := stmt.ExecContext(ctx, cs.DB)

	return err
}
