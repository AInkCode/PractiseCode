package routes

import (
	"apiserverv2/controller"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	routes := gin.Default()
	rGroup := routes.Group("api/v1")
	rGroup.POST("/login", controller.SignupHandler)
	return routes
}
