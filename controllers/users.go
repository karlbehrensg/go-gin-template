package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	password2 := c.PostForm("password2")

	fmt.Printf("username: %s; password: %s; password2: %s;\n", username, password, password2)

	c.String(http.StatusOK, "Hello, World!")
}
