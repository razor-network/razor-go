// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	big "math/big"

	common "github.com/ethereum/go-ethereum/common"
	ethclient "github.com/ethereum/go-ethereum/ethclient"

	mock "github.com/stretchr/testify/mock"
)

// OptionUtils is an autogenerated mock type for the OptionUtils type
type OptionUtils struct {
	mock.Mock
}

// GetPendingNonceAtWithRetry provides a mock function with given fields: _a0, _a1
func (_m *OptionUtils) GetPendingNonceAtWithRetry(_a0 *ethclient.Client, _a1 common.Address) (uint64, error) {
	ret := _m.Called(_a0, _a1)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(*ethclient.Client, common.Address) uint64); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, common.Address) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MultiplyFloatAndBigInt provides a mock function with given fields: _a0, _a1
func (_m *OptionUtils) MultiplyFloatAndBigInt(_a0 *big.Int, _a1 float64) *big.Int {
	ret := _m.Called(_a0, _a1)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*big.Int, float64) *big.Int); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	return r0
}

// SuggestGasPriceWithRetry provides a mock function with given fields: _a0
func (_m *OptionUtils) SuggestGasPriceWithRetry(_a0 *ethclient.Client) (*big.Int, error) {
	ret := _m.Called(_a0)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*ethclient.Client) *big.Int); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}