package controllers

import (
	"net/http"

	bootstrap "github.com/Johna210/A2SV-Backend-Track/Track7_clean_architecture/Bootstrap"
	domain "github.com/Johna210/A2SV-Backend-Track/Track7_clean_architecture/Domain"
	"github.com/gin-gonic/gin"
)

type PromoteController struct {
	UserUsecase domain.UserUsecase
	Env         *bootstrap.Env
}

func (pc *PromoteController) Promote(c *gin.Context) {
	id := c.Param("id")

	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	user, err := pc..GetUserByUsername(c, request.UserName)
	if err != nil {
		c.JSON(400, gin.H{"error": "User not found"})
		return
	}

	userRole := "Admin"
	user.User_Role = &userRole

	err = pc.UserUsecase.UpdateUser(c, &user)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "User promoted successfully"})
}
