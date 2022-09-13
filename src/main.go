package main

import (
	"os"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	docs "SampleGoService/docs"
	"SampleGoService/src/common"
	"SampleGoService/src/controllers"
)

func main() {
	router := gin.Default()
	docs.SwaggerInfo.BasePath = ""

	if strings.ToLower(os.Getenv("RUNNING_ENV")) != "prod" {
		log.Println("Not running in production, enabling CORS")
		router.Use(cors.Default())
	} else {
		log.Println("Running in production, setting gin to release mode")
		gin.SetMode(gin.ReleaseMode)
	}

	userController := CreateUserController() //  wire.Build(NewTaskController, NewUserController,) controllers.UserController{ UserDao: &db.UserDao{} }
	userController.RegisterController(router)

	taskController := controllers.TaskController{}
	taskController.RegisterController(router)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	ginSwagger.WrapHandler(swaggerfiles.Handler,
		ginSwagger.URL("http://localhost:8081/swagger/swagger.json"),
		ginSwagger.DefaultModelsExpandDepth(-1),
	)

	host := common.GetEnvOrDefault("HOST", "localhost")
	port := common.GetEnvOrDefault("PORT", "8080")

	err := router.Run(host + ":" + port)
	if err != nil {
		log.Fatal("Unable to start the router")
	}
}
