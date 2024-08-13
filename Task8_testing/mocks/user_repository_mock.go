package mocks

import (
	domain "github.com/Johna210/A2SV-Backend-Track/Task8_testing/Domain"
	"github.com/stretchr/testify/mock"
)

type UserRepositoryMock struct {
	mock.Mock
}

func (m *UserRepositoryMock) CreateUser(user *domain.User) error {
	ret := m.Called(user)

	var errorVal error
	if rf, ok := ret.Get(0).(func(*domain.User) error); ok {
		errorVal = rf(user)
	} else {
		errorVal = ret.Error(0)
	}

	return errorVal
}

func (m *UserRepositoryMock) Fetch() ([]domain.User, error) {
	ret := m.Called()

	var errorVal error
	if rf, ok := ret.Get(1).(func() error); ok {
		errorVal = rf()
	} else {
		errorVal = ret.Error(1)
	}

	var users []domain.User
	if rf, ok := ret.Get(0).(func() []domain.User); ok {
		users = rf()
	} else {
		users = ret.Get(0).([]domain.User)
	}

	return users, errorVal
}
