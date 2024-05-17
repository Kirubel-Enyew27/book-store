package main

import (
	"book-store/controllers"
	"book-store/db"
	"book-store/middlewares"

	"github.com/gin-gonic/gin"
)

func init() {
	db.LoadEnv()
	db.ConnectDB()
}

func main() {
	router := gin.Default()

	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)
	router.POST("/books", middlewares.AuthMiddleware(), controllers.CreateBook)
	router.GET("/books", controllers.GetBooks)
	router.GET("/books/:id", controllers.GetBookByID)
	router.PUT("/books/:id", middlewares.AuthMiddleware(), controllers.UpdateBook)
	router.DELETE("/books/:id", middlewares.AuthMiddleware(), controllers.DeleteBook)

	router.Run(":8080")
}
