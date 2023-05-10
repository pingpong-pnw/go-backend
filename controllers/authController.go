package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
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
		LastedUpdate: time.Now(),
	}

	database.DB.Create(&userData)

	ctx.JSON(http.StatusAccepted, &userData)
}
