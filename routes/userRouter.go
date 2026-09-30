package routes

import (
	"golang_jwt_auth/middleware"

	"github.com/gin-gonic/gin"
	controller "github.com/satya035/golang_jwt_auth/controllers"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
}
