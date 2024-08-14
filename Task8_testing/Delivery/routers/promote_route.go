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

func NewPromoteRoute(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repositories.NewUserRepository(db, "users")
	pc := controllers.PromoteController{
		UserUsecase: usecases.NewPromoteUsecase(ur, timeout),
		Env:         env,
	}
	group.PUT("/user/promote/:id", pc.Promote)
}