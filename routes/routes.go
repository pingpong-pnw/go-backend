package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pingpong-pnw/go-backend/controllers"
)

func Setup(app *fiber.App) {

	app.Post("/api/v1/register", controllers.Register)

}
