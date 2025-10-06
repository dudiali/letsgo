package controllers

import (
	"mini_backend/config"
	"mini_backend/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {
	var todos []models.Todo
	config.DB.Find(&todos)
	c.JSON(http.StatusOK, todos)
}

func CreateTodos(c *gin.Context) {
	var todo models.Todo

	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&todo)
	c.JSON(http.StatusOK, todo)
}

func UpdateTodos(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var todo models.Todo

	if err := config.DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	todo.Completed = !todo.Completed
	config.DB.Save(&todo)

	c.JSON(http.StatusOK, todo)
}

func DeleteTodos(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var todo models.Todo

	if err := config.DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	config.DB.Delete(&todo)
	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
}
