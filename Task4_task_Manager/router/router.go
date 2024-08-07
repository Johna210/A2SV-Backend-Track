package router

import (
	"github.com/Johna210/A2SV-Backend-Track/Track4_Task_Manager/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	router.GET("/tasks", controllers.GetTasksController)
	router.GET("/tasks/:id", controllers.GetTaskController)
	router.POST("/tasks", controllers.AddTaskController)
	router.PUT("/tasks/:id", controllers.UpdateTaskController)
	router.DELETE("/tasks/:id", controllers.DeleteTaskController)
}
