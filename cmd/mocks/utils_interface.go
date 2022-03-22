// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	big "math/big"
	bindings "razor/pkg/bindings"

	bind "github.com/ethereum/go-ethereum/accounts/abi/bind"

	common "github.com/ethereum/go-ethereum/common"

	ethclient "github.com/ethereum/go-ethereum/ethclient"

	mock "github.com/stretchr/testify/mock"

	pflag "github.com/spf13/pflag"

	types "razor/core/types"
)

// UtilsInterface is an autogenerated mock type for the UtilsInterface type
type UtilsInterface struct {
	mock.Mock
}

// AddJobToJSON provides a mock function with given fields: _a0, _a1
func (_m *UtilsInterface) AddJobToJSON(_a0 string, _a1 *types.StructsJob) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, *types.StructsJob) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AllZero provides a mock function with given fields: _a0
func (_m *UtilsInterface) AllZero(_a0 [32]byte) bool {
	ret := _m.Called(_a0)

	var r0 bool
	if rf, ok := ret.Get(0).(func([32]byte) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// AssignPassword provides a mock function with given fields: _a0
func (_m *UtilsInterface) AssignPassword(_a0 *pflag.FlagSet) string {
	ret := _m.Called(_a0)

	var r0 string
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// AssignStakerId provides a mock function with given fields: _a0, _a1, _a2
func (_m *UtilsInterface) AssignStakerId(_a0 *pflag.FlagSet, _a1 *ethclient.Client, _a2 string) (uint32, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 uint32
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet, *ethclient.Client, string) uint32); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*pflag.FlagSet, *ethclient.Client, string) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CalculateBlockTime provides a mock function with given fields: _a0
func (_m *UtilsInterface) CalculateBlockTime(_a0 *ethclient.Client) int64 {
	ret := _m.Called(_a0)

	var r0 int64
	if rf, ok := ret.Get(0).(func(*ethclient.Client) int64); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

// CheckAmountAndBalance provides a mock function with given fields: _a0, _a1
func (_m *UtilsInterface) CheckAmountAndBalance(_a0 *big.Int, _a1 *big.Int) *big.Int {
	ret := _m.Called(_a0, _a1)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*big.Int, *big.Int) *big.Int); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	return r0
}

// CheckEthBalanceIsZero provides a mock function with given fields: _a0, _a1
func (_m *UtilsInterface) CheckEthBalanceIsZero(_a0 *ethclient.Client, _a1 string) {
	_m.Called(_a0, _a1)
}

// ConnectToClient provides a mock function with given fields: _a0
func (_m *UtilsInterface) ConnectToClient(_a0 string) *ethclient.Client {
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

// ConvertBigIntArrayToUint32Array provides a mock function with given fields: _a0
func (_m *UtilsInterface) ConvertBigIntArrayToUint32Array(_a0 []*big.Int) []uint32 {
	ret := _m.Called(_a0)

	var r0 []uint32
	if rf, ok := ret.Get(0).(func([]*big.Int) []uint32); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]uint32)
		}
	}

	return r0
}

// ConvertRZRToSRZR provides a mock function with given fields: _a0, _a1, _a2
func (_m *UtilsInterface) ConvertRZRToSRZR(_a0 *big.Int, _a1 *big.Int, _a2 *big.Int) (*big.Int, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*big.Int, *big.Int, *big.Int) *big.Int); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*big.Int, *big.Int, *big.Int) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ConvertSRZRToRZR provides a mock function with given fields: _a0, _a1, _a2
func (_m *UtilsInterface) ConvertSRZRToRZR(_a0 *big.Int, _a1 *big.Int, _a2 *big.Int) *big.Int {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*big.Int, *big.Int, *big.Int) *big.Int); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	return r0
}

// ConvertUint32ArrayToBigIntArray provides a mock function with given fields: _a0
func (_m *UtilsInterface) ConvertUint32ArrayToBigIntArray(_a0 []uint32) []*big.Int {
	ret := _m.Called(_a0)

	var r0 []*big.Int
	if rf, ok := ret.Get(0).(func([]uint32) []*big.Int); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*big.Int)
		}
	}

	return r0
}

