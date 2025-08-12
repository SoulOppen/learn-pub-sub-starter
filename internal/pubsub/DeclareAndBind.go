package pubsub

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type SimpleQueueType string

const (
	Durable   SimpleQueueType = "durable"
	Transient SimpleQueueType = "transient"
)

func DeclareAndBind(
	conn *amqp.Connection,
	exchange,
	queueName,
	key string,
	queueType SimpleQueueType, // an enum to represent "durable" or "transient"
) (*amqp.Channel, amqp.Queue, error) {
	cha, err := conn.Channel()
	if err != nil {
		return nil, amqp.Queue{}, err
	}
	durable := queueType == "durable"
	autoDelete := queueType == "transient"
	exclusive := queueType == "transient"
	table := make(amqp.Table, 1)
	table["x-dead-letter-exchange"] = "peril_dlx"
	q, err := cha.QueueDeclare(queueName, durable, autoDelete, exclusive, false, table)
	if err != nil {
		return nil, amqp.Queue{}, err
	}
	err = cha.QueueBind(queueName, key, exchange, false, nil)
	if err != nil {
		return nil, amqp.Queue{}, err
	}
	return cha, q, nil
}
