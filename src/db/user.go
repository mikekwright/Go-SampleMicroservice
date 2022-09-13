package db

import (
	"SampleGoService/src/db/entity"
)

type UserDao struct {
}

func NewUserDao() UserDao {
	return UserDao{}
}

func (db UserDao) FindAllUsers() []entity.UserEntity {
	return []entity.UserEntity{
		{UserId: 1, Name: "UserEntity 1", Email: "test+1@test.com"},
		{UserId: 2, Name: "UserEntity 2", Email: "test+2@test.com"},
		{UserId: 3, Name: "UserEntity 3", Email: "test+3@test.com"},
	}
}
