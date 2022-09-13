package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"SampleGoService/src/db"
	"SampleGoService/src/models"
)

type UserController struct {
	UserDao db.UserDao
}

func NewUserController(userDao db.UserDao) UserController {
	return UserController{UserDao: userDao}
}

func (controller UserController) RegisterController(router *gin.Engine) {
	router.GET("/user", controller.getUsers)
}

// @BasePath /user

// getUsers godoc
// @Summary Get all users in example service
// @Schemes
// @Description list users
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {object} []models.User
// @Router /user [get]
func (controller UserController) getUsers(c *gin.Context) {
	endpointUsers := []*models.User{}

	for _, userEntity := range controller.UserDao.FindAllUsers() {
		userModel := models.User {
			UserId: userEntity.UserId,
			Name: userEntity.Name,
			Email: userEntity.Email,
		}

		endpointUsers = append(endpointUsers, &userModel)
	}

	c.JSON(http.StatusOK, endpointUsers)
}