package main

import (
	"rabbit-learn/models"
	repository "rabbit-learn/repository"
	"rabbit-learn/utils"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

func main() {

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	utils.HandleError(err, "Can't connect to AMQP")

	defer conn.Close()

	repos := repository.NewRepository(conn)

	go repos.Consume()
	time.Sleep(2 * time.Second)

	logrus.Info("Starting publishers")

	repos.Publish(models.LogItem{
		RequestID:   "system",
		ServiceName: "rabbit-learn",
		LogType:     "info",
		Info:        "1",
	})

	repos.Publish(models.LogItem{
		RequestID:   "system",
		ServiceName: "rabbit-learn",
		LogType:     "info",
		Info:        "2",
	})

	repos.Publish(models.LogItem{
		RequestID:   "system",
		ServiceName: "rabbit-learn",
		LogType:     "info",
		Info:        "3",
	})

	repos.Publish(models.LogItem{
		RequestID:   "system",
		ServiceName: "rabbit-learn",
		LogType:     "info",
		Info:        "4",
	})
}
