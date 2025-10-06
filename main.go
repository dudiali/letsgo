package main

import (
	"mini_backend/config"
	"mini_backend/models"
	"mini_backend/routes"
)

func main() {
	config.ConnectDB()

	config.DB.AutoMigrate(&models.User{}, &models.Post{})

	r := routes.SetupRouter()

	r.Run(":8080")
}
