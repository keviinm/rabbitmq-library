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

func (rmq *RabbitMQChannel) Publish(exchange, routingKey string, message []byte) error {
    return rmq.channel.Publish(
        exchange,
        routingKey,
        false,
        false,
        amqp.Publishing{
            ContentType: "text/plain",
            Body:        message,
        })
}

func (rmq *RabbitMQChannel) Close() {
    if rmq.channel != nil {
        rmq.channel.Close()
    }
}