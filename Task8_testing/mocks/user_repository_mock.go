package mocks

import (
	domain "github.com/Johna210/A2SV-Backend-Track/Task8_testing/Domain"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (mock *UserRepositoryMock) CreateUser(user *domain.User) error {
	args := mock.Called(user)
	_ = args.Get(0)
	return args.Error(1)
}

func (mock *UserRepositoryMock) Fetch() ([]domain.User, error) {
	args := mock.Called()
	return args.Get(0).([]domain.User), args.Error(1)
}
