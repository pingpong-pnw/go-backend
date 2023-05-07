package routes

import (
	"github.com/gofiber/fiber/v2"
)

func setUp(app *fiber.App) {

	app.Post("/", controllors.getHome)

}
