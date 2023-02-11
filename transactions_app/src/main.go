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

	rabbitmq.GetChannel()

    rabbitmq.SendMessages()
    rabbitmq.SendMessages()
    rabbitmq.SendMessages()

    go rabbitmq.ConsumeMessages()


	router.Run("0.0.0.0:8080")

}