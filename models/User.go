package models

import (
	"context"
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/karlbehrensg/go-web-server-template/schemas"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Username string `gorm:"not null;uniqueIndex" json:"username"`
	Name     string `json:"name"`
	Password string `gorm:"not null"`
}

func (u *User) GetUser(username string, db *gorm.DB) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return db.WithContext(ctx).Where("username = ?", username).First(&u).Error
}

func (u *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func (u *User) CreateTokens() (string, string) {
	access_token_payload := &schemas.JWTPayload{
		UserID: u.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
		},
	}
	refresh_token_payload := &schemas.JWTPayload{
		UserID: u.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 7).Unix(),
		},
	}

	access_token := jwt.NewWithClaims(jwt.SigningMethodHS256, access_token_payload)
	refresh_token := jwt.NewWithClaims(jwt.SigningMethodHS256, refresh_token_payload)

	access_token_string, _ := access_token.SignedString([]byte("secret"))
	refresh_token_string, _ := refresh_token.SignedString([]byte("secret"))

	return access_token_string, refresh_token_string
}

func (u *User) Login(form *schemas.Login, c *gin.Context, db *gorm.DB) (string, string, error) {
	err := u.GetUser(form.Username, db)

	if err != nil {
		return "", "", err
	}

	if !u.ValidatePassword(form.Password) {
		return "", "", errors.New("Invalid password")
	}

	access_token, refresh_token := u.CreateTokens()

	return access_token, refresh_token, nil
}
