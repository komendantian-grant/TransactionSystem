package main

import (
	"github.com/gin-gonic/gin"
)

import (
	"transactionSystem/models"
	"transactionSystem/controllers"
)

func main() {

	models.ConnectDatabase()

	router := gin.Default()

	router.GET("/get_users", controllers.GetUsers)
	router.GET("/get_user", controllers.GetUser)
	router.POST("/create_user", controllers.CreateUser)
	router.PATCH("/update_user", controllers.UpdateUser)
	router.DELETE("/delete_user", controllers.DeleteUser)
	router.POST("/withdraw_balance", controllers.WithdrawBalance)

	router.Run("0.0.0.0:8080")

}