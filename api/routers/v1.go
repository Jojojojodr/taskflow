package routers

import (
	"github.com/Jojojojodr/taskflow/api/controllers"
	"github.com/Jojojojodr/taskflow/api/middleware"

	"github.com/gin-gonic/gin"
)

func V1Routes(rg *gin.RouterGroup) {
	v1 := rg.Group("/v1")
	v1.GET("/health", controllers.GetHealth)
	
	// Auth routes (public)
	auth := v1.Group("/auth")
	auth.POST("/login", controllers.LoginUser)
	auth.POST("/register", controllers.CreateUser)
	auth.POST("/logout", controllers.LogoutUser)
	auth.GET("/validate", controllers.ValidateUser)
	
	// Protected routes
	protected := v1.Group("/")
	protected.Use(middleware.AuthMiddleware())

	protected.GET("/tasks", controllers.GetTasks)

	task := protected.Group("/task")
	task.Use(middleware.UserMiddleware())
	task.POST("/", controllers.CreateTask)
	task.GET("/", controllers.GetTask)
	task.PUT("/", controllers.UpdateTask)
	task.DELETE("/", controllers.DeleteTask)

	user := protected.Group("/user")
	user.Use(middleware.UserMiddleware())
	user.GET("/", controllers.GetUser)
	user.PUT("/", controllers.UpdateUser)
	user.DELETE("/", controllers.DeleteUser)
}