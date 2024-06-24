package model

type RegisterForm struct {
	UserName        string `json:"username"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

type User struct {
	UserID   string `json:"user_id" db:"user_id"`
	UserName string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}
