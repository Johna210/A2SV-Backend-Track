package router

import (
	controller "github.com/Johna210/A2SV-Backend-Track/Track6_Task_Manager/controllers"
	middleware "github.com/Johna210/A2SV-Backend-Track/Track6_Task_Manager/middleware"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	router.POST("/user/register", controller.RegisterUser)
	router.POST("/user/login", controller.LoginUser)
	router.GET("/tasks", middleware.AuthMiddleware(), controller.GetTasksController)
	router.GET("/tasks/:id", middleware.AuthMiddleware(), controller.GetTaskController)
	router.POST("/tasks", middleware.AuthMiddleware(), middleware.AdminMiddleware(), controller.AddTaskController)
	router.PUT("/tasks/:id", middleware.AuthMiddleware(), middleware.AdminMiddleware(), controller.UpdateTaskController)
	router.DELETE("/tasks/:id", middleware.AuthMiddleware(), middleware.AdminMiddleware(), controller.DeleteTaskController)
	router.PUT("/user/promote/:id", middleware.AuthMiddleware(), middleware.AdminMiddleware(), controller.Promote)
}
