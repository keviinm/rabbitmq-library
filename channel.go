package rabbitmq

import (
    "github.com/streadway/amqp"
)

type RabbitMQChannel struct {
    channel *amqp.Channel
}

func NewRabbitMQChannel(conn *amqp.Connection) (*RabbitMQChannel, error) {
    channel, err := conn.Channel()
    if err != nil {
        return nil, err
    }
    return &RabbitMQChannel{channel: channel}, nil
}

func (rmq *RabbitMQChannel) Close() {
    if rmq.channel != nil {
        rmq.channel.Close()
    }
}