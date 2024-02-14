// Handle connection
package rabbitmq

import (
	"github.com/streadway/amqp"
)

type RabbitMQConnection struct {
	conn *amqp.Connection
}

func NewRabbitMQConnection(url string) (*RabbitMQConnection, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}
	return &RabbitMQConnection{conn: conn}, nil
}

func (rmq *RabbitMQConnection) Close() {
	if rmq.conn != nil {
		rmq.conn.Close()
	}
}

func (rmq *RabbitMQConnection) Reconnect(url string) error {
	rmq.Close()
	conn, err := NewRabbitMQConnection(url)
	if err != nil {
		return err
	}
	rmq.conn = conn.conn
	return nil
}
