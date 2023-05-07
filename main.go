package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	routes.setUp(app)

}
