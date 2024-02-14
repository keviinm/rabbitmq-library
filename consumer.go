package rabbitmq

import (
    "github.com/streadway/amqp"
    "log"
)

type RabbitMQConsumer struct {
    channel  *RabbitMQChannel
    consumer <-chan amqp.Delivery
}

func NewRabbitMQConsumer(channel *RabbitMQChannel, queueName string) (*RabbitMQConsumer, error) {
    consumer, err := channel.channel.Consume(
        queueName,
        "",
        true,
        false,
        false,
        false,
        nil,
    )
    if err != nil {
        return nil, err
    }
    return &RabbitMQConsumer{channel: channel, consumer: consumer}, nil
}

func (rmq *RabbitMQConsumer) ConsumeMessages() {
    for msg := range rmq.consumer {
        log.Printf("Received message: %s", msg.Body)
    }
}

func (rmq *RabbitMQConsumer) Close() {
    if rmq.channel != nil {
        rmq.channel.Close()
    }
}

