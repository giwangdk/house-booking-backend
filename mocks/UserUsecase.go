// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	dto "final-project-backend/dto"
	entity "final-project-backend/entity"

	mock "github.com/stretchr/testify/mock"
)

// UserUsecase is an autogenerated mock type for the UserUsecase type
type UserUsecase struct {
	mock.Mock
}

// ChangePassword provides a mock function with given fields: r, userId
func (_m *UserUsecase) ChangePassword(r dto.ChangePasswordRequest, userId int) (*dto.UpdateUserResponse, error) {
	ret := _m.Called(r, userId)

	var r0 *dto.UpdateUserResponse
	if rf, ok := ret.Get(0).(func(dto.ChangePasswordRequest, int) *dto.UpdateUserResponse); ok {
		r0 = rf(r, userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.UpdateUserResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(dto.ChangePasswordRequest, int) error); ok {
		r1 = rf(r, userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateUser provides a mock function with given fields: r
func (_m *UserUsecase) CreateUser(r entity.User) (*entity.User, error) {
	ret := _m.Called(r)

	var r0 *entity.User
	if rf, ok := ret.Get(0).(func(entity.User) *entity.User); ok {
		r0 = rf(r)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(entity.User) error); ok {
		r1 = rf(r)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUser provides a mock function with given fields: userID
func (_m *UserUsecase) GetUser(userID int) (*dto.UserDetail, error) {
	ret := _m.Called(userID)

	var r0 *dto.UserDetail
	if rf, ok := ret.Get(0).(func(int) *dto.UserDetail); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.UserDetail)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByEmail provides a mock function with given fields: email
func (_m *UserUsecase) GetUserByEmail(email string) (*entity.User, error) {
	ret := _m.Called(email)

	var r0 *entity.User
	if rf, ok := ret.Get(0).(func(string) *entity.User); ok {
		r0 = rf(email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsUserExist provides a mock function with given fields: email
func (_m *UserUsecase) IsUserExist(email string) (*entity.User, bool) {
	ret := _m.Called(email)

	var r0 *entity.User
	if rf, ok := ret.Get(0).(func(string) *entity.User); ok {
		r0 = rf(email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.User)
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(string) bool); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// IsUserExistByEmail provides a mock function with given fields: email
func (_m *UserUsecase) IsUserExistByEmail(email string) (*entity.User, bool) {
	ret := _m.Called(email)

	var r0 *entity.User
	if rf, ok := ret.Get(0).(func(string) *entity.User); ok {
		r0 = rf(email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.User)
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(string) bool); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// UpdateRole provides a mock function with given fields: email, role
func (_m *UserUsecase) UpdateRole(email string, role string) (*entity.User, error) {
	ret := _m.Called(email, role)

	var r0 *entity.User
	if rf, ok := ret.Get(0).(func(string, string) *entity.User); ok {
		r0 = rf(email, role)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(email, role)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUser provides a mock function with given fields: u, userId
func (_m *UserUsecase) UpdateUser(u dto.UpdateUserRequest, userId int) (*dto.UpdateUserResponse, error) {
	ret := _m.Called(u, userId)

	var r0 *dto.UpdateUserResponse
	if rf, ok := ret.Get(0).(func(dto.UpdateUserRequest, int) *dto.UpdateUserResponse); ok {
		r0 = rf(u, userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.UpdateUserResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(dto.UpdateUserRequest, int) error); ok {
		r1 = rf(u, userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUserUsecase interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserUsecase creates a new instance of UserUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserUsecase(t mockConstructorTestingTNewUserUsecase) *UserUsecase {
	mock := &UserUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
