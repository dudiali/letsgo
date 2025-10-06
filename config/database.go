package config

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect database: ", err)
	}

	DB = database
}
