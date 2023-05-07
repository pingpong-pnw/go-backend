package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pingpong-pnw/go-backend/routes"
)

func main() {
	app := fiber.New()
	routes.Setup(app)
	app.Listen(":8080")
}
