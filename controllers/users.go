package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/karlbehrensg/go-web-server-template/database"
	"github.com/karlbehrensg/go-web-server-template/models"
	"github.com/karlbehrensg/go-web-server-template/schemas"
)

func CreateUser(c *gin.Context) {
	var form schemas.CreateUser
	var user models.User

	if err := c.Bind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if form.Password != form.Password2 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Passwords do not match"})
		return
	}

	user.Username = form.Username
	user.Password = form.Password

	createUser := database.DB.Create(&user)

	if createUser.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": createUser.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, &user)
}

func GetUsers(c *gin.Context) {
	var users []models.User

	getUsers := database.DB.Find(&users)

	if getUsers.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": getUsers.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, &users)
}

func GetUser(c *gin.Context) {
	c.String(http.StatusOK, "Get User")
}

func UpdateUser(c *gin.Context) {
	c.String(http.StatusOK, "Update User")
}

func DeleteUser(c *gin.Context) {
	c.String(http.StatusOK, "Delete User")
}
