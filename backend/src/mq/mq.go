package mq

import (
	"context"
	"encoding/json"

	"github.com/CZnavody19/music-manager/src/domain"
	amqp "github.com/rabbitmq/amqp091-go"
)

type MessageQueue struct {
	conn *amqp.Connection
}

func NewMessageQueue(conn *amqp.Connection) *MessageQueue {
	return &MessageQueue{
		conn: conn,
	}
}

func (mq *MessageQueue) Download(ctx context.Context, track *domain.Track) error {
	ch, err := mq.conn.Channel()
	if err != nil {
		return err
	}

	body, err := json.Marshal(track)
	if err != nil {
		return err
	}

	err = ch.PublishWithContext(ctx, "downloads", "tidal", false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        body,
	})

	if err != nil {
		return err
	}

	return nil
}
