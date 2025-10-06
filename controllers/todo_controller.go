package controllers

import (
	"mini_backend/models"
	"net/http"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	Todos  []models.Todo
	NextID = 1
	mu     sync.Mutex
)

func GetTodos(c *gin.Context) {
	mu.Lock()
	defer mu.Unlock()

	c.JSON(http.StatusOK, Todos)
}

func CreateTodos(c *gin.Context) {
	mu.Lock()
	defer mu.Unlock()

	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error ": err.Error()})
		return
	}

	todo.ID = NextID
	NextID++
	Todos = append(Todos, todo)
	c.JSON(http.StatusOK, todo)
}

func UpdateTodos(c *gin.Context) {
	mu.Lock()
	defer mu.Unlock()

	id, _ := strconv.Atoi(c.Param("id"))

	for i, t := range Todos {
		if t.ID == id {
			Todos[i].Completed = !t.Completed
			c.JSON(http.StatusOK, Todos[i])

			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
}

func DeleteTodos(c *gin.Context) {
	mu.Lock()
	defer mu.Unlock()

	id, _ := strconv.Atoi(c.Param("id"))

	for i, t := range Todos {
		if t.ID == id {
			Todos = append(Todos[:i], Todos[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Todo not Found"})
}
