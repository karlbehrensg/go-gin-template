package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/karlbehrensg/go-web-server-template/controllers"
)

func AuthRoutes(router *gin.Engine) {
	users := router.Group("/auth")
	users.POST("/signup", controllers.CreateUser)
	users.POST("/login", controllers.Login)
}
