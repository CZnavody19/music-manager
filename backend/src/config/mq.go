package config

type MQConfig struct {
	ConnectionURI string
}

func loadMqConfig() MQConfig {
	mqConfig := &MQConfig{}
	v := configViper("mq")
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}
	v.Unmarshal(mqConfig)
	return *mqConfig
}
