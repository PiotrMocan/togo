package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	todo "togoapi.com/internal/todo"
)

func main() {
	router := gin.Default()

	router.Use(cors.Default())

	config := cors.Config{
		AllowOrigins:     []string{"*"}, // Your frontend's origin
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}
	router.Use(cors.New(config))

	router.GET("/todos", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"todos": todo.TodoList(),
		})
	})

	println("Server is running on port 8080")

	router.Run()
}
