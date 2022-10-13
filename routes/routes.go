package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/karlbehrensg/go-web-server-template/controllers"
)

func ApplyRoutes(router *gin.Engine) {
	users := router.Group("/users")
	{
		users.POST("", controllers.CreateUser)
	}
}
