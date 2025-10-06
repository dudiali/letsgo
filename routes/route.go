package routes

import (
	"mini_backend/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.Default())

	r.POST("/users", controllers.CreateUser)
	r.GET("/users", controllers.GetUser)

	r.POST("/posts", controllers.CreatePost)
	r.GET("/posts", controllers.GetPosts)

	r.GET("todos", controllers.GetTodos)
	r.POST("todos", controllers.CreateTodos)
	r.PUT("todos/:id", controllers.UpdateTodos)
	r.DELETE("todos/:id", controllers.DeleteTodos)

	return r
}
