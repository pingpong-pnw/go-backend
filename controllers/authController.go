package controllers

import "github.com/gofiber/fiber/v2"

func Register(ctx *fiber.Ctx) error {
	var data map[string]string
	if err := ctx.BodyParser(&data); err != nil {
		return err
	}
	return ctx.JSON(data)
}
