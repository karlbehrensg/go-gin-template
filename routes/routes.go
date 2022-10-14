package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/karlbehrensg/go-web-server-template/controllers"
)

func UsersRoutes(router *gin.Engine) {
	users := router.Group("/users")
	users.PUT("/:id", controllers.UpdateUser)
	users.DELETE("/:id", controllers.DeleteUser)
}
