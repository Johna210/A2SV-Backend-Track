package routers

import (
	"time"

	bootstrap "github.com/Johna210/A2SV-Backend-Track/Task8_testing/Bootstrap"
	"github.com/Johna210/A2SV-Backend-Track/Task8_testing/Delivery/controllers"
	repositories "github.com/Johna210/A2SV-Backend-Track/Task8_testing/Repositories"
	usecases "github.com/Johna210/A2SV-Backend-Track/Task8_testing/Usecases"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewTaskUserRoute(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	tr := repositories.NewTaskRepository(db, "tasks")
	tc := controllers.TaskController{
		TaskUsecase: usecases.NewTaskUsecase(tr, timeout),
	}

	group.GET("/tasks", tc.Fetch)
	group.GET("/tasks/:id", tc.GetTaskByID)
}
