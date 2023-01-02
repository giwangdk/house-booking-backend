// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	entity "final-project-backend/entity"

	mock "github.com/stretchr/testify/mock"

	time "time"
)

// HouseRepository is an autogenerated mock type for the HouseRepository type
type HouseRepository struct {
	mock.Mock
}

// CreateHouse provides a mock function with given fields: u
func (_m *HouseRepository) CreateHouse(u entity.HouseProfile) (*entity.HouseProfile, error) {
	ret := _m.Called(u)

	var r0 *entity.HouseProfile
	if rf, ok := ret.Get(0).(func(entity.HouseProfile) *entity.HouseProfile); ok {
		r0 = rf(u)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.HouseProfile)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(entity.HouseProfile) error); ok {
		r1 = rf(u)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteHouse provides a mock function with given fields: id
func (_m *HouseRepository) DeleteHouse(id int) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetHouseById provides a mock function with given fields: id
func (_m *HouseRepository) GetHouseById(id int) (*entity.House, error) {
	ret := _m.Called(id)

	var r0 *entity.House
	if rf, ok := ret.Get(0).(func(int) *entity.House); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.House)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetHouses provides a mock function with given fields: userId, page, limit, sortBy, sort, searchBy, filterByCity, checkIn, checkOut
func (_m *HouseRepository) GetHouses(userId int, page int, limit int, sortBy string, sort string, searchBy string, filterByCity int, checkIn string, checkOut string) (*[]entity.House, int, error) {
	ret := _m.Called(userId, page, limit, sortBy, sort, searchBy, filterByCity, checkIn, checkOut)

	var r0 *[]entity.House
	if rf, ok := ret.Get(0).(func(int, int, int, string, string, string, int, string, string) *[]entity.House); ok {
		r0 = rf(userId, page, limit, sortBy, sort, searchBy, filterByCity, checkIn, checkOut)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]entity.House)
		}
	}

	var r1 int
	if rf, ok := ret.Get(1).(func(int, int, int, string, string, string, int, string, string) int); ok {
		r1 = rf(userId, page, limit, sortBy, sort, searchBy, filterByCity, checkIn, checkOut)
	} else {
		r1 = ret.Get(1).(int)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(int, int, int, string, string, string, int, string, string) error); ok {
		r2 = rf(userId, page, limit, sortBy, sort, searchBy, filterByCity, checkIn, checkOut)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// IsBooked provides a mock function with given fields: id, deletedTime
func (_m *HouseRepository) IsBooked(id int, deletedTime time.Time) (bool, *entity.House) {
	ret := _m.Called(id, deletedTime)

	var r0 bool
	if rf, ok := ret.Get(0).(func(int, time.Time) bool); ok {
		r0 = rf(id, deletedTime)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 *entity.House
	if rf, ok := ret.Get(1).(func(int, time.Time) *entity.House); ok {
		r1 = rf(id, deletedTime)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*entity.House)
		}
	}

	return r0, r1
}

// UpdateHouse provides a mock function with given fields: u, userId
func (_m *HouseRepository) UpdateHouse(u entity.HouseProfile, userId int) (*entity.HouseProfile, error) {
	ret := _m.Called(u, userId)

	var r0 *entity.HouseProfile
	if rf, ok := ret.Get(0).(func(entity.HouseProfile, int) *entity.HouseProfile); ok {
		r0 = rf(u, userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.HouseProfile)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(entity.HouseProfile, int) error); ok {
		r1 = rf(u, userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewHouseRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewHouseRepository creates a new instance of HouseRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewHouseRepository(t mockConstructorTestingTNewHouseRepository) *HouseRepository {
	mock := &HouseRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
