package main

import (
	"book-store/db"
	"book-store/models"
)

func init() {
	db.LoadEnv()
	db.ConnectDB()
}

func main() {
	db.DB.AutoMigrate(&models.Book{})
	db.DB.AutoMigrate(&models.User{})
}
