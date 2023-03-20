package main

import (
	"Go-Tutorials/controllers"
	"Go-Tutorials/models"
	"github.com/gin-gonic/gin"
)

func main() {
	// ...
	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.FindBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/books/:id", controllers.DeleteBook) // new

	err := r.Run()
	if err != nil {
		return
	}
}
