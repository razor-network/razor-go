// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	big "math/big"
	bindings "razor/pkg/bindings"

	bind "github.com/ethereum/go-ethereum/accounts/abi/bind"

	common "github.com/ethereum/go-ethereum/common"

	coretypes "razor/core/types"

	ethclient "github.com/ethereum/go-ethereum/ethclient"

	ethereum "github.com/ethereum/go-ethereum"

	mock "github.com/stretchr/testify/mock"

	types "github.com/ethereum/go-ethereum/core/types"
)

// Utils is an autogenerated mock type for the Utils type
type Utils struct {
	mock.Mock
}

// Aggregate provides a mock function with given fields: _a0, _a1, _a2
func (_m *Utils) Aggregate(_a0 *ethclient.Client, _a1 uint32, _a2 bindings.StructsCollection) (*big.Int, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*ethclient.Client, uint32, bindings.StructsCollection) *big.Int); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, uint32, bindings.StructsCollection) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BalanceAtWithRetry provides a mock function with given fields: _a0, _a1
func (_m *Utils) BalanceAtWithRetry(_a0 *ethclient.Client, _a1 common.Address) (*big.Int, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*ethclient.Client, common.Address) *big.Int); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, common.Address) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// EstimateGasWithRetry provides a mock function with given fields: _a0, _a1
func (_m *Utils) EstimateGasWithRetry(_a0 *ethclient.Client, _a1 ethereum.CallMsg) (uint64, error) {
	ret := _m.Called(_a0, _a1)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(*ethclient.Client, ethereum.CallMsg) uint64); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, ethereum.CallMsg) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FetchPreviousValue provides a mock function with given fields: _a0, _a1, _a2
func (_m *Utils) FetchPreviousValue(_a0 *ethclient.Client, _a1 uint32, _a2 uint16) (uint32, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 uint32
	if rf, ok := ret.Get(0).(func(*ethclient.Client, uint32, uint16) uint32); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, uint32, uint16) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FilterLogsWithRetry provides a mock function with given fields: _a0, _a1
func (_m *Utils) FilterLogsWithRetry(_a0 *ethclient.Client, _a1 ethereum.FilterQuery) ([]types.Log, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []types.Log
	if rf, ok := ret.Get(0).(func(*ethclient.Client, ethereum.FilterQuery) []types.Log); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.Log)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, ethereum.FilterQuery) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetActiveAssetIds provides a mock function with given fields: _a0
func (_m *Utils) GetActiveAssetIds(_a0 *ethclient.Client) ([]uint16, error) {
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

// GetActiveAssetsData provides a mock function with given fields: _a0, _a1
func (_m *Utils) GetActiveAssetsData(_a0 *ethclient.Client, _a1 uint32) ([]*big.Int, error) {
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

// GetActiveCollection provides a mock function with given fields: _a0, _a1
func (_m *Utils) GetActiveCollection(_a0 *ethclient.Client, _a1 uint16) (bindings.StructsCollection, error) {
	ret := _m.Called(_a0, _a1)

	var r0 bindings.StructsCollection
	if rf, ok := ret.Get(0).(func(*ethclient.Client, uint16) bindings.StructsCollection); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(bindings.StructsCollection)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, uint16) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetActiveJob provides a mock function with given fields: _a0, _a1
func (_m *Utils) GetActiveJob(_a0 *ethclient.Client, _a1 uint16) (bindings.StructsJob, error) {
	ret := _m.Called(_a0, _a1)

	var r0 bindings.StructsJob
	if rf, ok := ret.Get(0).(func(*ethclient.Client, uint16) bindings.StructsJob); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(bindings.StructsJob)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, uint16) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAssetManager provides a mock function with given fields: _a0
func (_m *Utils) GetAssetManager(_a0 *ethclient.Client) *bindings.AssetManager {
	ret := _m.Called(_a0)

	var r0 *bindings.AssetManager
	if rf, ok := ret.Get(0).(func(*ethclient.Client) *bindings.AssetManager); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bindings.AssetManager)
		}
	}

	return r0
}

