// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	entity "final-project-backend/entity"

	mock "github.com/stretchr/testify/mock"
)

// HousePhotoRepository is an autogenerated mock type for the HousePhotoRepository type
type HousePhotoRepository struct {
	mock.Mock
}

// CreateHousePhoto provides a mock function with given fields: u
func (_m *HousePhotoRepository) CreateHousePhoto(u entity.HousePhoto) (*entity.HousePhoto, error) {
	ret := _m.Called(u)

	var r0 *entity.HousePhoto
	if rf, ok := ret.Get(0).(func(entity.HousePhoto) *entity.HousePhoto); ok {
		r0 = rf(u)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.HousePhoto)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(entity.HousePhoto) error); ok {
		r1 = rf(u)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteHousePhoto provides a mock function with given fields: id
func (_m *HousePhotoRepository) DeleteHousePhoto(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewHousePhotoRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewHousePhotoRepository creates a new instance of HousePhotoRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewHousePhotoRepository(t mockConstructorTestingTNewHousePhotoRepository) *HousePhotoRepository {
	mock := &HousePhotoRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
