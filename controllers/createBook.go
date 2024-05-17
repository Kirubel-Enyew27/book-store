package controllers

import (
	"book-store/db"
	"book-store/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateBook(ctx *gin.Context) {
	var bookField models.Book
	if err := ctx.ShouldBindJSON(&bookField); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "error parsing JSON", "error": err.Error()})
		return
	}

	var book models.Book
	db.DB.Where("id = ?", bookField.ID).Find(&book)
	if book.ID != 0 {
		ctx.JSON(http.StatusConflict, gin.H{"message": "the book already posted"})
		return
	}

	if err := db.DB.Create(&bookField).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not create book", "error": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, gin.H{"Book": bookField, "message": "the book posted successfully"})
}
