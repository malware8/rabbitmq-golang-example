package repository

import (
	"rabbit-learn/models"
	consumer "rabbit-learn/repository/consumer"
	"rabbit-learn/repository/publisher"

	"github.com/streadway/amqp"
)

type ConsumerI interface {
	Consume()
}

type PublisherI interface {
	Publish(payload models.LogItem)
}

type Repository struct {
	ConsumerI
	PublisherI
}

func NewRepository(rabbitConnection *amqp.Connection) *Repository {
	return &Repository{
		ConsumerI:  consumer.NewConsumer(rabbitConnection),
		PublisherI: publisher.NewPublisher(rabbitConnection),
	}
}
