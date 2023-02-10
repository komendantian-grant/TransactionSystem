package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := "host=postgres_app user=postgres password=pass123 dbname=accounts port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Unable to connect to database.")
	}

	database.AutoMigrate(&User{})

	DB = database
}