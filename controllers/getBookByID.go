package controllers

import (
	"book-store/db"
	"book-store/models"
	"net/http"

	"errors"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetBookByID(ctx *gin.Context) {
	ID := ctx.Param("id")
	var book models.Book

	result := db.DB.Where("id = ?", ID).First(&book)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "book with id " + ID + " not found"})
		return
	} else if result.Error != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "could not retrieve book", "error": result.Error.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"book": book})
}
