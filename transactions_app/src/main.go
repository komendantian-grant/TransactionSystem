package main

import (
	"github.com/gin-gonic/gin"
	"time"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

import (
	"transactionSystem/models"
	"transactionSystem/controllers"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {

    time.Sleep(time.Second * 5)

	models.ConnectDatabase()

	router := gin.Default()

	router.GET("/get_users", controllers.GetUsers)
	router.GET("/get_user", controllers.GetUser)
	router.POST("/create_user", controllers.CreateUser)
	router.PATCH("/update_user", controllers.UpdateUser)
	router.DELETE("/delete_user", controllers.DeleteUser)
	router.POST("/withdraw_balance", controllers.WithdrawBalance)

	ticker := time.NewTicker(time.Second*5)

	fmt.Println("TestTestTestTestTestTestTestTestTestTestTestTestTestTest!!!")

	go func() {
	    for {
	        select {
	        case t:= <- ticker.C:
	            fmt.Println("Ticker ticked!", t)
	        }
	    }
	}()

	conn, err := amqp.Dial("amqp://rabbitmq:rabbitmq@rabbitmq_app:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
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
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")


	router.Run("0.0.0.0:8080")

}