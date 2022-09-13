//+build wireinject

package main

import (
	"github.com/google/wire"

	"SampleGoService/src/controllers"
	"SampleGoService/src/db"
)

func CreateUserController() controllers.UserController {
	wire.Build(
		db.NewUserDao,
		controllers.NewUserController,
	)
	return controllers.UserController{}
}
