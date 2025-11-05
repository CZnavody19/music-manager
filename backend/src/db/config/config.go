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

func (cs *ConfigStore) GetGeneralConfig(ctx context.Context) (*domain.GeneralConfig, error) {
	stmt := table.General.SELECT(table.General.AllColumns).WHERE(table.General.Active.IS_TRUE())

	var dest model.General
	err := stmt.QueryContext(ctx, cs.DB, &dest)
	if err != nil {
		return nil, err
	}

	return cs.Mapper.MapGeneralConfig(&dest), nil
}

func (cs *ConfigStore) SaveGeneralConfig(ctx context.Context, config *domain.GeneralConfig) error {
	stmt := table.General.INSERT(table.General.AllColumns).MODEL(struct {
		Active bool
		domain.GeneralConfig
	}{
		Active:        true,
		GeneralConfig: *config,
	})

	stmt = db.DoUpsert(stmt, table.General.Active, table.General.MutableColumns, table.General.EXCLUDED.MutableColumns)

	_, err := stmt.ExecContext(ctx, cs.DB)

	return err
}
