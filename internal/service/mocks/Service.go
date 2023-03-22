// Code generated by mockery v2.23.0. DO NOT EDIT.

package mocks

import (
	context "context"

	models "github.com/otaxhu/clean-code-project/internal/models"
	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// AddUserRole provides a mock function with given fields: ctx, userId, roleId
func (_m *Service) AddUserRole(ctx context.Context, userId string, roleId int) error {
	ret := _m.Called(ctx, userId, roleId)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int) error); ok {
		r0 = rf(ctx, userId, roleId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LoginUser provides a mock function with given fields: ctx, email, password
func (_m *Service) LoginUser(ctx context.Context, email string, password string) (*models.User, error) {
	ret := _m.Called(ctx, email, password)

	var r0 *models.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (*models.User, error)); ok {
		return rf(ctx, email, password)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *models.User); ok {
		r0 = rf(ctx, email, password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, email, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterUser provides a mock function with given fields: ctx, email, name, password
func (_m *Service) RegisterUser(ctx context.Context, email string, name string, password string) error {
	ret := _m.Called(ctx, email, name, password)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) error); ok {
		r0 = rf(ctx, email, name, password)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveUserRole provides a mock function with given fields: ctx, userId, roleId
func (_m *Service) RemoveUserRole(ctx context.Context, userId string, roleId int) error {
	ret := _m.Called(ctx, userId, roleId)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int) error); ok {
		r0 = rf(ctx, userId, roleId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewService interface {
	mock.TestingT
	Cleanup(func())
}

// NewService creates a new instance of Service. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewService(t mockConstructorTestingTNewService) *Service {
	mock := &Service{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
