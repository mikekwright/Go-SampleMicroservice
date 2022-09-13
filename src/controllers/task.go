package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"SampleGoService/src/models"
)

type TaskController struct {
}

func NewTaskController() *TaskController {
	return &TaskController{}
}

func (controller TaskController) RegisterController(router *gin.Engine) {
	router.GET("/task", controller.getAllTasks)
}

// @BasePath /task

// getAllTasks godoc
// @Summary Get a list of tasks
// @Schemes
// @Description list tasks
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {object} []models.Task
// @Router /task [get]
func (controller TaskController) getAllTasks(c *gin.Context) {
	c.JSON(http.StatusOK, []models.Task{
		{TaskId: 1, TaskName: "TaskController 1", Description: "Test Description", IsComplete: true},
		{TaskId: 2, TaskName: "TaskController 2", Description: "Test Description", IsComplete: false},
		{TaskId: 3, TaskName: "TaskController 3", Description: "Test Description", IsComplete: false},
		{TaskId: 4, TaskName: "TaskController 4", Description: "Test Description", IsComplete: true},
	})
}