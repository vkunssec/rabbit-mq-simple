package main

import (
	"os"

	"github.com/vkunssec/rabbit-mq-simple/internal/config/environment"
	"github.com/vkunssec/rabbit-mq-simple/pkg/repository"
)

func main() {
	environment.LoadEnv()
	amqpServerURL := os.Getenv("AMQP_SERVER_URL")

	_, channelRabbitMQ, err := repository.OpenConnectionRabbitMQ(amqpServerURL)
	if err != nil {
		panic(err)
	}

	repository.ReceiveMessageRabbitMQ(channelRabbitMQ)
}
