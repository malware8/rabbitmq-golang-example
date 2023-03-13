package publisher

import (
	"encoding/json"
	"rabbit-learn/models"
	"rabbit-learn/utils"

	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

func (r *Publisher) Publish(payload models.LogItem) {

	queue, err := r.RabbitChannel.QueueDeclare("logs", true, false, false, false, nil)
	utils.HandleError(err, "Could't declare `logs` queue")

	body, err := json.Marshal(payload)
	if err != nil {
		utils.HandleError(err, "Error encoding JSON")
	}

	err = r.RabbitChannel.Publish("", queue.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "text/plain",
		Body:         []byte(body),
	})
	if err != nil {
		logrus.Fatalf("Error publishing message: %s", err)
	}
	logrus.Infof("Published payload: ", payload)
}
