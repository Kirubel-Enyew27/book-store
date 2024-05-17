package controllers

import (
	"book-store/db"
	"book-store/models"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteBook(ctx *gin.Context) {
	ID := ctx.Param("id")
	var book models.Book

	result := db.DB.Where("id = ?", ID).First(&book)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "book with id " + ID + " not found"})
		return
	} else if result.Error != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": result.Error.Error()})
		return
	}

	result = db.DB.Delete(&book)
	if result.Error != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "unable to delete book with id " + ID, "error": result.Error.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "book with id " + ID + " deleted successfully"})
}
