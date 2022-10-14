package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/karlbehrensg/go-web-server-template/models"
	"github.com/karlbehrensg/go-web-server-template/schemas"
)

func CreateUser(c *gin.Context) {
	var form schemas.CreateUser
	var response schemas.UserData
	var user models.User

	if err := c.Bind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := user.Register(&form)

	if err != nil {
		if err.Error() == "ERROR: duplicate key value violates unique constraint \"idx_go_gin_users_username\" (SQLSTATE 23505)" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
			return
		} else if err.Error() == "Passwords do not match" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Passwords do not match"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong"})
		return
	}

	response.ID = user.ID
	response.Username = user.Username
	response.Name = user.Name
	response.CreatedAt = user.CreatedAt.String()
	response.UpdatedAt = user.UpdatedAt.String()

	c.JSON(http.StatusCreated, &response)
}

func Login(c *gin.Context) {
	var form schemas.Login
	var response schemas.LoginResponse
	var user models.User

	if err := c.Bind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	access_token, refresh_token, err := user.Login(&form)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid username or password"})
		return
	}

	response.AccessToken = access_token
	response.RefreshToken = refresh_token

	c.JSON(http.StatusOK, &response)
}
