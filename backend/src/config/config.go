package config

import (
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Server        ServerConfig
	LoggingConfig LoggingConfig
	DBConfig      DBConfig
	MQConfig      MQConfig
}

func LoadConfig() *Config {
	return &Config{
		Server:        loadServerConfig(),
		LoggingConfig: loadLoggingConfig(),
		DBConfig:      loadDbConfig(),
		MQConfig:      loadMqConfig(),
	}
}

func configViper(configName string) *viper.Viper {
	v := viper.New()
	v.SetEnvPrefix(strings.ToUpper(configName))
	v.AutomaticEnv()
	v.SetConfigName(configName)
	v.SetConfigType("yaml")
	v.AddConfigPath("./configuration/")
	return v
}
