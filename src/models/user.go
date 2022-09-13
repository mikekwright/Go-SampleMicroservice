package models

type User struct {
	UserId int32 `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}
