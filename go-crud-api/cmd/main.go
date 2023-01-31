package main

import (
	"fmt"
	"github.com/Lefree111/go-gin-rest-api/go-crud-api/controller"
	"github.com/Lefree111/go-gin-rest-api/go-crud-api/database"

	"github.com/gin-gonic/gin"
)

func main() {

	fmt.Println("Starting application ...")
	database.DatabaseConnection()

	r := gin.Default()
	r.GET("/books/:id", controller.ReadBook)
	r.GET("/books", controller.ReadBooks)
	r.POST("/books", controller.CreateBook)
	r.PUT("/books/:id", controller.UpdateBook)
	r.DELETE("/books/:id", controller.DeleteBook)
	err := r.Run(":3000")
	if err != nil {
		return
	}
}
