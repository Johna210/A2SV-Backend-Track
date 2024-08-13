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

func NewSignupRoute(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repositories.NewUserRepository(db, "users")
	sc := controllers.SignupController{
		SignupUsecase: usecases.NewSignupUsecase(ur, timeout),
		Env:           env,
	}
	group.POST("/user/register", sc.Signup)
}
