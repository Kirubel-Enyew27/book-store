package main

import (
	"book-store/db"

	"github.com/gin-gonic/gin"
)

func init() {
	db.LoadEnv()
	db.ConnectDB()
}

func main() {
	router := gin.Default()

	router.Run(":8080")
}
