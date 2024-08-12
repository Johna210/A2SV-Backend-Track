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

	user, err := pc.UserUsecase.Promote(c, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var response = make(map[string]interface{})
	response["message"] = "User promoted successfully"
	response["user"] = user

	c.JSON(http.StatusOK, response)
}
