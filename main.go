package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/karlbehrensg/go-web-server-template/database"
	"github.com/karlbehrensg/go-web-server-template/models"
	"github.com/karlbehrensg/go-web-server-template/routes"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	database.DBConnection()

	database.DB.AutoMigrate(&models.User{}, &models.Task{})

	router := gin.Default()
	routes.ApplyRoutes(router)
	router.Run(":" + port)
}
