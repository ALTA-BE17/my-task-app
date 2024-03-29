// Code generated by mockery v2.28.1. DO NOT EDIT.

package mocks

import (
	user "github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/features/user"
	mock "github.com/stretchr/testify/mock"
)

// UserData is an autogenerated mock type for the UserData type
type UserData struct {
	mock.Mock
}

// Deactive provides a mock function with given fields: userId
func (_m *UserData) Deactive(userId string) (user.Core, error) {
	ret := _m.Called(userId)

	var r0 user.Core
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (user.Core, error)); ok {
		return rf(userId)
	}
	if rf, ok := ret.Get(0).(func(string) user.Core); ok {
		r0 = rf(userId)
	} else {
		r0 = ret.Get(0).(user.Core)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Login provides a mock function with given fields: request
func (_m *UserData) Login(request user.Core) (user.Core, string, error) {
	ret := _m.Called(request)

	var r0 user.Core
	var r1 string
	var r2 error
	if rf, ok := ret.Get(0).(func(user.Core) (user.Core, string, error)); ok {
		return rf(request)
	}
	if rf, ok := ret.Get(0).(func(user.Core) user.Core); ok {
		r0 = rf(request)
	} else {
		r0 = ret.Get(0).(user.Core)
	}

	if rf, ok := ret.Get(1).(func(user.Core) string); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(user.Core) error); ok {
		r2 = rf(request)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Profile provides a mock function with given fields: userId
func (_m *UserData) Profile(userId string) (user.Core, error) {
	ret := _m.Called(userId)

	var r0 user.Core
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (user.Core, error)); ok {
		return rf(userId)
	}
	if rf, ok := ret.Get(0).(func(string) user.Core); ok {
		r0 = rf(userId)
	} else {
		r0 = ret.Get(0).(user.Core)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Register provides a mock function with given fields: request
func (_m *UserData) Register(request user.Core) (user.Core, error) {
	ret := _m.Called(request)

	var r0 user.Core
	var r1 error
	if rf, ok := ret.Get(0).(func(user.Core) (user.Core, error)); ok {
		return rf(request)
	}
	if rf, ok := ret.Get(0).(func(user.Core) user.Core); ok {
		r0 = rf(request)
	} else {
		r0 = ret.Get(0).(user.Core)
	}

	if rf, ok := ret.Get(1).(func(user.Core) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SearchUsers provides a mock function with given fields: userId, pattern
func (_m *UserData) SearchUsers(userId string, pattern string) ([]user.Core, error) {
	ret := _m.Called(userId, pattern)

	var r0 []user.Core
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) ([]user.Core, error)); ok {
		return rf(userId, pattern)
	}
	if rf, ok := ret.Get(0).(func(string, string) []user.Core); ok {
		r0 = rf(userId, pattern)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]user.Core)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(userId, pattern)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateProfile provides a mock function with given fields: userId, request
func (_m *UserData) UpdateProfile(userId string, request user.Core) (user.Core, error) {
	ret := _m.Called(userId, request)

	var r0 user.Core
	var r1 error
	if rf, ok := ret.Get(0).(func(string, user.Core) (user.Core, error)); ok {
		return rf(userId, request)
	}
	if rf, ok := ret.Get(0).(func(string, user.Core) user.Core); ok {
		r0 = rf(userId, request)
	} else {
		r0 = ret.Get(0).(user.Core)
	}

	if rf, ok := ret.Get(1).(func(string, user.Core) error); ok {
		r1 = rf(userId, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUserData interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserData creates a new instance of UserData. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserData(t mockConstructorTestingTNewUserData) *UserData {
	mock := &UserData{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
