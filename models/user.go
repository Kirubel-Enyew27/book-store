package models

type User struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	Password string `json:"password"`
}
