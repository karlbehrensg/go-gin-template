package main

import (
	"github.com/gin-gonic/gin"
	"github.com/karlbehrensg/go-web-server-template/routes"
)

func main() {
	router := gin.Default()
	routes.ApplyRoutes(router)
	router.Run()
}
