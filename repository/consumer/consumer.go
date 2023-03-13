package consumer

import (
	"rabbit-learn/utils"

	"github.com/streadway/amqp"
)

type Consumer struct {
	RabbitConnection *amqp.Connection
	RabbitChannel    *amqp.Channel
}

func NewConsumer(rabbitConnection *amqp.Connection) *Consumer {
	consumer := Consumer{RabbitConnection: rabbitConnection}
	ch, err := consumer.RabbitConnection.Channel()
	utils.HandleError(err, "Can't create a amqpChannel")
	consumer.RabbitChannel = ch
	return &consumer
}
