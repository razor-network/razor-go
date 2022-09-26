// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"
	big "math/big"

	common "github.com/ethereum/go-ethereum/common"

	ethclient "github.com/ethereum/go-ethereum/ethclient"

	ethereum "github.com/ethereum/go-ethereum"

	mock "github.com/stretchr/testify/mock"

	types "github.com/ethereum/go-ethereum/core/types"
)

// ClientUtils is an autogenerated mock type for the ClientUtils type
type ClientUtils struct {
	mock.Mock
}

// BalanceAt provides a mock function with given fields: client, ctx, account, blockNumber
func (_m *ClientUtils) BalanceAt(client *ethclient.Client, ctx context.Context, account common.Address, blockNumber *big.Int) (*big.Int, error) {
	ret := _m.Called(client, ctx, account, blockNumber)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*ethclient.Client, context.Context, common.Address, *big.Int) *big.Int); ok {
		r0 = rf(client, ctx, account, blockNumber)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, context.Context, common.Address, *big.Int) error); ok {
		r1 = rf(client, ctx, account, blockNumber)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EstimateGas provides a mock function with given fields: client, ctx, msg
func (_m *ClientUtils) EstimateGas(client *ethclient.Client, ctx context.Context, msg ethereum.CallMsg) (uint64, error) {
	ret := _m.Called(client, ctx, msg)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(*ethclient.Client, context.Context, ethereum.CallMsg) uint64); ok {
		r0 = rf(client, ctx, msg)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, context.Context, ethereum.CallMsg) error); ok {
		r1 = rf(client, ctx, msg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FilterLogs provides a mock function with given fields: client, ctx, q
func (_m *ClientUtils) FilterLogs(client *ethclient.Client, ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error) {
	ret := _m.Called(client, ctx, q)

	var r0 []types.Log
	if rf, ok := ret.Get(0).(func(*ethclient.Client, context.Context, ethereum.FilterQuery) []types.Log); ok {
		r0 = rf(client, ctx, q)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.Log)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, context.Context, ethereum.FilterQuery) error); ok {
		r1 = rf(client, ctx, q)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HeaderByNumber provides a mock function with given fields: client, ctx, number
func (_m *ClientUtils) HeaderByNumber(client *ethclient.Client, ctx context.Context, number *big.Int) (*types.Header, error) {
	ret := _m.Called(client, ctx, number)

	var r0 *types.Header
	if rf, ok := ret.Get(0).(func(*ethclient.Client, context.Context, *big.Int) *types.Header); ok {
		r0 = rf(client, ctx, number)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Header)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, context.Context, *big.Int) error); ok {
		r1 = rf(client, ctx, number)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NonceAt provides a mock function with given fields: client, ctx, account
func (_m *ClientUtils) NonceAt(client *ethclient.Client, ctx context.Context, account common.Address) (uint64, error) {
	ret := _m.Called(client, ctx, account)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(*ethclient.Client, context.Context, common.Address) uint64); ok {
		r0 = rf(client, ctx, account)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, context.Context, common.Address) error); ok {
		r1 = rf(client, ctx, account)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SuggestGasPrice provides a mock function with given fields: client, ctx
func (_m *ClientUtils) SuggestGasPrice(client *ethclient.Client, ctx context.Context) (*big.Int, error) {
	ret := _m.Called(client, ctx)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*ethclient.Client, context.Context) *big.Int); ok {
		r0 = rf(client, ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, context.Context) error); ok {
		r1 = rf(client, ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TransactionReceipt provides a mock function with given fields: client, ctx, txHash
func (_m *ClientUtils) TransactionReceipt(client *ethclient.Client, ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	ret := _m.Called(client, ctx, txHash)

	var r0 *types.Receipt
	if rf, ok := ret.Get(0).(func(*ethclient.Client, context.Context, common.Hash) *types.Receipt); ok {
		r0 = rf(client, ctx, txHash)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Receipt)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, context.Context, common.Hash) error); ok {
		r1 = rf(client, ctx, txHash)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
