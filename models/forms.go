package models

type LoginForm struct {
	Username string
	Password string
	HasError bool
	ErrorMsg string
}

type CurrentUser struct {
	ID       uint
	Username string
	Role     Role
}
