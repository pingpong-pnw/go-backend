package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/pingpong-pnw/go-backend/controllers"
)

func Setup(app *gin.Engine) {

	app.POST("/api/v1/register", controllers.Register)
	app.POST("/api/v1/login", controllers.Login)
	app.POST("/api/v1/logout", controllers.Logout)

}
