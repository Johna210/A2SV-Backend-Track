package routers

import (
	"time"

	bootstrap "github.com/Johna210/A2SV-Backend-Track/Track7_clean_architecture/Bootstrap"
	repositories "github.com/Johna210/A2SV-Backend-Track/Track7_clean_architecture/Repositories"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewTaskAdminRoute(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repositories.NewUserRepository(db, "tasks")
}
