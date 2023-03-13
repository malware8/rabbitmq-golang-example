package publisher

import (
	"rabbit-learn/utils"

	"github.com/streadway/amqp"
)

type Publisher struct {
	RabbitConnection *amqp.Connection
	RabbitChannel    *amqp.Channel
}

func NewPublisher(rabbitConnection *amqp.Connection) *Publisher {
	publisher := Publisher{RabbitConnection: rabbitConnection}
	ch, err := publisher.RabbitConnection.Channel()
	utils.HandleError(err, "Can't create a amqpChannel")
	publisher.RabbitChannel = ch
	return &publisher
}
