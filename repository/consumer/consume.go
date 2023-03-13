package consumer

import (
	"encoding/json"
	"os"
	"rabbit-learn/models"
	"rabbit-learn/utils"

	"github.com/sirupsen/logrus"
)

func (r *Consumer) Consume() {

	queue, err := r.RabbitChannel.QueueDeclare("logs", true, false, false, false, nil)
	utils.HandleError(err, "Could't declare `logs` queue")

	messageChannel, err := r.RabbitChannel.Consume(
		queue.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	utils.HandleError(err, "Could't register consumer")

	stopChan := make(chan bool)

	go func() {
		logrus.Infof("Consumer ready, PID: %d", os.Getpid())
		for d := range messageChannel {
			logrus.Info("Received message")
			var messagePayload *models.LogItem
			err := json.Unmarshal(d.Body, &messagePayload)
			if err != nil {
				logrus.Infof("Error decoding JSON: %s", err)
			}

			if messagePayload != nil {
				logrus.Info("Message info - ", messagePayload.Info)
				logrus.Info("Message request id  - ", messagePayload.RequestID)
				logrus.Info("Message log type  - ", messagePayload.LogType)
				logrus.Info("Message service name  - ", messagePayload.ServiceName)
			} else {
				logrus.Info("Message is nil")
			}

			if err := d.Ack(false); err != nil {
				logrus.Info("Error acknowledging message: %s", err)
			} else {
				logrus.Info("Acknowledged message")
			}
		}
	}()

	<-stopChan
}
