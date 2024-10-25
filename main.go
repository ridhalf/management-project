package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"management-project/app"
	"management-project/auth"
	"management-project/controller"
	"management-project/middleware"
	"management-project/repository"
	"management-project/service"
	"os"
)

func main() {
	app.Env()
	db := app.NewDB()

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	projectRepository := repository.NewProjectRepository(db)
	projectService := service.NewProjectService(projectRepository)

	authJwt := auth.NewJwt()
	authMiddleware := middleware.AuthMiddleware(authJwt, userService)

	userController := controller.NewUserController(userService, authJwt)
	projectController := controller.NewProjectController(projectService)

	router := gin.Default()
	//blocked by cors policy
	router.Use(cors.Default())
	//blocked by cors policy
	api := router.Group("/api/v1")
	api.POST("/users", userController.Register)
	api.POST("/users/login", userController.Login)
	api.GET("/users/:id", userController.FindById)

	api.GET("/projects", projectController.FindAll)
	api.GET("/projects/:id", projectController.FindById)
	api.POST("/projects", authMiddleware, projectController.Add)
	api.PUT("/projects", projectController.Update)
	api.DELETE("/projects/:id", projectController.Delete)

	err := router.Run(os.Getenv("DOMAIN"))
	if err != nil {
		panic(err)
	}
}
