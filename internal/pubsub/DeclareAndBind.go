package pubsub

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type simpleQueueType string

const (
	Durable   simpleQueueType = "durable"
	Transient simpleQueueType = "transient"
)

func DeclareAndBind(
	conn *amqp.Connection,
	exchange,
	queueName,
	key string,
	queueType simpleQueueType, // an enum to represent "durable" or "transient"
) (*amqp.Channel, amqp.Queue, error) {
	cha, err := conn.Channel()
	if err != nil {
		return nil, amqp.Queue{}, err
	}
	durable := queueType == "durable"
	autoDelete := queueType == "transient"
	exclusive := queueType == "transient"
	q, err := cha.QueueDeclare(queueName, durable, autoDelete, exclusive, false, nil)
	if err != nil {
		return nil, amqp.Queue{}, err
	}
	err = cha.QueueBind(queueName, key, exchange, false, nil)
	if err != nil {
		return nil, amqp.Queue{}, err
	}
	return cha, q, nil
}
