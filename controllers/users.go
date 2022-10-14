package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/karlbehrensg/go-web-server-template/database"
	"github.com/karlbehrensg/go-web-server-template/models"
	"github.com/karlbehrensg/go-web-server-template/schemas"
)

func UpdateUser(c *gin.Context) {
	var form schemas.UpdateUser
	var user models.User

	if err := c.Bind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	access_token := c.GetHeader("Authorization")

	if access_token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Access token is required"})
		return
	}

	access_token = strings.Split(access_token, "Bearer ")[1]

	if err := user.Update(&form, access_token); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response := &schemas.UserData{
		ID:        user.ID,
		Username:  user.Username,
		Name:      user.Name,
		CreatedAt: user.CreatedAt.String(),
		UpdatedAt: user.UpdatedAt.String(),
	}

	c.JSON(http.StatusOK, &response)

}

func DeleteUser(c *gin.Context) {
	var user models.User

	access_token := c.GetHeader("Authorization")

	if access_token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Access token is required"})
		return
	}

	access_token = strings.Split(access_token, "Bearer ")[1]

	if err := user.Delete(access_token); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Delete(&user)

	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}
