// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	dto "final-project-backend/dto"

	mock "github.com/stretchr/testify/mock"
)

// PickupUsecase is an autogenerated mock type for the PickupUsecase type
type PickupUsecase struct {
	mock.Mock
}

// CreatePickup provides a mock function with given fields: r
func (_m *PickupUsecase) CreatePickup(r dto.CreatePickupRequest) (*dto.CreatePickupResponse, error) {
	ret := _m.Called(r)

	var r0 *dto.CreatePickupResponse
	if rf, ok := ret.Get(0).(func(dto.CreatePickupRequest) *dto.CreatePickupResponse); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.CreatePickupResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(dto.CreatePickupRequest) error); ok {
		r1 = rf(r)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPickups provides a mock function with given fields: page, limit, sortBy, sort, searchBy, filterByStatus
func (_m *PickupUsecase) GetPickups(page int, limit int, sortBy string, sort string, searchBy string, filterByStatus int) (*dto.PickupLists, error) {
	ret := _m.Called(page, limit, sortBy, sort, searchBy, filterByStatus)

	var r0 *dto.PickupLists
	if rf, ok := ret.Get(0).(func(int, int, string, string, string, int) *dto.PickupLists); ok {
		r0 = rf(page, limit, sortBy, sort, searchBy, filterByStatus)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.PickupLists)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int, string, string, string, int) error); ok {
		r1 = rf(page, limit, sortBy, sort, searchBy, filterByStatus)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateStatusPickup provides a mock function with given fields: id, statusID
func (_m *PickupUsecase) UpdateStatusPickup(id int, statusID int) (*dto.CreatePickupResponse, error) {
	ret := _m.Called(id, statusID)

	var r0 *dto.CreatePickupResponse
	if rf, ok := ret.Get(0).(func(int, int) *dto.CreatePickupResponse); ok {
		r0 = rf(id, statusID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.CreatePickupResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, int) error); ok {
		r1 = rf(id, statusID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewPickupUsecase interface {
	mock.TestingT
	Cleanup(func())
}

// NewPickupUsecase creates a new instance of PickupUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPickupUsecase(t mockConstructorTestingTNewPickupUsecase) *PickupUsecase {
	mock := &PickupUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
