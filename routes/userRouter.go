package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jeevangb/golang-jwt-gin-authenticatin-project/controllers"
	"github.com/jeevangb/golang-jwt-gin-authenticatin-project/middleware"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controllers.GetUsers())
	incomingRoutes.GET("/users/:user_id", controllers.GetUser())
}
