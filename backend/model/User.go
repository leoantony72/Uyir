package model

type User struct {
	Id       string `json:"id" redis:"id"`
	Name     string `json:"username" redis:"username"`
	Email    string `json:"email" redis:"email"`
	Password string `json:"password" redis:"password"`
}
