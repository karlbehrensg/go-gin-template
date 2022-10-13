package schemas

type CreateUser struct {
	Username  string `form:"username" binding:"required,email"`
	Password  string `form:"password" binding:"required"`
	Password2 string `form:"password2" binding:"required"`
}

type UserData struct {
	ID        uint   `json:"id"`
	Username  string `json:"username"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UpdateUser struct {
	Name string `json:"name" binding:"required"`
}