// ConvertUintArrayToUint16Array provides a mock function with given fields: _a0
func (_m *UtilsInterface) ConvertUintArrayToUint16Array(_a0 []uint) []uint16 {
	ret := _m.Called(_a0)

	var r0 []uint16
	if rf, ok := ret.Get(0).(func([]uint) []uint16); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]uint16)
		}
	}

	return r0
}

// ConvertWeiToEth provides a mock function with given fields: _a0
func (_m *UtilsInterface) ConvertWeiToEth(_a0 *big.Int) (*big.Float, error) {
	ret := _m.Called(_a0)

	var r0 *big.Float
	if rf, ok := ret.Get(0).(func(*big.Int) *big.Float); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Float)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*big.Int) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteJobFromJSON provides a mock function with given fields: _a0, _a1
func (_m *UtilsInterface) DeleteJobFromJSON(_a0 string, _a1 string) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FetchBalance provides a mock function with given fields: _a0, _a1
func (_m *UtilsInterface) FetchBalance(_a0 *ethclient.Client, _a1 string) (*big.Int, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*ethclient.Client, string) *big.Int); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetActiveCollectionIds provides a mock function with given fields: _a0
func (_m *UtilsInterface) GetActiveCollectionIds(_a0 *ethclient.Client) ([]uint16, error) {
	ret := _m.Called(_a0)

	var r0 []uint16
	if rf, ok := ret.Get(0).(func(*ethclient.Client) []uint16); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]uint16)
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

// GetAggregatedDataOfCollection provides a mock function with given fields: client, collectionId, epoch
func (_m *UtilsInterface) GetAggregatedDataOfCollection(client *ethclient.Client, collectionId uint16, epoch uint32) (*big.Int, error) {
	ret := _m.Called(client, collectionId, epoch)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*ethclient.Client, uint16, uint32) *big.Int); ok {
		r0 = rf(client, collectionId, epoch)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, uint16, uint32) error); ok {
		r1 = rf(client, collectionId, epoch)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAmountInDecimal provides a mock function with given fields: _a0
func (_m *UtilsInterface) GetAmountInDecimal(_a0 *big.Int) *big.Float {
	ret := _m.Called(_a0)

	var r0 *big.Float
	if rf, ok := ret.Get(0).(func(*big.Int) *big.Float); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Float)
		}
	}

	return r0
}

// GetAmountInWei provides a mock function with given fields: _a0
func (_m *UtilsInterface) GetAmountInWei(_a0 *big.Int) *big.Int {
	ret := _m.Called(_a0)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*big.Int) *big.Int); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	return r0
}

// GetBlockManager provides a mock function with given fields: _a0
func (_m *UtilsInterface) GetBlockManager(_a0 *ethclient.Client) *bindings.BlockManager {
	ret := _m.Called(_a0)

	var r0 *bindings.BlockManager
	if rf, ok := ret.Get(0).(func(*ethclient.Client) *bindings.BlockManager); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bindings.BlockManager)
		}
	}

	return r0
}

// GetCollections provides a mock function with given fields: _a0
func (_m *UtilsInterface) GetCollections(_a0 *ethclient.Client) ([]bindings.StructsCollection, error) {
	ret := _m.Called(_a0)

	var r0 []bindings.StructsCollection
	if rf, ok := ret.Get(0).(func(*ethclient.Client) []bindings.StructsCollection); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]bindings.StructsCollection)
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

