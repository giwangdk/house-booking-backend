// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	dto "final-project-backend/dto"
	entity "final-project-backend/entity"

	mock "github.com/stretchr/testify/mock"
)

// ReservationUsecase is an autogenerated mock type for the ReservationUsecase type
type ReservationUsecase struct {
	mock.Mock
}

// CreateReservation provides a mock function with given fields: r
func (_m *ReservationUsecase) CreateReservation(r entity.Reservation) (*entity.Reservation, error) {
	ret := _m.Called(r)

	var r0 *entity.Reservation
	if rf, ok := ret.Get(0).(func(entity.Reservation) *entity.Reservation); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Reservation)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(entity.Reservation) error); ok {
		r1 = rf(r)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateReservationWithUser provides a mock function with given fields: r
func (_m *ReservationUsecase) CreateReservationWithUser(r dto.CreateReservationRequest) (*dto.CreateReservationResponse, error) {
	ret := _m.Called(r)

	var r0 *dto.CreateReservationResponse
	if rf, ok := ret.Get(0).(func(dto.CreateReservationRequest) *dto.CreateReservationResponse); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.CreateReservationResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(dto.CreateReservationRequest) error); ok {
		r1 = rf(r)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetReservationByBookingCode provides a mock function with given fields: code
func (_m *ReservationUsecase) GetReservationByBookingCode(code string) (*dto.ReservationDetail, error) {
	ret := _m.Called(code)

	var r0 *dto.ReservationDetail
	if rf, ok := ret.Get(0).(func(string) *dto.ReservationDetail); ok {
		r0 = rf(code)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.ReservationDetail)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(code)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetReservationById provides a mock function with given fields: id
func (_m *ReservationUsecase) GetReservationById(id int) (*entity.Reservation, error) {
	ret := _m.Called(id)

	var r0 *entity.Reservation
	if rf, ok := ret.Get(0).(func(int) *entity.Reservation); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Reservation)
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

// GetReservationByUserId provides a mock function with given fields: userId
func (_m *ReservationUsecase) GetReservationByUserId(userId int) (*dto.ReservationList, error) {
	ret := _m.Called(userId)

	var r0 *dto.ReservationList
	if rf, ok := ret.Get(0).(func(int) *dto.ReservationList); ok {
		r0 = rf(userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.ReservationList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateStatusReservation provides a mock function with given fields: id, statusID
func (_m *ReservationUsecase) UpdateStatusReservation(id int, statusID int) (*entity.Reservation, error) {
	ret := _m.Called(id, statusID)

	var r0 *entity.Reservation
	if rf, ok := ret.Get(0).(func(int, int) *entity.Reservation); ok {
		r0 = rf(id, statusID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Reservation)
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

type mockConstructorTestingTNewReservationUsecase interface {
	mock.TestingT
	Cleanup(func())
}

// NewReservationUsecase creates a new instance of ReservationUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewReservationUsecase(t mockConstructorTestingTNewReservationUsecase) *ReservationUsecase {
	mock := &ReservationUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}