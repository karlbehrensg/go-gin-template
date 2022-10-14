package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/karlbehrensg/go-web-server-template/database"
	"github.com/karlbehrensg/go-web-server-template/models"
	"github.com/karlbehrensg/go-web-server-template/schemas"
	"golang.org/x/crypto/bcrypt"
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

	bytes, err := bcrypt.GenerateFromPassword([]byte(form.Password), 14)

	if err != nil {
		log.Println("Error hashing password")
		c.JSON(http.StatusInternalServerError, "Something went wrong")
		return
	}

	user.Username = form.Username
	user.Password = string(bytes)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	createUser := database.DB.WithContext(ctx).Create(&user)

	if createUser.Error != nil {
		log.Println(createUser.Error)
		if createUser.Error.Error() == "ERROR: duplicate key value violates unique constraint \"idx_go_gin_users_username\" (SQLSTATE 23505)" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": "Something went wrong"})
		return
	}

	var response schemas.UserData

	response.ID = user.ID
	response.Username = user.Username
	response.Name = user.Name
	response.CreatedAt = user.CreatedAt.String()
	response.UpdatedAt = user.UpdatedAt.String()

	c.JSON(http.StatusCreated, &response)

}
