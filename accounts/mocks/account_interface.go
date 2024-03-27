// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	ecdsa "crypto/ecdsa"

	accounts "github.com/ethereum/go-ethereum/accounts"

	mock "github.com/stretchr/testify/mock"

	types "razor/core/types"
)

// AccountInterface is an autogenerated mock type for the AccountInterface type
type AccountInterface struct {
	mock.Mock
}

// Accounts provides a mock function with given fields: path
func (_m *AccountInterface) Accounts(path string) []accounts.Account {
	ret := _m.Called(path)

	var r0 []accounts.Account
	if rf, ok := ret.Get(0).(func(string) []accounts.Account); ok {
		r0 = rf(path)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]accounts.Account)
		}
	}

	return r0
}

// CreateAccount provides a mock function with given fields: path, password
func (_m *AccountInterface) CreateAccount(path string, password string) accounts.Account {
	ret := _m.Called(path, password)

	var r0 accounts.Account
	if rf, ok := ret.Get(0).(func(string, string) accounts.Account); ok {
		r0 = rf(path, password)
	} else {
		r0 = ret.Get(0).(accounts.Account)
	}

	return r0
}

// GetPrivateKey provides a mock function with given fields: address, password, keystorePath
func (_m *AccountInterface) GetPrivateKey(address string, password string, keystorePath string) (*ecdsa.PrivateKey, error) {
	ret := _m.Called(address, password, keystorePath)

	var r0 *ecdsa.PrivateKey
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, string) (*ecdsa.PrivateKey, error)); ok {
		return rf(address, password, keystorePath)
	}
	if rf, ok := ret.Get(0).(func(string, string, string) *ecdsa.PrivateKey); ok {
		r0 = rf(address, password, keystorePath)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ecdsa.PrivateKey)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string, string) error); ok {
		r1 = rf(address, password, keystorePath)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewAccount provides a mock function with given fields: path, passphrase
func (_m *AccountInterface) NewAccount(path string, passphrase string) (accounts.Account, error) {
	ret := _m.Called(path, passphrase)

	var r0 accounts.Account
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (accounts.Account, error)); ok {
		return rf(path, passphrase)
	}
	if rf, ok := ret.Get(0).(func(string, string) accounts.Account); ok {
		r0 = rf(path, passphrase)
	} else {
		r0 = ret.Get(0).(accounts.Account)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(path, passphrase)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Sign provides a mock function with given fields: digestHash, prv
func (_m *AccountInterface) Sign(digestHash []byte, prv *ecdsa.PrivateKey) ([]byte, error) {
	ret := _m.Called(digestHash, prv)

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func([]byte, *ecdsa.PrivateKey) ([]byte, error)); ok {
		return rf(digestHash, prv)
	}
	if rf, ok := ret.Get(0).(func([]byte, *ecdsa.PrivateKey) []byte); ok {
		r0 = rf(digestHash, prv)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func([]byte, *ecdsa.PrivateKey) error); ok {
		r1 = rf(digestHash, prv)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SignData provides a mock function with given fields: hash, account, defaultPath
func (_m *AccountInterface) SignData(hash []byte, account types.Account, defaultPath string) ([]byte, error) {
	ret := _m.Called(hash, account, defaultPath)

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func([]byte, types.Account, string) ([]byte, error)); ok {
		return rf(hash, account, defaultPath)
	}
	if rf, ok := ret.Get(0).(func([]byte, types.Account, string) []byte); ok {
		r0 = rf(hash, account, defaultPath)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func([]byte, types.Account, string) error); ok {
		r1 = rf(hash, account, defaultPath)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewAccountInterface creates a new instance of AccountInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewAccountInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *AccountInterface {
	mock := &AccountInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
