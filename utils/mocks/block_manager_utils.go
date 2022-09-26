// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	big "math/big"
	bindings "razor/pkg/bindings"

	ethclient "github.com/ethereum/go-ethereum/ethclient"

	mock "github.com/stretchr/testify/mock"
)

// BlockManagerUtils is an autogenerated mock type for the BlockManagerUtils type
type BlockManagerUtils struct {
	mock.Mock
}

// GetBlock provides a mock function with given fields: client, epoch
func (_m *BlockManagerUtils) GetBlock(client *ethclient.Client, epoch uint32) (bindings.StructsBlock, error) {
	ret := _m.Called(client, epoch)

	var r0 bindings.StructsBlock
	if rf, ok := ret.Get(0).(func(*ethclient.Client, uint32) bindings.StructsBlock); ok {
		r0 = rf(client, epoch)
	} else {
		r0 = ret.Get(0).(bindings.StructsBlock)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, uint32) error); ok {
		r1 = rf(client, epoch)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBlockIndexToBeConfirmed provides a mock function with given fields: client
func (_m *BlockManagerUtils) GetBlockIndexToBeConfirmed(client *ethclient.Client) (int8, error) {
	ret := _m.Called(client)

	var r0 int8
	if rf, ok := ret.Get(0).(func(*ethclient.Client) int8); ok {
		r0 = rf(client)
	} else {
		r0 = ret.Get(0).(int8)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client) error); ok {
		r1 = rf(client)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEpochLastProposed provides a mock function with given fields: client, stakerId
func (_m *BlockManagerUtils) GetEpochLastProposed(client *ethclient.Client, stakerId uint32) (uint32, error) {
	ret := _m.Called(client, stakerId)

	var r0 uint32
	if rf, ok := ret.Get(0).(func(*ethclient.Client, uint32) uint32); ok {
		r0 = rf(client, stakerId)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, uint32) error); ok {
		r1 = rf(client, stakerId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNumProposedBlocks provides a mock function with given fields: client, epoch
func (_m *BlockManagerUtils) GetNumProposedBlocks(client *ethclient.Client, epoch uint32) (uint8, error) {
	ret := _m.Called(client, epoch)

	var r0 uint8
	if rf, ok := ret.Get(0).(func(*ethclient.Client, uint32) uint8); ok {
		r0 = rf(client, epoch)
	} else {
		r0 = ret.Get(0).(uint8)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, uint32) error); ok {
		r1 = rf(client, epoch)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProposedBlock provides a mock function with given fields: client, epoch, proposedBlock
func (_m *BlockManagerUtils) GetProposedBlock(client *ethclient.Client, epoch uint32, proposedBlock uint32) (bindings.StructsBlock, error) {
	ret := _m.Called(client, epoch, proposedBlock)

	var r0 bindings.StructsBlock
	if rf, ok := ret.Get(0).(func(*ethclient.Client, uint32, uint32) bindings.StructsBlock); ok {
		r0 = rf(client, epoch, proposedBlock)
	} else {
		r0 = ret.Get(0).(bindings.StructsBlock)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, uint32, uint32) error); ok {
		r1 = rf(client, epoch, proposedBlock)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MaxAltBlocks provides a mock function with given fields: client
func (_m *BlockManagerUtils) MaxAltBlocks(client *ethclient.Client) (uint8, error) {
	ret := _m.Called(client)

	var r0 uint8
	if rf, ok := ret.Get(0).(func(*ethclient.Client) uint8); ok {
		r0 = rf(client)
	} else {
		r0 = ret.Get(0).(uint8)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client) error); ok {
		r1 = rf(client)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MinStake provides a mock function with given fields: client
func (_m *BlockManagerUtils) MinStake(client *ethclient.Client) (*big.Int, error) {
	ret := _m.Called(client)

	var r0 *big.Int
	if rf, ok := ret.Get(0).(func(*ethclient.Client) *big.Int); ok {
		r0 = rf(client)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client) error); ok {
		r1 = rf(client)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SortedProposedBlockIds provides a mock function with given fields: client, arg0, arg1
func (_m *BlockManagerUtils) SortedProposedBlockIds(client *ethclient.Client, arg0 uint32, arg1 *big.Int) (uint32, error) {
	ret := _m.Called(client, arg0, arg1)

	var r0 uint32
	if rf, ok := ret.Get(0).(func(*ethclient.Client, uint32, *big.Int) uint32); ok {
		r0 = rf(client, arg0, arg1)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client, uint32, *big.Int) error); ok {
		r1 = rf(client, arg0, arg1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StateBuffer provides a mock function with given fields: client
func (_m *BlockManagerUtils) StateBuffer(client *ethclient.Client) (uint8, error) {
	ret := _m.Called(client)

	var r0 uint8
	if rf, ok := ret.Get(0).(func(*ethclient.Client) uint8); ok {
		r0 = rf(client)
	} else {
		r0 = ret.Get(0).(uint8)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*ethclient.Client) error); ok {
		r1 = rf(client)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
