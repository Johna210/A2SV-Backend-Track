package routers

import (
	"time"

	bootstrap "github.com/Johna210/A2SV-Backend-Track/Track7_clean_architecture/Bootstrap"
	"github.com/Johna210/A2SV-Backend-Track/Track7_clean_architecture/Infrastructure/middleware"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db mongo.Database, gin *gin.Engine) {
	publicRouter := gin.Group("")

	// Public APIs
	NewLoginRoute(env, timeout, db, publicRouter)
	NewSignupRoute(env, timeout, db, publicRouter)

	// Protected APIs
	protectedRouter := gin.Group("")
	protectedRouter.Use(middleware.AuthMiddleware(env))
	NewTaskUserRoute(env, timeout, db, protectedRouter)

	// Protected and admin only APIs
	protectedAdminRouter := gin.Group("")
	protectedAdminRouter.Use(middleware.AuthMiddleware(env))
	protectedAdminRouter.Use(middleware.AdminMiddleware())
	NewPromoteRoute(env, timeout, db, protectedAdminRouter)
	NewTaskAdminRoute(env, timeout, db, protectedAdminRouter)
}
