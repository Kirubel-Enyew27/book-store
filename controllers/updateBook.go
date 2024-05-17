package controllers

import (
	"book-store/db"
	"book-store/models"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UpdateBook(ctx *gin.Context) {
	ID := ctx.Param("id")
	var book models.Book

	result := db.DB.Where("id = ?", ID).First(&book)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "book with id " + ID + " not found"})
		return
	} else if result.Error != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	var updatedData models.Book
	if err := ctx.ShouldBindJSON(&updatedData); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "error parsing request body", "error": err.Error()})
		return
	}

	result = db.DB.Model(&book).Updates(updatedData)
	if result.Error != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "unable to update book with id " + ID, "error": result.Error.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"updated book": book})
}
