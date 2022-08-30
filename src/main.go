package main

import (
	"first-api/src/components"
	"first-api/src/controllers"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	db, err := components.Database()

	if err != nil {
		log.Println(err)
		return
	}
	db.DB()
	router := gin.Default()
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	router.GET("/todos", controllers.GetTodos)
	router.POST("/todo", controllers.CreateTodo)
	router.GET("/todos/:id", controllers.GetTodo)
	router.Run()
}