// GetAssetManagerWithOpts provides a mock function with given fields: _a0
func (_m *Utils) GetAssetManagerWithOpts(_a0 *ethclient.Client) (*bindings.AssetManager, bind.CallOpts) {
	ret := _m.Called(_a0)

	var r0 *bindings.AssetManager
	if rf, ok := ret.Get(0).(func(*ethclient.Client) *bindings.AssetManager); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bindings.AssetManager)
		}
	}

	var r1 bind.CallOpts
	if rf, ok := ret.Get(1).(func(*ethclient.Client) bind.CallOpts); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Get(1).(bind.CallOpts)
	}

	return r0, r1
}

// GetAssetType provides a mock function with given fields: _a0, _a1
func (_m *Utils) GetAssetType(_a0 *ethclient.Client, _a1 uint16) (uint8, error) {
	ret := _m.Called(_a0, _a1)

	var r0 uint8
	if rf, ok := ret.Get(0).(func(*ethclient.Client, uint16) uint8); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(uint8)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, uint16) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBlockManager provides a mock function with given fields: _a0
func (_m *Utils) GetBlockManager(_a0 *ethclient.Client) *bindings.BlockManager {
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

// GetBlockManagerWithOpts provides a mock function with given fields: _a0
func (_m *Utils) GetBlockManagerWithOpts(_a0 *ethclient.Client) (*bindings.BlockManager, bind.CallOpts) {
	ret := _m.Called(_a0)

	var r0 *bindings.BlockManager
	if rf, ok := ret.Get(0).(func(*ethclient.Client) *bindings.BlockManager); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bindings.BlockManager)
		}
	}

	var r1 bind.CallOpts
	if rf, ok := ret.Get(1).(func(*ethclient.Client) bind.CallOpts); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Get(1).(bind.CallOpts)
	}

	return r0, r1
}

// GetCollection provides a mock function with given fields: _a0, _a1
func (_m *Utils) GetCollection(_a0 *ethclient.Client, _a1 uint16) (bindings.StructsCollection, error) {
	ret := _m.Called(_a0, _a1)

	var r0 bindings.StructsCollection
	if rf, ok := ret.Get(0).(func(*ethclient.Client, uint16) bindings.StructsCollection); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(bindings.StructsCollection)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, uint16) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCollections provides a mock function with given fields: _a0
