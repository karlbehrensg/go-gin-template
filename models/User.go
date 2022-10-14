package models

import (
	"context"
	"errors"
	"time"

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

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		return err
	}

	user.Password = string(bytes)
	return nil
}

func (user *User) Register(form *schemas.CreateUser, db *gorm.DB) error {
	if form.Password != form.Password2 {
		return errors.New("Passwords do not match")
	}

	user.Username = form.Username
	err := user.HashPassword(form.Password)

	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	createUser := db.WithContext(ctx).Create(&user)

	if createUser.Error != nil {
		return createUser.Error
	}

	return nil
}

func (user *User) GetUser(username string, db *gorm.DB) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return db.WithContext(ctx).Where("username = ?", username).First(&user).Error
}

func (user *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}

func (user *User) CreateTokens() (string, string) {
	access_token_payload := &schemas.JWTPayload{
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
		},
	}
	refresh_token_payload := &schemas.JWTPayload{
		UserID: user.ID,
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

func (user *User) Login(form *schemas.Login, db *gorm.DB) (string, string, error) {
	err := user.GetUser(form.Username, db)

	if err != nil {
		return "", "", err
	}

	if !user.ValidatePassword(form.Password) {
		return "", "", errors.New("Invalid password")
	}

	access_token, refresh_token := user.CreateTokens()

	return access_token, refresh_token, nil
}
