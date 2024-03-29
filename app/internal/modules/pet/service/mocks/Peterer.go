// Code generated by mockery v2.35.4. DO NOT EDIT.

package mocks

import (
	models "projects/LDmitryLD/petstore/app/internal/models"
	service "projects/LDmitryLD/petstore/app/internal/modules/pet/service"

	mock "github.com/stretchr/testify/mock"
)

// Peterer is an autogenerated mock type for the Peterer type
type Peterer struct {
	mock.Mock
}

// Create provides a mock function with given fields: in
func (_m *Peterer) Create(in service.CreateIn) service.CreateOut {
	ret := _m.Called(in)

	var r0 service.CreateOut
	if rf, ok := ret.Get(0).(func(service.CreateIn) service.CreateOut); ok {
		r0 = rf(in)
	} else {
		r0 = ret.Get(0).(service.CreateOut)
	}

	return r0
}

// Delete provides a mock function with given fields: petID
func (_m *Peterer) Delete(petID int) error {
	ret := _m.Called(petID)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(petID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindByStatus provides a mock function with given fields: values
func (_m *Peterer) FindByStatus(values []string) []models.Pet {
	ret := _m.Called(values)

	var r0 []models.Pet
	if rf, ok := ret.Get(0).(func([]string) []models.Pet); ok {
		r0 = rf(values)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Pet)
		}
	}

	return r0
}

// GetByID provides a mock function with given fields: in
func (_m *Peterer) GetByID(in service.GetByIDIn) service.GetByIDOut {
	ret := _m.Called(in)

	var r0 service.GetByIDOut
	if rf, ok := ret.Get(0).(func(service.GetByIDIn) service.GetByIDOut); ok {
		r0 = rf(in)
	} else {
		r0 = ret.Get(0).(service.GetByIDOut)
	}

	return r0
}

// GetList provides a mock function with given fields:
func (_m *Peterer) GetList() service.GetListOut {
	ret := _m.Called()

	var r0 service.GetListOut
	if rf, ok := ret.Get(0).(func() service.GetListOut); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(service.GetListOut)
	}

	return r0
}

// Update provides a mock function with given fields: in
func (_m *Peterer) Update(in service.UpdateIn) service.UpdateOut {
	ret := _m.Called(in)

	var r0 service.UpdateOut
	if rf, ok := ret.Get(0).(func(service.UpdateIn) service.UpdateOut); ok {
		r0 = rf(in)
	} else {
		r0 = ret.Get(0).(service.UpdateOut)
	}

	return r0
}

// UpdateByID provides a mock function with given fields: in
func (_m *Peterer) UpdateByID(in service.UpdateByIDIn) error {
	ret := _m.Called(in)

	var r0 error
	if rf, ok := ret.Get(0).(func(service.UpdateByIDIn) error); ok {
		r0 = rf(in)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UploadImage provides a mock function with given fields: in
func (_m *Peterer) UploadImage(in service.UploadIn) service.UploadOut {
	ret := _m.Called(in)

	var r0 service.UploadOut
	if rf, ok := ret.Get(0).(func(service.UploadIn) service.UploadOut); ok {
		r0 = rf(in)
	} else {
		r0 = ret.Get(0).(service.UploadOut)
	}

	return r0
}

// NewPeterer creates a new instance of Peterer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPeterer(t interface {
	mock.TestingT
	Cleanup(func())
}) *Peterer {
	mock := &Peterer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