// GetCommitments provides a mock function with given fields: _a0, _a1
func (_m *UtilsInterface) GetCommitments(_a0 *ethclient.Client, _a1 string) ([32]byte, error) {
	ret := _m.Called(_a0, _a1)

	var r0 [32]byte
	if rf, ok := ret.Get(0).(func(*ethclient.Client, string) [32]byte); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([32]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetConfigFilePath provides a mock function with given fields:
func (_m *UtilsInterface) GetConfigFilePath() (string, error) {
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

// GetDefaultPath provides a mock function with given fields:
func (_m *UtilsInterface) GetDefaultPath() (string, error) {
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
func (_m *UtilsInterface) GetDelayedState(_a0 *ethclient.Client, _a1 int32) (int64, error) {
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
func (_m *UtilsInterface) GetEpoch(_a0 *ethclient.Client) (uint32, error) {
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

// GetEpochLastCommitted provides a mock function with given fields: _a0, _a1
func (_m *UtilsInterface) GetEpochLastCommitted(_a0 *ethclient.Client, _a1 uint32) (uint32, error) {
	ret := _m.Called(_a0, _a1)

	var r0 uint32
	if rf, ok := ret.Get(0).(func(*ethclient.Client, uint32) uint32); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, uint32) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEpochLastRevealed provides a mock function with given fields: _a0, _a1
func (_m *UtilsInterface) GetEpochLastRevealed(_a0 *ethclient.Client, _a1 uint32) (uint32, error) {
	ret := _m.Called(_a0, _a1)

	var r0 uint32
	if rf, ok := ret.Get(0).(func(*ethclient.Client, uint32) uint32); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, uint32) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEpochLimitForUpdateCommission provides a mock function with given fields: _a0
func (_m *UtilsInterface) GetEpochLimitForUpdateCommission(_a0 *ethclient.Client) (uint16, error) {
	ret := _m.Called(_a0)

	var r0 uint16
	if rf, ok := ret.Get(0).(func(*ethclient.Client) uint16); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(uint16)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetFractionalAmountInWei provides a mock function with given fields: _a0, _a1
func (_m *UtilsInterface) GetFractionalAmountInWei(_a0 *big.Int, _a1 string) (*big.Int, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*big.Int, string) *big.Int); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*big.Int, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetInfluenceSnapshot provides a mock function with given fields: _a0, _a1, _a2
func (_m *UtilsInterface) GetInfluenceSnapshot(_a0 *ethclient.Client, _a1 uint32, _a2 uint32) (*big.Int, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*ethclient.Client, uint32, uint32) *big.Int); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, uint32, uint32) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetJobFilePath provides a mock function with given fields:
func (_m *UtilsInterface) GetJobFilePath() (string, error) {
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

// GetJobs provides a mock function with given fields: _a0
func (_m *UtilsInterface) GetJobs(_a0 *ethclient.Client) ([]bindings.StructsJob, error) {
	ret := _m.Called(_a0)

	var r0 []bindings.StructsJob
	if rf, ok := ret.Get(0).(func(*ethclient.Client) []bindings.StructsJob); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]bindings.StructsJob)
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

// GetLock provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *UtilsInterface) GetLock(_a0 *ethclient.Client, _a1 string, _a2 uint32, _a3 uint8) (types.Locks, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 types.Locks
	if rf, ok := ret.Get(0).(func(*ethclient.Client, string, uint32, uint8) types.Locks); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		r0 = ret.Get(0).(types.Locks)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, string, uint32, uint8) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMaxAltBlocks provides a mock function with given fields: _a0
func (_m *UtilsInterface) GetMaxAltBlocks(_a0 *ethclient.Client) (uint8, error) {
	ret := _m.Called(_a0)

	var r0 uint8
	if rf, ok := ret.Get(0).(func(*ethclient.Client) uint8); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(uint8)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMaxCommission provides a mock function with given fields: _a0
func (_m *UtilsInterface) GetMaxCommission(_a0 *ethclient.Client) (uint8, error) {
	ret := _m.Called(_a0)

	var r0 uint8
	if rf, ok := ret.Get(0).(func(*ethclient.Client) uint8); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(uint8)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNumActiveCollections provides a mock function with given fields: _a0
func (_m *UtilsInterface) GetNumActiveCollections(_a0 *ethclient.Client) (uint16, error) {
	ret := _m.Called(_a0)

	var r0 uint16
	if rf, ok := ret.Get(0).(func(*ethclient.Client) uint16); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(uint16)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNumberOfProposedBlocks provides a mock function with given fields: _a0, _a1
func (_m *UtilsInterface) GetNumberOfProposedBlocks(_a0 *ethclient.Client, _a1 uint32) (uint8, error) {
	ret := _m.Called(_a0, _a1)

	var r0 uint8
	if rf, ok := ret.Get(0).(func(*ethclient.Client, uint32) uint8); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(uint8)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, uint32) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNumberOfStakers provides a mock function with given fields: _a0
func (_m *UtilsInterface) GetNumberOfStakers(_a0 *ethclient.Client) (uint32, error) {
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

// GetOptions provides a mock function with given fields:
func (_m *UtilsInterface) GetOptions() bind.CallOpts {
	ret := _m.Called()

	var r0 bind.CallOpts
	if rf, ok := ret.Get(0).(func() bind.CallOpts); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bind.CallOpts)
	}

	return r0
}

// GetProposedBlock provides a mock function with given fields: _a0, _a1, _a2
func (_m *UtilsInterface) GetProposedBlock(_a0 *ethclient.Client, _a1 uint32, _a2 uint32) (bindings.StructsBlock, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 bindings.StructsBlock
	if rf, ok := ret.Get(0).(func(*ethclient.Client, uint32, uint32) bindings.StructsBlock); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).(bindings.StructsBlock)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, uint32, uint32) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRogueRandomValue provides a mock function with given fields: _a0
func (_m *UtilsInterface) GetRogueRandomValue(_a0 int) *big.Int {
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

// GetSortedProposedBlockIds provides a mock function with given fields: _a0, _a1
func (_m *UtilsInterface) GetSortedProposedBlockIds(_a0 *ethclient.Client, _a1 uint32) ([]uint32, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []uint32
	if rf, ok := ret.Get(0).(func(*ethclient.Client, uint32) []uint32); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]uint32)
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

// GetStake provides a mock function with given fields: _a0, _a1
func (_m *UtilsInterface) GetStake(_a0 *ethclient.Client, _a1 uint32) (*big.Int, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*ethclient.Client, uint32) *big.Int); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
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

// GetStakeSnapshot provides a mock function with given fields: _a0, _a1, _a2
func (_m *UtilsInterface) GetStakeSnapshot(_a0 *ethclient.Client, _a1 uint32, _a2 uint32) (*big.Int, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*ethclient.Client, uint32, uint32) *big.Int); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, uint32, uint32) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStakedToken provides a mock function with given fields: _a0, _a1
func (_m *UtilsInterface) GetStakedToken(_a0 *ethclient.Client, _a1 common.Address) *bindings.StakedToken {
	ret := _m.Called(_a0, _a1)

	var r0 *bindings.StakedToken
	if rf, ok := ret.Get(0).(func(*ethclient.Client, common.Address) *bindings.StakedToken); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bindings.StakedToken)
		}
	}

	return r0
}

