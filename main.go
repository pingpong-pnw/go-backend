package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pingpong-pnw/go-backend/database"
	"github.com/pingpong-pnw/go-backend/routes"
)

func main() {

	database.PostgresConnect()

	app := fiber.New()
	routes.Setup(app)
	app.Listen(":8000")
}
