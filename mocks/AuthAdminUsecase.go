// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	dto "final-project-backend/dto"

	mock "github.com/stretchr/testify/mock"
)

// AuthAdminUsecase is an autogenerated mock type for the AuthAdminUsecase type
type AuthAdminUsecase struct {
	mock.Mock
}

// Login provides a mock function with given fields: request
func (_m *AuthAdminUsecase) Login(request dto.LoginRequest) (*dto.LoginResponse, error) {
	ret := _m.Called(request)

	var r0 *dto.LoginResponse
	if rf, ok := ret.Get(0).(func(dto.LoginRequest) *dto.LoginResponse); ok {
		r0 = rf(request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.LoginResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(dto.LoginRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Register provides a mock function with given fields: request
func (_m *AuthAdminUsecase) Register(request dto.RegisterRequest) (*dto.RegisterResponse, error) {
	ret := _m.Called(request)

	var r0 *dto.RegisterResponse
	if rf, ok := ret.Get(0).(func(dto.RegisterRequest) *dto.RegisterResponse); ok {
		r0 = rf(request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.RegisterResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(dto.RegisterRequest) error); ok {
		r1 = rf(request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewAuthAdminUsecase interface {
	mock.TestingT
	Cleanup(func())
}

// NewAuthAdminUsecase creates a new instance of AuthAdminUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewAuthAdminUsecase(t mockConstructorTestingTNewAuthAdminUsecase) *AuthAdminUsecase {
	mock := &AuthAdminUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
