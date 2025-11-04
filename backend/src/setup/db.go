package setup

import (
	"database/sql"

	"github.com/CZnavody19/music-manager/src/config"
	"github.com/XSAM/otelsql"
	_ "github.com/lib/pq"
	"github.com/wexder/goose/v3"
	"go.uber.org/zap"
)

func SetupDb(config *config.DBConfig) (*sql.DB, error) {
	dbConn, err := otelsql.Open(config.DriverName, config.ConnectionURI)
	if err != nil {
		zap.L().Error("Error opening DB connection", zap.Error(err))
		return nil, err
	}

	dbConn.SetMaxOpenConns(config.MaxOpenConnections)
	dbConn.SetMaxIdleConns(config.MaxIdleConnections)
	dbConn.SetConnMaxLifetime(config.MaxConnectionLifetime)

	err = dbConn.Ping()
	if err != nil {
		zap.L().Error("Error pinging DB", zap.Error(err))
		return nil, err
	}

	if config.RunMigration {
		zap.L().Info("Running DB migrations")
		err := goose.Up(dbConn, "./migrations")
		if err != nil {
			zap.L().Error("Error running DB migrations", zap.Error(err))
			return nil, err
		}
		zap.L().Info("DB migrations completed successfully")
	}

	return dbConn, nil
}
