package controllers

import (
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pingpong-pnw/go-backend/database"
	"github.com/pingpong-pnw/go-backend/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
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
			"message": "Unknown error please try again later",
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

	dbResult := database.DB.Create(&userData)
	if errors.Is(dbResult.Error, gorm.ErrDuplicatedKey) {
		ctx.AbortWithStatusJSON(http.StatusMethodNotAllowed, gin.H{
			"message": "This registered email already exists. Please try a different email.",
		})
		return
	}

	ctx.JSON(http.StatusOK, &userData)
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

	dbResult := database.DB.Where("email = ?", data.Email).First(&user)

	if errors.Is(dbResult.Error, gorm.ErrRecordNotFound) {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized, your email are not exists. Please register an account.",
		})
		return
	}

	// Compare hash password in returned data with logged in password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password)); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized, password has incorrect. Please use forgot password.",
			"error":   err.Error(),
		})
		return
	}

	if data.RememberMe {
		// Create JWT token to save session for user if user check remember me
		claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"Issuer":    user.Id.String(),
			"ExpiresAt": time.Now().Add(time.Hour * 24).Unix(),
		})
		token, err := claims.SignedString([]byte(os.Getenv("SECRET_KEY")))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"message": "Could not login. Please try again later.",
				"error":   err.Error(),
			})
		}

		ctx.SetSameSite(http.SameSiteLaxMode)
		ctx.SetCookie("Authorization", token, (3600 * 24), "", "", false, true)

		ctx.JSON(http.StatusOK, gin.H{
			"message": "Login successful",
		})
	} else {
		// Return message success to login
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Login successful",
		})
	}
}
