// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	big "math/big"
	bindings "razor/pkg/bindings"

	bind "github.com/ethereum/go-ethereum/accounts/abi/bind"

	coretypes "razor/core/types"

	ethclient "github.com/ethereum/go-ethereum/ethclient"

	mock "github.com/stretchr/testify/mock"

	types "github.com/ethereum/go-ethereum/core/types"
)

// StakeManagerInterface is an autogenerated mock type for the StakeManagerInterface type
type StakeManagerInterface struct {
	mock.Mock
}

// ApproveUnstake provides a mock function with given fields: client, opts, staker, amount
func (_m *StakeManagerInterface) ApproveUnstake(client *ethclient.Client, opts *bind.TransactOpts, staker bindings.StructsStaker, amount *big.Int) (*types.Transaction, error) {
	ret := _m.Called(client, opts, staker, amount)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.TransactOpts, bindings.StructsStaker, *big.Int) *types.Transaction); ok {
		r0 = rf(client, opts, staker, amount)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, *bind.TransactOpts, bindings.StructsStaker, *big.Int) error); ok {
		r1 = rf(client, opts, staker, amount)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ClaimStakerReward provides a mock function with given fields: client, opts
func (_m *StakeManagerInterface) ClaimStakerReward(client *ethclient.Client, opts *bind.TransactOpts) (*types.Transaction, error) {
	ret := _m.Called(client, opts)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.TransactOpts) *types.Transaction); ok {
		r0 = rf(client, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, *bind.TransactOpts) error); ok {
		r1 = rf(client, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delegate provides a mock function with given fields: client, opts, stakerId, amount
func (_m *StakeManagerInterface) Delegate(client *ethclient.Client, opts *bind.TransactOpts, stakerId uint32, amount *big.Int) (*types.Transaction, error) {
	ret := _m.Called(client, opts, stakerId, amount)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.TransactOpts, uint32, *big.Int) *types.Transaction); ok {
		r0 = rf(client, opts, stakerId, amount)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, *bind.TransactOpts, uint32, *big.Int) error); ok {
		r1 = rf(client, opts, stakerId, amount)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBountyLock provides a mock function with given fields: client, opts, bountyId
func (_m *StakeManagerInterface) GetBountyLock(client *ethclient.Client, opts *bind.CallOpts, bountyId uint32) (coretypes.BountyLock, error) {
	ret := _m.Called(client, opts, bountyId)

	var r0 coretypes.BountyLock
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.CallOpts, uint32) coretypes.BountyLock); ok {
		r0 = rf(client, opts, bountyId)
	} else {
		r0 = ret.Get(0).(coretypes.BountyLock)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, *bind.CallOpts, uint32) error); ok {
		r1 = rf(client, opts, bountyId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMaturity provides a mock function with given fields: client, opts, age
func (_m *StakeManagerInterface) GetMaturity(client *ethclient.Client, opts *bind.CallOpts, age uint32) (uint16, error) {
	ret := _m.Called(client, opts, age)

	var r0 uint16
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.CallOpts, uint32) uint16); ok {
		r0 = rf(client, opts, age)
	} else {
		r0 = ret.Get(0).(uint16)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, *bind.CallOpts, uint32) error); ok {
		r1 = rf(client, opts, age)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InitiateWithdraw provides a mock function with given fields: client, opts, stakerId
func (_m *StakeManagerInterface) InitiateWithdraw(client *ethclient.Client, opts *bind.TransactOpts, stakerId uint32) (*types.Transaction, error) {
	ret := _m.Called(client, opts, stakerId)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.TransactOpts, uint32) *types.Transaction); ok {
		r0 = rf(client, opts, stakerId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, *bind.TransactOpts, uint32) error); ok {
		r1 = rf(client, opts, stakerId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RedeemBounty provides a mock function with given fields: client, opts, bountyId
func (_m *StakeManagerInterface) RedeemBounty(client *ethclient.Client, opts *bind.TransactOpts, bountyId uint32) (*types.Transaction, error) {
	ret := _m.Called(client, opts, bountyId)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.TransactOpts, uint32) *types.Transaction); ok {
		r0 = rf(client, opts, bountyId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, *bind.TransactOpts, uint32) error); ok {
		r1 = rf(client, opts, bountyId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResetUnstakeLock provides a mock function with given fields: client, opts, stakerId
func (_m *StakeManagerInterface) ResetUnstakeLock(client *ethclient.Client, opts *bind.TransactOpts, stakerId uint32) (*types.Transaction, error) {
	ret := _m.Called(client, opts, stakerId)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.TransactOpts, uint32) *types.Transaction); ok {
		r0 = rf(client, opts, stakerId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, *bind.TransactOpts, uint32) error); ok {
		r1 = rf(client, opts, stakerId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetDelegationAcceptance provides a mock function with given fields: client, opts, status
func (_m *StakeManagerInterface) SetDelegationAcceptance(client *ethclient.Client, opts *bind.TransactOpts, status bool) (*types.Transaction, error) {
	ret := _m.Called(client, opts, status)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.TransactOpts, bool) *types.Transaction); ok {
		r0 = rf(client, opts, status)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, *bind.TransactOpts, bool) error); ok {
		r1 = rf(client, opts, status)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Stake provides a mock function with given fields: client, txnOpts, epoch, amount
func (_m *StakeManagerInterface) Stake(client *ethclient.Client, txnOpts *bind.TransactOpts, epoch uint32, amount *big.Int) (*types.Transaction, error) {
	ret := _m.Called(client, txnOpts, epoch, amount)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.TransactOpts, uint32, *big.Int) *types.Transaction); ok {
		r0 = rf(client, txnOpts, epoch, amount)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, *bind.TransactOpts, uint32, *big.Int) error); ok {
		r1 = rf(client, txnOpts, epoch, amount)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StakerInfo provides a mock function with given fields: client, opts, stakerId
func (_m *StakeManagerInterface) StakerInfo(client *ethclient.Client, opts *bind.CallOpts, stakerId uint32) (coretypes.Staker, error) {
	ret := _m.Called(client, opts, stakerId)

	var r0 coretypes.Staker
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.CallOpts, uint32) coretypes.Staker); ok {
		r0 = rf(client, opts, stakerId)
	} else {
		r0 = ret.Get(0).(coretypes.Staker)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, *bind.CallOpts, uint32) error); ok {
		r1 = rf(client, opts, stakerId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UnlockWithdraw provides a mock function with given fields: client, opts, stakerId
func (_m *StakeManagerInterface) UnlockWithdraw(client *ethclient.Client, opts *bind.TransactOpts, stakerId uint32) (*types.Transaction, error) {
	ret := _m.Called(client, opts, stakerId)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.TransactOpts, uint32) *types.Transaction); ok {
		r0 = rf(client, opts, stakerId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, *bind.TransactOpts, uint32) error); ok {
		r1 = rf(client, opts, stakerId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Unstake provides a mock function with given fields: client, opts, stakerId, sAmount
func (_m *StakeManagerInterface) Unstake(client *ethclient.Client, opts *bind.TransactOpts, stakerId uint32, sAmount *big.Int) (*types.Transaction, error) {
	ret := _m.Called(client, opts, stakerId, sAmount)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.TransactOpts, uint32, *big.Int) *types.Transaction); ok {
		r0 = rf(client, opts, stakerId, sAmount)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, *bind.TransactOpts, uint32, *big.Int) error); ok {
		r1 = rf(client, opts, stakerId, sAmount)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateCommission provides a mock function with given fields: client, opts, commission
func (_m *StakeManagerInterface) UpdateCommission(client *ethclient.Client, opts *bind.TransactOpts, commission uint8) (*types.Transaction, error) {
	ret := _m.Called(client, opts, commission)

	var r0 *types.Transaction
	if rf, ok := ret.Get(0).(func(*ethclient.Client, *bind.TransactOpts, uint8) *types.Transaction); ok {
		r0 = rf(client, opts, commission)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, *bind.TransactOpts, uint8) error); ok {
		r1 = rf(client, opts, commission)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewStakeManagerInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewStakeManagerInterface creates a new instance of StakeManagerInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewStakeManagerInterface(t mockConstructorTestingTNewStakeManagerInterface) *StakeManagerInterface {
	mock := &StakeManagerInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
