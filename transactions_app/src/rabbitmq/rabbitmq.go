package rabbitmq

import (
	"log"
	"context"
	"time"
	amqp "github.com/rabbitmq/amqp091-go"
)

var CH *amqp.Channel

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func GetChannel(){
	conn, err := amqp.Dial("amqp://rabbitmq:rabbitmq@rabbitmq_app:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")

	CH = ch
}

func ConsumeMessages(consume_channel chan []byte) {
	ch := CH

	q, err := ch.QueueDeclare(
		"withdraw_balance", // name
		true,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	go func() {
		for d := range msgs {
		    consume_channel <- d.Body
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
}

func SendMessages(body string) {
    ch := CH

	q, err := ch.QueueDeclare(
		"withdraw_balance", // name
		true,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//body := "Hello World!"
	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", body)
}
