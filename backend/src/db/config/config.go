package config

import (
	"database/sql"

	"github.com/CZnavody19/music-manager/src/db"
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
