package mq

import amqp "github.com/rabbitmq/amqp091-go"

type MessageQueue struct {
	conn *amqp.Connection
}

func NewMessageQueue(conn *amqp.Connection) *MessageQueue {
	return &MessageQueue{
		conn: conn,
	}
}
