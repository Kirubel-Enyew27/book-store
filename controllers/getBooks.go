package controllers

import (
	"book-store/db"
	"book-store/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBooks(ctx *gin.Context) {
	var books []models.Book

	result := db.DB.Find(&books)
	if result.Error != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "could not retrieve books", "error": result.Error.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"books": books})
}
