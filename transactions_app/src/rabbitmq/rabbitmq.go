package rabbitmq

import (
	"log"
	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func GetChannel() *amqp.Channel {
	conn, err := amqp.Dial("amqp://rabbitmq:rabbitmq@rabbitmq_app:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")

	return ch
}