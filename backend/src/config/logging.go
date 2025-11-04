package config

type LoggingConfig struct {
	EnableDebugLogger bool
	EnableFileLogger  bool
	FileLogLevel      string
	FileLogOutput     string
	LoggerConfig      LoggerConfig
}

type LoggerConfig struct {
	Development       bool
	DisableCaller     bool
	DisableStacktrace bool
	Sampling          SamplingConfig
	SamplingEnable    bool
	InitialFields     map[string]interface{}
}

type SamplingConfig struct {
	Initial    int
	Thereafter int
}

func loadLoggingConfig() LoggingConfig {
	loggingConfig := &LoggingConfig{}
	v := configViper("logging")
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}
	v.Unmarshal(loggingConfig)
	return *loggingConfig
}
