package controllers

import (
	"mini_backend/config"
	"mini_backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	var post models.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&post)
	c.JSON(http.StatusOK, post)
}

func GetPosts(c *gin.Context) {
	var posts []models.Post
	config.DB.Find(&posts)
	c.JSON(http.StatusOK, posts)
}
