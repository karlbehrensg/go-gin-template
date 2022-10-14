package schemas

import "github.com/golang-jwt/jwt"

type CreateUser struct {
	Username  string `form:"username" binding:"required,email"`
	Password  string `form:"password" binding:"required"`
	Password2 string `form:"password2" binding:"required"`
}

type Login struct {
	Username string `form:"username" binding:"required,email"`
	Password string `form:"password" binding:"required"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type JWTPayload struct {
	jwt.StandardClaims
	UserID uint
}
