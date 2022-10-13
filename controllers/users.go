package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/karlbehrensg/go-web-server-template/schemas"
)

func CreateUser(c *gin.Context) {
	var form schemas.CreateUser

	if err := c.Bind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Printf("username: %s; password: %s; password2: %s;\n", form.Username, form.Password, form.Password2)

	c.String(http.StatusOK, "Hello, World!")
}
