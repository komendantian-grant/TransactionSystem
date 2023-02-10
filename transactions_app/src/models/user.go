package models

type User struct {
	ID uint `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Balance int `json:"balance"`
}