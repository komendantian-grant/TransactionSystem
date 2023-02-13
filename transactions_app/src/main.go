package main

import (
	"github.com/gin-gonic/gin"
	"time"
	"fmt"
	"log"
)

import (
	"transactionSystem/models"
	"transactionSystem/controllers"
	"transactionSystem/rabbitmq"
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

	fmt.Println("TestTestTestTestTestTestTestTestTestTestTestTestTestTest!!!")

	rabbitmq.GetChannel()

    rabbitmq.SendMessages("hello")
    rabbitmq.SendMessages("hi")
    rabbitmq.SendMessages("privet")

    consume_channel := make(chan []byte)
    go func() {
        for {
            //fmt.Println("Consumed is:", string(<-consume_channel))
            controllers.WithdrawBalanceReceive(<-consume_channel)
        }
    }()
    go rabbitmq.ConsumeMessages(consume_channel)


	router.Run("0.0.0.0:8080")

}