// Code generated by mockery v2.35.4. DO NOT EDIT.

package mocks

import (
	multipart "mime/multipart"
	models "projects/LDmitryLD/petstore/app/internal/models"

	mock "github.com/stretchr/testify/mock"
)

// PetStorager is an autogenerated mock type for the PetStorager type
type PetStorager struct {
	mock.Mock
}

// Create provides a mock function with given fields: pet
func (_m *PetStorager) Create(pet models.Pet) models.Pet {
	ret := _m.Called(pet)

	var r0 models.Pet
	if rf, ok := ret.Get(0).(func(models.Pet) models.Pet); ok {
		r0 = rf(pet)
	} else {
		r0 = ret.Get(0).(models.Pet)
	}

	return r0
}

// Delete provides a mock function with given fields: petID
func (_m *PetStorager) Delete(petID int) error {
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
func (_m *PetStorager) FindByStatus(values []string) []models.Pet {
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

// GetByID provides a mock function with given fields: petID
func (_m *PetStorager) GetByID(petID int) (models.Pet, error) {
	ret := _m.Called(petID)

	var r0 models.Pet
	var r1 error
	if rf, ok := ret.Get(0).(func(int) (models.Pet, error)); ok {
		return rf(petID)
	}
	if rf, ok := ret.Get(0).(func(int) models.Pet); ok {
		r0 = rf(petID)
	} else {
		r0 = ret.Get(0).(models.Pet)
	}

	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(petID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetList provides a mock function with given fields:
func (_m *PetStorager) GetList() []models.Pet {
	ret := _m.Called()

	var r0 []models.Pet
	if rf, ok := ret.Get(0).(func() []models.Pet); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Pet)
		}
	}

	return r0
}

// Update provides a mock function with given fields: pet
func (_m *PetStorager) Update(pet models.Pet) error {
	ret := _m.Called(pet)

	var r0 error
	if rf, ok := ret.Get(0).(func(models.Pet) error); ok {
		r0 = rf(pet)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateByID provides a mock function with given fields: petID, name, status
func (_m *PetStorager) UpdateByID(petID int, name string, status string) error {
	ret := _m.Called(petID, name, status)

	var r0 error
	if rf, ok := ret.Get(0).(func(int, string, string) error); ok {
		r0 = rf(petID, name, status)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UploadImage provides a mock function with given fields: file, fileName, id
func (_m *PetStorager) UploadImage(file multipart.File, fileName string, id int) error {
	ret := _m.Called(file, fileName, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(multipart.File, string, int) error); ok {
		r0 = rf(file, fileName, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewPetStorager creates a new instance of PetStorager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPetStorager(t interface {
	mock.TestingT
	Cleanup(func())
}) *PetStorager {
	mock := &PetStorager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}