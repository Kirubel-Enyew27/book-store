package main

import (
	"book-store/controllers"
	"book-store/db"

	"github.com/gin-gonic/gin"
)

func init() {
	db.LoadEnv()
	db.ConnectDB()
}

func main() {
	router := gin.Default()

	router.POST("/books", controllers.CreateBook)

	router.Run(":8080")
}
