// Code generated by mockery v2.16.0. DO NOT EDIT.

package mocks

import (
	user "sirloinapi/features/user"

	mock "github.com/stretchr/testify/mock"
)

// UserService is an autogenerated mock type for the UserService type
type UserService struct {
	mock.Mock
}

// Register provides a mock function with given fields: newUser
func (_m *UserService) Register(newUser user.Core) (user.Core, error) {
	ret := _m.Called(newUser)

	var r0 user.Core
	if rf, ok := ret.Get(0).(func(user.Core) user.Core); ok {
		r0 = rf(newUser)
	} else {
		r0 = ret.Get(0).(user.Core)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(user.Core) error); ok {
		r1 = rf(newUser)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUserService interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserService creates a new instance of UserService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserService(t mockConstructorTestingTNewUserService) *UserService {
	mock := &UserService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
