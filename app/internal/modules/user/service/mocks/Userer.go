// Code generated by mockery v2.35.4. DO NOT EDIT.

package mocks

import (
	models "projects/LDmitryLD/petstore/app/internal/models"

	mock "github.com/stretchr/testify/mock"
)

// Userer is an autogenerated mock type for the Userer type
type Userer struct {
	mock.Mock
}

// Create provides a mock function with given fields: user
func (_m *Userer) Create(user models.User) int {
	ret := _m.Called(user)

	var r0 int
	if rf, ok := ret.Get(0).(func(models.User) int); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// Delete provides a mock function with given fields: username
func (_m *Userer) Delete(username string) error {
	ret := _m.Called(username)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(username)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetByUsername provides a mock function with given fields: username
func (_m *Userer) GetByUsername(username string) (models.User, error) {
	ret := _m.Called(username)

	var r0 models.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (models.User, error)); ok {
		return rf(username)
	}
	if rf, ok := ret.Get(0).(func(string) models.User); ok {
		r0 = rf(username)
	} else {
		r0 = ret.Get(0).(models.User)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: username, user
func (_m *Userer) Update(username string, user models.User) error {
	ret := _m.Called(username, user)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, models.User) error); ok {
		r0 = rf(username, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewUserer creates a new instance of Userer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserer(t interface {
	mock.TestingT
	Cleanup(func())
}) *Userer {
	mock := &Userer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}