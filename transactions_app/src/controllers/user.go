package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

import (
	"transactionSystem/models"
)


type WithdrawBalanceInput struct {
	Id   int `json:"id" binding:"required"`
	Amount int `json:"amount" binding:"required"`
}

// curl -v -H 'Content-Type: application/json' -d '{"id":2, "amount":16}' -X POST 127.0.0.1:8080/change_balance
func WithdrawBalance(c *gin.Context) {
	var input WithdrawBalanceInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := models.DB.Where("id = ?", input.Id).First(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	if input.Amount > user.Balance {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Insufficient funds"})
	}

	balance_withdrawn := input.Amount
	new_balance := user.Balance - balance_withdrawn
	updatedUser := models.User{Balance: new_balance}

	models.DB.Model(&user).Updates(&updatedUser)

	c.JSON(http.StatusOK, gin.H{"data": user})

}


// curl -v -H 'Content-Type: application/json' -d '{}' -X GET 127.0.0.1:8080/get_users
func GetUsers(c *gin.Context) {
	var users []models.User
	models.DB.Find(&users)
	c.JSON(http.StatusOK, gin.H{"data": users})
}

type GetUserInput struct {
	Id int `json:"id" binding:"required"`
}

// curl -v -H 'Content-Type: application/json' -d '{"id": 3}' -X GET 127.0.0.1:8080/get_user
func GetUser(c *gin.Context) {
	var input GetUserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := models.DB.Where("id = ?", input.Id).First(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}


type CreateUserInput struct {
	Name string `json:"name" binding:"required"`
	Balance int `json:"balance" binding:"required"`
}

// curl -v -H 'Content-Type: application/json' -d '{"name":"Petrov Petr Petrovich", "balance":16}' -X POST 127.0.0.1:8080/create_user
func CreateUser(c *gin.Context) {
	var input CreateUserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{Name: input.Name, Balance: input.Balance}
	models.DB.Create(&user)
	c.JSON(http.StatusOK, gin.H{"data": user})
}

type UpdateUserInput struct {
	Id int `json:"id" binding:"required"`
	Name string `json:"name"`
	Balance int `json:"balance"`
}


// curl -v -H 'Content-Type: application/json' -d '{"id": 2, "name":"Petrov Petr Petrovich", "balance":16}' -X PATCH 127.0.0.1:8080/update_user
func UpdateUser(c *gin.Context) {

	var input UpdateUserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := models.DB.Where("id = ?", input.Id).First(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	updatedUser := models.User{Name: input.Name, Balance: input.Balance}
	models.DB.Model(&user).Updates(&updatedUser)
	c.JSON(http.StatusOK, gin.H{"data": user})

}

type DeleteUserInput struct {
	Id int `json:"id" binding:"required"` 
}

//curl -v -H 'Content-Type: application/json' -d '{"id": 3}' -X DELETE 127.0.0.1:8080/delete_user
func DeleteUser(c *gin.Context) {
	var input DeleteUserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	if err := models.DB.Where("id = ?", input.Id).First(&user).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	models.DB.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"data": "success"})
}



