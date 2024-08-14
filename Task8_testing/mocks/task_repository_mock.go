// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/Johna210/A2SV-Backend-Track/Task8_testing/Domain"
	mock "github.com/stretchr/testify/mock"
)

// TaskRepository is an autogenerated mock type for the TaskRepository type
type MockTaskRepository struct {
	mock.Mock
}

// CreateTask provides a mock function with given fields: c, task
func (_m *MockTaskRepository) CreateTask(c context.Context, task *domain.Task) error {
	ret := _m.Called(c, task)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Task) error); ok {
		r0 = rf(c, task)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteTask provides a mock function with given fields: c, id
func (_m *MockTaskRepository) DeleteTask(c context.Context, id string) error {
	ret := _m.Called(c, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(c, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Fetch provides a mock function with given fields: c
func (_m *MockTaskRepository) Fetch(c context.Context) ([]domain.Task, error) {
	ret := _m.Called(c)

	var r0 []domain.Task
	if rf, ok := ret.Get(0).(func(context.Context) []domain.Task); ok {
		r0 = rf(c)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Task)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(c)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByID provides a mock function with given fields: c, id
func (_m *MockTaskRepository) GetByID(c context.Context, id string) (domain.Task, error) {
	ret := _m.Called(c, id)

	var r0 domain.Task
	if rf, ok := ret.Get(0).(func(context.Context, string) domain.Task); ok {
		r0 = rf(c, id)
	} else {
		r0 = ret.Get(0).(domain.Task)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(c, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateTask provides a mock function with given fields: c, task, id
func (_m *MockTaskRepository) UpdateTask(c context.Context, task *domain.TaskUpdate, id string) (domain.Task, error) {
	ret := _m.Called(c, task, id)

	var r0 domain.Task
	if rf, ok := ret.Get(0).(func(context.Context, *domain.TaskUpdate, string) domain.Task); ok {
		r0 = rf(c, task, id)
	} else {
		r0 = ret.Get(0).(domain.Task)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *domain.TaskUpdate, string) error); ok {
		r1 = rf(c, task, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewTaskRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewTaskRepository creates a new instance of MockTaskRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTaskRepository(t mockConstructorTestingTNewTaskRepository) *MockTaskRepository {
	mock := &MockTaskRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}