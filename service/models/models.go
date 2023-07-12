package models

type User struct {
	UserId   int64  `json:"user_id" db:"user_id"`
	Username string `json:"username" db:"username" validate:"required , min=2 , max=100"`
	Password string `json:"password" db:"password" validate:"required, min=6"`
}
