package setup

import (
	"github.com/CZnavody19/music-manager/src/config"
	amqp "github.com/rabbitmq/amqp091-go"
)

func SetupMq(config *config.MQConfig) (*amqp.Connection, error) {
	conn, err := amqp.Dial(config.ConnectionURI)
	if err != nil {
		return nil, err
	}

	chann, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	err = chann.ExchangeDeclare("downloads", "direct", true, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	err = chann.ExchangeDeclare("downloads_complete", "direct", true, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	err = chann.ExchangeDeclare("reload", "direct", true, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
