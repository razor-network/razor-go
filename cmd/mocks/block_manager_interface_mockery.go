// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	big "math/big"

	bind "github.com/ethereum/go-ethereum/accounts/abi/bind"

	ethclient "github.com/ethereum/go-ethereum/ethclient"

	mock "github.com/stretchr/testify/mock"

	types "github.com/ethereum/go-ethereum/core/types"
)

// BlockManagerInterfaceMockery is an autogenerated mock type for the BlockManagerInterfaceMockery type
type BlockManagerInterfaceMockery struct {
	mock.Mock
}

// ClaimBlockReward provides a mock function with given fields: _a0, _a1
func (_m *BlockManagerInterfaceMockery) ClaimBlockReward(_a0 *ethclient.Client, _a1 *bind.TransactOpts) (*types.Transaction, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.TransactOpts) *types.Transaction); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, *bind.TransactOpts) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DisputeBiggestInfluenceProposed provides a mock function with given fields: _a0, _a1, _a2, _a3, _a4
func (_m *BlockManagerInterfaceMockery) DisputeBiggestInfluenceProposed(_a0 *ethclient.Client, _a1 *bind.TransactOpts, _a2 uint32, _a3 uint8, _a4 uint32) (*types.Transaction, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3, _a4)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.TransactOpts, uint32, uint8, uint32) *types.Transaction); ok {
		r0 = rf(_a0, _a1, _a2, _a3, _a4)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, *bind.TransactOpts, uint32, uint8, uint32) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3, _a4)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FinalizeDispute provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *BlockManagerInterfaceMockery) FinalizeDispute(_a0 *ethclient.Client, _a1 *bind.TransactOpts, _a2 uint32, _a3 uint8) (*types.Transaction, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.TransactOpts, uint32, uint8) *types.Transaction); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, *bind.TransactOpts, uint32, uint8) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Propose provides a mock function with given fields: _a0, _a1, _a2, _a3, _a4, _a5
func (_m *BlockManagerInterfaceMockery) Propose(_a0 *ethclient.Client, _a1 *bind.TransactOpts, _a2 uint32, _a3 []uint32, _a4 *big.Int, _a5 uint32) (*types.Transaction, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3, _a4, _a5)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.TransactOpts, uint32, []uint32, *big.Int, uint32) *types.Transaction); ok {
		r0 = rf(_a0, _a1, _a2, _a3, _a4, _a5)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, *bind.TransactOpts, uint32, []uint32, *big.Int, uint32) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3, _a4, _a5)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}