package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pingpong-pnw/go-backend/database"
	"github.com/pingpong-pnw/go-backend/models"
	"golang.org/x/crypto/bcrypt"
)

func Register(ctx *gin.Context) {
	var data models.Users
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid inputs, Please check your inputs",
			"error":   err.Error(),
		})
		return
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), 12)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"message": "Unknow error please try again later",
			"error":   err.Error(),
		})
		return
	}

	userData := models.Users{
		FirstName:    data.FirstName,
		LastName:     data.LastName,
		Email:        data.Email,
		Password:     string(hashPassword),
		LastedUpdate: time.Now().Unix(),
	}

	database.DB.Create(&userData)

	ctx.JSON(http.StatusOK, userData)
}

func Login(ctx *gin.Context) {

	var data models.Login
	var user models.Users
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid inputs, Please check your inputs",
			"error":   err.Error(),
		})
		return
	}

	database.DB.Where("email = ?", data.Email).First(&user)

	if user.Id == uuid.Nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized, your Email are not exists. Please register an account.",
		})
		return
	}

	// Compare Hash Password in returned data with Logged in Password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password)); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized, Password has incorrect. Please use forgot password.",
			"error":   err.Error(),
		})
		return
	}

	// return Login Success message and user data to client
	ctx.JSON(http.StatusOK, user)
}
