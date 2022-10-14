package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/karlbehrensg/go-web-server-template/database"
	"github.com/karlbehrensg/go-web-server-template/models"
	"github.com/karlbehrensg/go-web-server-template/schemas"
)

func UpdateUser(c *gin.Context) {
	var form schemas.UpdateUser

	if err := c.Bind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User

	getUser := database.DB.First(&user, c.Param("id"))

	if getUser.Error != nil {
		if getUser.Error.Error() == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": getUser.Error.Error()})
	}

	user.Name = form.Name

	updateUser := database.DB.Save(&user)

	if updateUser.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": updateUser.Error.Error()})
		return
	}

	var response schemas.UserData

	response.ID = user.ID
	response.Username = user.Username
	response.Name = user.Name
	response.CreatedAt = user.CreatedAt.String()
	response.UpdatedAt = user.UpdatedAt.String()

	c.JSON(http.StatusOK, &response)

}

func DeleteUser(c *gin.Context) {
	var user models.User

	getUser := database.DB.First(&user, c.Param("id"))

	if getUser.Error != nil {
		if getUser.Error.Error() == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": getUser.Error.Error()})
	}

	deleteUser := database.DB.Delete(&user)

	if deleteUser.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": deleteUser.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}