func (_m *Utils) GetCollections(_a0 *ethclient.Client) ([]bindings.StructsCollection, error) {
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
func (_m *Utils) GetCommitments(_a0 *ethclient.Client, _a1 string) ([32]byte, error) {
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

// GetDataFromAPI provides a mock function with given fields: _a0
func (_m *Utils) GetDataFromAPI(_a0 string) ([]byte, error) {
	ret := _m.Called(_a0)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(string) []byte); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDataFromHTML provides a mock function with given fields: _a0, _a1
func (_m *Utils) GetDataFromHTML(_a0 string, _a1 string) (string, error) {
	ret := _m.Called(_a0, _a1)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDataFromJSON provides a mock function with given fields: _a0, _a1
func (_m *Utils) GetDataFromJSON(_a0 map[string]interface{}, _a1 string) (interface{}, error) {
	ret := _m.Called(_a0, _a1)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(map[string]interface{}, string) interface{}); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(map[string]interface{}, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDataToCommitFromJob provides a mock function with given fields: _a0
func (_m *Utils) GetDataToCommitFromJob(_a0 bindings.StructsJob) (*big.Int, error) {
	ret := _m.Called(_a0)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(bindings.StructsJob) *big.Int); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(bindings.StructsJob) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDataToCommitFromJobs provides a mock function with given fields: _a0
func (_m *Utils) GetDataToCommitFromJobs(_a0 []bindings.StructsJob) ([]*big.Int, []uint8, error) {
	ret := _m.Called(_a0)

	var r0 []*big.Int
	if rf, ok := ret.Get(0).(func([]bindings.StructsJob) []*big.Int); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*big.Int)
		}
	}

	var r1 []uint8
	if rf, ok := ret.Get(1).(func([]bindings.StructsJob) []uint8); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]uint8)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func([]bindings.StructsJob) error); ok {
		r2 = rf(_a0)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetEpochLastCommitted provides a mock function with given fields: _a0, _a1
func (_m *Utils) GetEpochLastCommitted(_a0 *ethclient.Client, _a1 uint32) (uint32, error) {
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
func (_m *Utils) GetEpochLastRevealed(_a0 *ethclient.Client, _a1 uint32) (uint32, error) {
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

// GetGasLimit provides a mock function with given fields: _a0, _a1
func (_m *Utils) GetGasLimit(_a0 coretypes.TransactionOptions, _a1 *bind.TransactOpts) (uint64, error) {
	ret := _m.Called(_a0, _a1)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(coretypes.TransactionOptions, *bind.TransactOpts) uint64); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(coretypes.TransactionOptions, *bind.TransactOpts) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetGasPrice provides a mock function with given fields: _a0, _a1
func (_m *Utils) GetGasPrice(_a0 *ethclient.Client, _a1 coretypes.Configurations) *big.Int {
	ret := _m.Called(_a0, _a1)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*ethclient.Client, coretypes.Configurations) *big.Int); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	return r0
}

// GetInfluenceSnapshot provides a mock function with given fields: _a0, _a1, _a2
func (_m *Utils) GetInfluenceSnapshot(_a0 *ethclient.Client, _a1 uint32, _a2 uint32) (*big.Int, error) {
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

// GetJobs provides a mock function with given fields: _a0
func (_m *Utils) GetJobs(_a0 *ethclient.Client) ([]bindings.StructsJob, error) {
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

// GetLatestBlockWithRetry provides a mock function with given fields: _a0
func (_m *Utils) GetLatestBlockWithRetry(_a0 *ethclient.Client) (*types.Header, error) {
	ret := _m.Called(_a0)

	var r0 *types.Header
	if rf, ok := ret.Get(0).(func(*ethclient.Client) *types.Header); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Header)
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

// GetMaxAltBlocks provides a mock function with given fields: _a0
func (_m *Utils) GetMaxAltBlocks(_a0 *ethclient.Client) (uint8, error) {
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

// GetMinStakeAmount provides a mock function with given fields: _a0
func (_m *Utils) GetMinStakeAmount(_a0 *ethclient.Client) (*big.Int, error) {
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

// GetNumActiveAssets provides a mock function with given fields: _a0
func (_m *Utils) GetNumActiveAssets(_a0 *ethclient.Client) (*big.Int, error) {
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

// GetNumAssets provides a mock function with given fields: _a0
func (_m *Utils) GetNumAssets(_a0 *ethclient.Client) (uint16, error) {
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
func (_m *Utils) GetNumberOfProposedBlocks(_a0 *ethclient.Client, _a1 uint32) (uint8, error) {
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

// GetOptions provides a mock function with given fields:
func (_m *Utils) GetOptions() bind.CallOpts {
	ret := _m.Called()

	var r0 bind.CallOpts
	if rf, ok := ret.Get(0).(func() bind.CallOpts); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bind.CallOpts)
	}

	return r0
}

// GetPendingNonceAtWithRetry provides a mock function with given fields: _a0, _a1
func (_m *Utils) GetPendingNonceAtWithRetry(_a0 *ethclient.Client, _a1 common.Address) (uint64, error) {
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

// GetProposedBlock provides a mock function with given fields: _a0, _a1, _a2
func (_m *Utils) GetProposedBlock(_a0 *ethclient.Client, _a1 uint32, _a2 uint32) (bindings.StructsBlock, error) {
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

// GetRandaoHash provides a mock function with given fields: _a0
func (_m *Utils) GetRandaoHash(_a0 *ethclient.Client) ([32]byte, error) {
	ret := _m.Called(_a0)

	var r0 [32]byte
	if rf, ok := ret.Get(0).(func(*ethclient.Client) [32]byte); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([32]byte)
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

// GetSortedProposedBlockId provides a mock function with given fields: _a0, _a1, _a2
func (_m *Utils) GetSortedProposedBlockId(_a0 *ethclient.Client, _a1 uint32, _a2 *big.Int) (uint32, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 uint32
	if rf, ok := ret.Get(0).(func(*ethclient.Client, uint32, *big.Int) uint32); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, uint32, *big.Int) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSortedProposedBlockIds provides a mock function with given fields: _a0, _a1
func (_m *Utils) GetSortedProposedBlockIds(_a0 *ethclient.Client, _a1 uint32) ([]uint32, error) {
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

// GetStakeManager provides a mock function with given fields: _a0
func (_m *Utils) GetStakeManager(_a0 *ethclient.Client) *bindings.StakeManager {
	ret := _m.Called(_a0)

	var r0 *bindings.StakeManager
	if rf, ok := ret.Get(0).(func(*ethclient.Client) *bindings.StakeManager); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bindings.StakeManager)
		}
	}

	return r0
}

// GetStakeSnapshot provides a mock function with given fields: _a0, _a1, _a2
func (_m *Utils) GetStakeSnapshot(_a0 *ethclient.Client, _a1 uint32, _a2 uint32) (*big.Int, error) {
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
func (_m *Utils) GetStakedToken(_a0 *ethclient.Client, _a1 common.Address) *bindings.StakedToken {
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

// GetStakerId provides a mock function with given fields: _a0, _a1
func (_m *Utils) GetStakerId(_a0 *ethclient.Client, _a1 string) (uint32, error) {
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

// GetTokenManager provides a mock function with given fields: _a0
func (_m *Utils) GetTokenManager(_a0 *ethclient.Client) *bindings.RAZOR {
	ret := _m.Called(_a0)

	var r0 *bindings.RAZOR
	if rf, ok := ret.Get(0).(func(*ethclient.Client) *bindings.RAZOR); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bindings.RAZOR)
		}
	}

	return r0
}

// GetTotalInfluenceRevealed provides a mock function with given fields: _a0, _a1
func (_m *Utils) GetTotalInfluenceRevealed(_a0 *ethclient.Client, _a1 uint32) (*big.Int, error) {
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

// GetTxnOpts provides a mock function with given fields: _a0
func (_m *Utils) GetTxnOpts(_a0 coretypes.TransactionOptions) *bind.TransactOpts {
	ret := _m.Called(_a0)

	var r0 *bind.TransactOpts
	if rf, ok := ret.Get(0).(func(coretypes.TransactionOptions) *bind.TransactOpts); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bind.TransactOpts)
		}
	}

	return r0
}

// GetVoteManager provides a mock function with given fields: _a0
func (_m *Utils) GetVoteManager(_a0 *ethclient.Client) *bindings.VoteManager {
	ret := _m.Called(_a0)

	var r0 *bindings.VoteManager
	if rf, ok := ret.Get(0).(func(*ethclient.Client) *bindings.VoteManager); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bindings.VoteManager)
		}
	}

	return r0
}

// GetVoteManagerWithOpts provides a mock function with given fields: _a0
func (_m *Utils) GetVoteManagerWithOpts(_a0 *ethclient.Client) (*bindings.VoteManager, bind.CallOpts) {
	ret := _m.Called(_a0)

	var r0 *bindings.VoteManager
	if rf, ok := ret.Get(0).(func(*ethclient.Client) *bindings.VoteManager); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*bindings.VoteManager)
		}
	}

	var r1 bind.CallOpts
	if rf, ok := ret.Get(1).(func(*ethclient.Client) bind.CallOpts); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Get(1).(bind.CallOpts)
	}

	return r0, r1
}

