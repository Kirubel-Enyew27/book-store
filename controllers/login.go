package controllers

import (
	"book-store/db"
	"book-store/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func Login(ctx *gin.Context) {
	var userField models.User

	if err := ctx.ShouldBindJSON(&userField); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "error parsing request body"})
		return
	}

	var user models.User

	result := db.DB.Where("username=?", userField.Username).First(&user)
	if result.Error != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"message": result.Error.Error()})
		return
	}

	if user.UserID == 0 || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userField.Password)) != nil {
		ctx.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "invalid username or password"})
		return
	}

	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       userField.UserID,
		"username": userField.Username,
		"exp":      time.Now().Add(time.Hour).Unix(),
	})

	token, err := generateToken.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		ctx.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "failed to generate token"})
		return
	}
	ctx.SetCookie("token", token, int(time.Hour.Seconds()), "/", "", false, true)
	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "user loged in successfully"})
}
