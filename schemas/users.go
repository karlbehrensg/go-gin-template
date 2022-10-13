package schemas

type CreateUser struct {
	Username  string `form:"username" binding:"required,email"`
	Password  string `form:"password" binding:"required"`
	Password2 string `form:"password2" binding:"required"`
}
