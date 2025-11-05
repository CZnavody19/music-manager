package config

import (
	"context"

	"github.com/CZnavody19/music-manager/src/db"
	"github.com/CZnavody19/music-manager/src/db/gen/musicdb/config/model"
	"github.com/CZnavody19/music-manager/src/db/gen/musicdb/config/table"
	"github.com/CZnavody19/music-manager/src/domain"
)

func (cs *ConfigStore) GetAuthConfig(ctx context.Context) (*domain.AuthConfig, error) {
	stmt := table.Auth.SELECT(table.Auth.AllColumns).WHERE(table.Auth.Active.IS_TRUE())

	var dest model.Auth
	err := stmt.QueryContext(ctx, cs.DB, &dest)
	if err != nil {
		return nil, err
	}

	return cs.Mapper.MapAuthConfig(&dest), nil
}

func (cs *ConfigStore) SaveAuthConfig(ctx context.Context, config *domain.AuthConfig) error {
	stmt := table.Auth.INSERT(table.Auth.AllColumns).MODEL(struct {
		Active bool
		domain.AuthConfig
	}{
		Active:     true,
		AuthConfig: *config,
	})

	stmt = db.DoUpsert(stmt, table.Auth.Active, table.Auth.MutableColumns, table.Auth.EXCLUDED.MutableColumns)

	_, err := stmt.ExecContext(ctx, cs.DB)

	return err
}
