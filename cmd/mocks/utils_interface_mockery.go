// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	big "math/big"

	bind "github.com/ethereum/go-ethereum/accounts/abi/bind"

	ethclient "github.com/ethereum/go-ethereum/ethclient"

	mock "github.com/stretchr/testify/mock"

	pflag "github.com/spf13/pflag"

	time "time"

	types "razor/core/types"
)

// UtilsInterfaceMockery is an autogenerated mock type for the UtilsInterfaceMockery type
type UtilsInterfaceMockery struct {
	mock.Mock
}

// AssignPassword provides a mock function with given fields: _a0
func (_m *UtilsInterfaceMockery) AssignPassword(_a0 *pflag.FlagSet) string {
	ret := _m.Called(_a0)

	var r0 string
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// CalculateBlockTime provides a mock function with given fields: _a0
func (_m *UtilsInterfaceMockery) CalculateBlockTime(_a0 *ethclient.Client) int64 {
	ret := _m.Called(_a0)

	var r0 int64
	if rf, ok := ret.Get(0).(func(*ethclient.Client) int64); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

// ConnectToClient provides a mock function with given fields: _a0
func (_m *UtilsInterfaceMockery) ConnectToClient(_a0 string) *ethclient.Client {
	ret := _m.Called(_a0)

	var r0 *ethclient.Client
	if rf, ok := ret.Get(0).(func(string) *ethclient.Client); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ethclient.Client)
		}
	}

	return r0
}

// GetActiveAssetsData provides a mock function with given fields: _a0, _a1
func (_m *UtilsInterfaceMockery) GetActiveAssetsData(_a0 *ethclient.Client, _a1 uint32) ([]*big.Int, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []*big.Int
	if rf, ok := ret.Get(0).(func(*ethclient.Client, uint32) []*big.Int); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, uint32) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetConfigFilePath provides a mock function with given fields:
func (_m *UtilsInterfaceMockery) GetConfigFilePath() (string, error) {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDelayedState provides a mock function with given fields: _a0, _a1
func (_m *UtilsInterfaceMockery) GetDelayedState(_a0 *ethclient.Client, _a1 int32) (int64, error) {
	ret := _m.Called(_a0, _a1)

	var r0 int64
	if rf, ok := ret.Get(0).(func(*ethclient.Client, int32) int64); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, int32) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEpoch provides a mock function with given fields: _a0
func (_m *UtilsInterfaceMockery) GetEpoch(_a0 *ethclient.Client) (uint32, error) {
	ret := _m.Called(_a0)

	var r0 uint32
	if rf, ok := ret.Get(0).(func(*ethclient.Client) uint32); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNumActiveAssets provides a mock function with given fields: _a0
func (_m *UtilsInterfaceMockery) GetNumActiveAssets(_a0 *ethclient.Client) (*big.Int, error) {
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

// GetOptions provides a mock function with given fields:
func (_m *UtilsInterfaceMockery) GetOptions() bind.CallOpts {
	ret := _m.Called()

	var r0 bind.CallOpts
	if rf, ok := ret.Get(0).(func() bind.CallOpts); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bind.CallOpts)
	}

	return r0
}

// GetRogueRandomValue provides a mock function with given fields: _a0
func (_m *UtilsInterfaceMockery) GetRogueRandomValue(_a0 int) *big.Int {
	ret := _m.Called(_a0)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(int) *big.Int); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	return r0
}

// GetStringAddress provides a mock function with given fields: _a0
func (_m *UtilsInterfaceMockery) GetStringAddress(_a0 *pflag.FlagSet) (string, error) {
	ret := _m.Called(_a0)

	var r0 string
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTxnOpts provides a mock function with given fields: _a0
func (_m *UtilsInterfaceMockery) GetTxnOpts(_a0 types.TransactionOptions) *bind.TransactOpts {
	ret := _m.Called(_a0)

	var r0 *bind.TransactOpts
	if rf, ok := ret.Get(0).(func(types.TransactionOptions) *bind.TransactOpts); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bind.TransactOpts)
		}
	}

	return r0
}

// GetUint32BountyId provides a mock function with given fields: _a0
func (_m *UtilsInterfaceMockery) GetUint32BountyId(_a0 *pflag.FlagSet) (uint32, error) {
	ret := _m.Called(_a0)

	var r0 uint32
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) uint32); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Sleep provides a mock function with given fields: _a0
func (_m *UtilsInterfaceMockery) Sleep(_a0 time.Duration) {
	_m.Called(_a0)
}

// ViperWriteConfigAs provides a mock function with given fields: _a0
func (_m *UtilsInterfaceMockery) ViperWriteConfigAs(_a0 string) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WaitForBlockCompletion provides a mock function with given fields: _a0, _a1
func (_m *UtilsInterfaceMockery) WaitForBlockCompletion(_a0 *ethclient.Client, _a1 string) int {
	ret := _m.Called(_a0, _a1)

	var r0 int
	if rf, ok := ret.Get(0).(func(*ethclient.Client, string) int); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}