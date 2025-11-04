package config

import "time"

type DBConfig struct {
	ConnectionURI         string
	DriverName            string
	RunMigration          bool
	MaxOpenConnections    int
	MaxIdleConnections    int
	MaxConnectionLifetime time.Duration
}

func loadDbConfig() DBConfig {
	dbConfig := &DBConfig{}
	v := configViper("db")
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}
	v.Unmarshal(dbConfig)
	return *dbConfig
}
