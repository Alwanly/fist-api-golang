package controllers

import (
	"log"
	"net/http"

	"first-api/src/components"
	"first-api/src/models"

	"github.com/gin-gonic/gin"
)

type NewTodo struct {
	Title string `json:"title" binding:"required"`
	Task  string `json:"task" binding:"required"`
}

func GetTodos(c *gin.Context) {
	var todos []models.Todo
	db, err := components.Database()
	if err != nil {
		log.Println(err)
	}

	if err := db.Find(&todos).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todos)

}

func CreateTodo(c *gin.Context) {
	var todo NewTodo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	NewTodo := models.Todo{Title: todo.Title, Task: todo.Task}
	db, err := components.Database()
	if err != nil {
		log.Println(err)
	}
	if err := db.Create(&NewTodo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, NewTodo)
}

func GetTodo(c *gin.Context) {

	var todo models.Todo

	db, err := components.Database()
	if err != nil {
		log.Println(err)
	}

	if err := db.Where("id = ?", c.Param("id")).First(&todo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "todo not found"})
		return
	}
	c.JSON(http.StatusOK, todo)
}
