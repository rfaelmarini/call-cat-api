package dto

type LoginCredentials struct {
	Username string `form:"username"`
	Password string `form:"password"`
}
