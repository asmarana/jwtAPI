package models

type User struct {
	Username string `json:"username" validate:"required , min=2 , max=100"`
	Password string `json:"password" validate:"required, min=6"`
}
