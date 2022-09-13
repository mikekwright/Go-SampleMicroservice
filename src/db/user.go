package db

import (
	"SampleGoService/src/db/entity"
)

type User struct {
}

func (db User) FindAllUsers() []entity.User {
	return []entity.User {
		{UserId: 1, Name: "User 1", Email: "test+1@test.com"},
		{UserId: 2, Name: "User 2", Email: "test+2@test.com"},
		{UserId: 3, Name: "User 3", Email: "test+3@test.com"},
	}
}
