package controllers

import (
	"mini_backend/config"
	"mini_backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&user)
	c.JSON(http.StatusOK, user)
}

func GetUser(c *gin.Context) {
	var users []models.User
	config.DB.Preload("Posts").Find(&users)
	c.JSON(http.StatusOK, users)
}
