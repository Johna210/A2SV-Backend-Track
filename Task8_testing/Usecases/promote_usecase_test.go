package usecases

import (
	"context"
	"testing"
	"time"

	domain "github.com/Johna210/A2SV-Backend-Track/Task8_testing/Domain"
	"github.com/Johna210/A2SV-Backend-Track/Task8_testing/mocks"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var testUserId = primitive.NewObjectID()

type UserPromoteUsecase struct {
	suite.Suite
	mockUserRepo   *mocks.MockUserRepository
	promoteUsecase domain.UserUsecase
	contextTimeout time.Duration
	testUser       *domain.User
}

func (suite *UserPromoteUsecase) SetupSuite() {
	suite.mockUserRepo = new(mocks.MockUserRepository)
	suite.contextTimeout = time.Second * 2
	suite.promoteUsecase = NewPromoteUsecase(suite.mockUserRepo, suite.contextTimeout)
	suite.testUser = &domain.User{
		ID:         testUserId,
		First_Name: &firstName,
		Last_Name:  &lastName,
		User_Name:  &userName,
		User_Role:  &userRole,
		Password:   &password,
		Email:      &email,
		Created_At: time.Now(),
		Updated_At: time.Now(),
	}
}

// TestPromoteUser tests the PromoteUser method in the use case
func (suite *UserPromoteUsecase) TestPromoteUser() {
	suite.mockUserRepo.On("Promote", suite.testUser.ID.Hex()).Return(*suite.testUser, nil)
	_, err := suite.promoteUsecase.Promote(context.Background(), suite.testUser.ID.Hex())
	suite.NoError(err)
}

func TestUserPromoteUsecase(t *testing.T) {
	suite.Run(t, new(UserSignupUsecase))
}
