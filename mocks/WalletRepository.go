// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	entity "final-project-backend/entity"

	decimal "github.com/shopspring/decimal"

	mock "github.com/stretchr/testify/mock"
)

// WalletRepository is an autogenerated mock type for the WalletRepository type
type WalletRepository struct {
	mock.Mock
}

// CreateWallet provides a mock function with given fields: userId
func (_m *WalletRepository) CreateWallet(userId int) (*entity.Wallet, error) {
	ret := _m.Called(userId)

	var r0 *entity.Wallet
	if rf, ok := ret.Get(0).(func(int) *entity.Wallet); ok {
		r0 = rf(userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Wallet)
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

// DecreaseBalance provides a mock function with given fields: amount, wallet
func (_m *WalletRepository) DecreaseBalance(amount decimal.Decimal, wallet entity.Wallet) (*entity.Wallet, error) {
	ret := _m.Called(amount, wallet)

	var r0 *entity.Wallet
	if rf, ok := ret.Get(0).(func(decimal.Decimal, entity.Wallet) *entity.Wallet); ok {
		r0 = rf(amount, wallet)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Wallet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(decimal.Decimal, entity.Wallet) error); ok {
		r1 = rf(amount, wallet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetWalletByUserID provides a mock function with given fields: userId
func (_m *WalletRepository) GetWalletByUserID(userId int) (*entity.Wallet, error) {
	ret := _m.Called(userId)

	var r0 *entity.Wallet
	if rf, ok := ret.Get(0).(func(int) *entity.Wallet); ok {
		r0 = rf(userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Wallet)
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

// IncreaseBalance provides a mock function with given fields: amount, wallet
func (_m *WalletRepository) IncreaseBalance(amount decimal.Decimal, wallet entity.Wallet) (*entity.Wallet, error) {
	ret := _m.Called(amount, wallet)

	var r0 *entity.Wallet
	if rf, ok := ret.Get(0).(func(decimal.Decimal, entity.Wallet) *entity.Wallet); ok {
		r0 = rf(amount, wallet)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Wallet)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(decimal.Decimal, entity.Wallet) error); ok {
		r1 = rf(amount, wallet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsValidBalance provides a mock function with given fields: amount, wallet
func (_m *WalletRepository) IsValidBalance(amount decimal.Decimal, wallet entity.Wallet) bool {
	ret := _m.Called(amount, wallet)

	var r0 bool
	if rf, ok := ret.Get(0).(func(decimal.Decimal, entity.Wallet) bool); ok {
		r0 = rf(amount, wallet)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

type mockConstructorTestingTNewWalletRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewWalletRepository creates a new instance of WalletRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewWalletRepository(t mockConstructorTestingTNewWalletRepository) *WalletRepository {
	mock := &WalletRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
