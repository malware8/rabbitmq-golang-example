
Before starting example, you need docker container

```
docker run --rm -p 15672:15672 -p 5672:5672 rabbitmq:3.10.7-management
```

Run

```
go run main.go
```

Resources used
* https://medium.com/nuances-of-programming/рабочая-очередь-в-go-с-rabbitmq-d543183028c

Used libs
* github.com/streadway/amqp
* github.com/sirupsen/logrus