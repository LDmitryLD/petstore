// Code generated by mockery v2.35.4. DO NOT EDIT.

package mocks

import (
	models "projects/LDmitryLD/petstore/app/internal/models"

	mock "github.com/stretchr/testify/mock"
)

// StoreStorager is an autogenerated mock type for the StoreStorager type
type StoreStorager struct {
	mock.Mock
}

// CreateOrder provides a mock function with given fields: order
func (_m *StoreStorager) CreateOrder(order models.Order) models.Order {
	ret := _m.Called(order)

	var r0 models.Order
	if rf, ok := ret.Get(0).(func(models.Order) models.Order); ok {
		r0 = rf(order)
	} else {
		r0 = ret.Get(0).(models.Order)
	}

	return r0
}

// Delete provides a mock function with given fields: id
func (_m *StoreStorager) Delete(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetByID provides a mock function with given fields: id
func (_m *StoreStorager) GetByID(id int) (models.Order, error) {
	ret := _m.Called(id)

	var r0 models.Order
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (models.Order, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(int) models.Order); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Get(0).(models.Order)
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Inventory provides a mock function with given fields:
func (_m *StoreStorager) Inventory() map[string]int {
	ret := _m.Called()

	var r0 map[string]int
	if rf, ok := ret.Get(0).(func() map[string]int); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]int)
		}
	}

	return r0
}

// NewStoreStorager creates a new instance of StoreStorager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewStoreStorager(t interface {
	mock.TestingT
	Cleanup(func())
}) *StoreStorager {
	mock := &StoreStorager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
