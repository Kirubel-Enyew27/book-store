package controllers

import (
	"book-store/db"
	"book-store/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(ctx *gin.Context) {
	var userField models.User
	if err := ctx.ShouldBindJSON(&userField); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "error parsing request body"})
		return
	}

	if userField.Username == "" || userField.Password == "" {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "username and password can not be empty"})
		return
	}

	var user models.User

	db.DB.Where("username=?", userField.Username).First(&user)

	if user.UserID != 0 {
		ctx.IndentedJSON(http.StatusConflict, gin.H{"message": "username already exists"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userField.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "error hashing a password", "error": err.Error()})
		return
	}

	userData := models.User{
		UserID:   userField.UserID,
		Username: userField.Username,
		Password: string(hashedPassword),
	}

	db.DB.Create(&userData)
	ctx.IndentedJSON(http.StatusCreated, gin.H{"user": userData, "message": "user registered successfully"})
}
