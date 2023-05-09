package controllers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/pingpong-pnw/go-backend/database"
	"github.com/pingpong-pnw/go-backend/models"
	"golang.org/x/crypto/bcrypt"
)

func Register(ctx *fiber.Ctx) error {
	var data map[string]string
	if err := ctx.BodyParser(&data); err != nil {
		return err
	}

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 12)
	userData := models.Users{
		FirstName:    data["firstName"],
		LastName:     data["lastName"],
		Email:        data["email"],
		Password:     hashPassword,
		LastedUpdate: time.Now(),
	}

	database.DB.Create(&userData)

	return ctx.JSON(userData)
}