// GetStaker provides a mock function with given fields: _a0, _a1
func (_m *UtilsInterface) GetStaker(_a0 *ethclient.Client, _a1 uint32) (bindings.StructsStaker, error) {
	ret := _m.Called(_a0, _a1)

	var r0 bindings.StructsStaker
	if rf, ok := ret.Get(0).(func(*ethclient.Client, uint32) bindings.StructsStaker); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(bindings.StructsStaker)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, uint32) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStakerId provides a mock function with given fields: _a0, _a1
func (_m *UtilsInterface) GetStakerId(_a0 *ethclient.Client, _a1 string) (uint32, error) {
	ret := _m.Called(_a0, _a1)

	var r0 uint32
	if rf, ok := ret.Get(0).(func(*ethclient.Client, string) uint32); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStakerSRZRBalance provides a mock function with given fields: _a0, _a1
func (_m *UtilsInterface) GetStakerSRZRBalance(_a0 *ethclient.Client, _a1 bindings.StructsStaker) (*big.Int, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*ethclient.Client, bindings.StructsStaker) *big.Int); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, bindings.StructsStaker) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStringAddress provides a mock function with given fields: _a0
func (_m *UtilsInterface) GetStringAddress(_a0 *pflag.FlagSet) (string, error) {
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

// GetTotalInfluenceRevealed provides a mock function with given fields: _a0, _a1, _a2
func (_m *UtilsInterface) GetTotalInfluenceRevealed(_a0 *ethclient.Client, _a1 uint32, _a2 uint16) (*big.Int, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*ethclient.Client, uint32, uint16) *big.Int); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, uint32, uint16) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTxnOpts provides a mock function with given fields: _a0
func (_m *UtilsInterface) GetTxnOpts(_a0 types.TransactionOptions) *bind.TransactOpts {
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
func (_m *UtilsInterface) GetUint32BountyId(_a0 *pflag.FlagSet) (uint32, error) {
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

// GetUpdatedStaker provides a mock function with given fields: _a0, _a1
func (_m *UtilsInterface) GetUpdatedStaker(_a0 *ethclient.Client, _a1 uint32) (bindings.StructsStaker, error) {
	ret := _m.Called(_a0, _a1)

	var r0 bindings.StructsStaker
	if rf, ok := ret.Get(0).(func(*ethclient.Client, uint32) bindings.StructsStaker); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(bindings.StructsStaker)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, uint32) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetVoteValue provides a mock function with given fields: client, epoch, stakerId, medianIndex
func (_m *UtilsInterface) GetVoteValue(client *ethclient.Client, epoch uint32, stakerId uint32, medianIndex uint16) (uint32, error) {
	ret := _m.Called(client, epoch, stakerId, medianIndex)

	var r0 uint32
	if rf, ok := ret.Get(0).(func(*ethclient.Client, uint32, uint32, uint16) uint32); ok {
		r0 = rf(client, epoch, stakerId, medianIndex)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, uint32, uint32, uint16) error); ok {
		r1 = rf(client, epoch, stakerId, medianIndex)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetWithdrawInitiationPeriod provides a mock function with given fields: _a0
func (_m *UtilsInterface) GetWithdrawInitiationPeriod(_a0 *ethclient.Client) (uint8, error) {
	ret := _m.Called(_a0)

	var r0 uint8
	if rf, ok := ret.Get(0).(func(*ethclient.Client) uint8); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(uint8)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsFlagPassed provides a mock function with given fields: _a0
func (_m *UtilsInterface) IsFlagPassed(_a0 string) bool {
	ret := _m.Called(_a0)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// PasswordPrompt provides a mock function with given fields:
func (_m *UtilsInterface) PasswordPrompt() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// PrivateKeyPrompt provides a mock function with given fields:
func (_m *UtilsInterface) PrivateKeyPrompt() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// ReadDataFromFile provides a mock function with given fields: _a0
func (_m *UtilsInterface) ReadDataFromFile(_a0 string) (uint32, []*big.Int, error) {
	ret := _m.Called(_a0)

	var r0 uint32
	if rf, ok := ret.Get(0).(func(string) uint32); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	var r1 []*big.Int
	if rf, ok := ret.Get(1).(func(string) []*big.Int); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]*big.Int)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string) error); ok {
		r2 = rf(_a0)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// SaveDataToFile provides a mock function with given fields: _a0, _a1, _a2
func (_m *UtilsInterface) SaveDataToFile(_a0 string, _a1 uint32, _a2 []*big.Int) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, uint32, []*big.Int) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SecondsToReadableTime provides a mock function with given fields: _a0
func (_m *UtilsInterface) SecondsToReadableTime(_a0 int) string {
	ret := _m.Called(_a0)

	var r0 string
	if rf, ok := ret.Get(0).(func(int) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// WaitForBlockCompletion provides a mock function with given fields: _a0, _a1
func (_m *UtilsInterface) WaitForBlockCompletion(_a0 *ethclient.Client, _a1 string) int {
	ret := _m.Called(_a0, _a1)

	var r0 int
	if rf, ok := ret.Get(0).(func(*ethclient.Client, string) int); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// WaitTillNextNSecs provides a mock function with given fields: _a0
func (_m *UtilsInterface) WaitTillNextNSecs(_a0 int32) {
	_m.Called(_a0)
}
