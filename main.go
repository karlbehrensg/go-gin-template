package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/karlbehrensg/go-web-server-template/routes"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := gin.Default()
	routes.ApplyRoutes(router)
	router.Run(":" + port)
}
