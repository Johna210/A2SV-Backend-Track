package controllers

import (
	"net/http"
	"time"

	bootstrap "github.com/Johna210/A2SV-Backend-Track/Track7_clean_architecture/Bootstrap"
	domain "github.com/Johna210/A2SV-Backend-Track/Track7_clean_architecture/Domain"
	infrastructure "github.com/Johna210/A2SV-Backend-Track/Track7_clean_architecture/Infrastructure"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SignupController struct {
	SignupUsecase domain.SignupUsecase
	Env           *bootstrap.Env
}

func (sc *SignupController) Signup(c *gin.Context) {
	var request domain.SignupRequest

	err := c.ShouldBindJSON(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = sc.SignupUsecase.GetUserByEmail(c, request.Email)
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
		return
	}

	_, err = sc.SignupUsecase.GetUserByUsername(c, request.User_Name)
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
		return
	}

	encryptedPassword, err := infrastructure.HashPassword(request.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	request.Password = encryptedPassword
	userRole := "User"

	user := domain.User{
		ID:         primitive.NewObjectID(),
		First_Name: &request.First_Name,
		Last_Name:  &request.Last_Name,
		Email:      &request.Email,
		User_Name:  &request.User_Name,
		Password:   &request.Password,
		User_Role:  &userRole,
		Created_At: time.Now(),
		Updated_At: time.Now(),
	}

	err = sc.SignupUsecase.Create(c, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	accessToken, err := infrastructure.CreateAccessToken(&user, sc.Env.AccessTokenSecret, sc.Env.AccessTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	singupResponse := domain.SignupResponse{
		Message:      "User created successfully",
		Access_Token: accessToken,
		User_ID:      user.ID.Hex(),
	}

	c.JSON(http.StatusCreated, gin.H{"data": singupResponse})

}
