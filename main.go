package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pingpong-pnw/go-backend/database"
	"github.com/pingpong-pnw/go-backend/routes"
)

func main() {

	database.PostgresConnect()

	app := gin.Default()
	routes.Setup(app)
	app.Run(":8000")

}
