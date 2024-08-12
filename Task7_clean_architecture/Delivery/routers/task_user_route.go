package routers

import (
	"time"

	bootstrap "github.com/Johna210/A2SV-Backend-Track/Track7_clean_architecture/Bootstrap"
	"github.com/Johna210/A2SV-Backend-Track/Track7_clean_architecture/Delivery/controllers"
	repositories "github.com/Johna210/A2SV-Backend-Track/Track7_clean_architecture/Repositories"
	usecases "github.com/Johna210/A2SV-Backend-Track/Track7_clean_architecture/Usecases"
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