// GetVoteValue provides a mock function with given fields: _a0, _a1, _a2
func (_m *Utils) GetVoteValue(_a0 *ethclient.Client, _a1 uint16, _a2 uint32) (*big.Int, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*ethclient.Client, uint16, uint32) *big.Int); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, uint16, uint32) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetVotes provides a mock function with given fields: _a0, _a1
func (_m *Utils) GetVotes(_a0 *ethclient.Client, _a1 uint32) (bindings.StructsVote, error) {
	ret := _m.Called(_a0, _a1)

	var r0 bindings.StructsVote
	if rf, ok := ret.Get(0).(func(*ethclient.Client, uint32) bindings.StructsVote); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(bindings.StructsVote)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, uint32) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IncreaseGasLimitValue provides a mock function with given fields: _a0, _a1, _a2
func (_m *Utils) IncreaseGasLimitValue(_a0 *ethclient.Client, _a1 uint64, _a2 float32) (uint64, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(*ethclient.Client, uint64, float32) uint64); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, uint64, float32) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MultiplyFloatAndBigInt provides a mock function with given fields: _a0, _a1
func (_m *Utils) MultiplyFloatAndBigInt(_a0 *big.Int, _a1 float64) *big.Int {
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
func (_m *Utils) SuggestGasPriceWithRetry(_a0 *ethclient.Client) (*big.Int, error) {
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
