// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	big "math/big"

	mock "github.com/stretchr/testify/mock"
)

// MerkleTreeInterface is an autogenerated mock type for the MerkleTreeInterface type
type MerkleTreeInterface struct {
	mock.Mock
}

// CreateMerkle provides a mock function with given fields: values
func (_m *MerkleTreeInterface) CreateMerkle(values []*big.Int) ([][][]byte, error) {
	ret := _m.Called(values)

	var r0 [][][]byte
	if rf, ok := ret.Get(0).(func([]*big.Int) [][][]byte); ok {
		r0 = rf(values)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([][][]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]*big.Int) error); ok {
		r1 = rf(values)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMerkleRoot provides a mock function with given fields: tree
func (_m *MerkleTreeInterface) GetMerkleRoot(tree [][][]byte) [32]byte {
	ret := _m.Called(tree)

	var r0 [32]byte
	if rf, ok := ret.Get(0).(func([][][]byte) [32]byte); ok {
		r0 = rf(tree)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([32]byte)
		}
	}

	return r0
}

// GetProofPath provides a mock function with given fields: tree, assetId
func (_m *MerkleTreeInterface) GetProofPath(tree [][][]byte, assetId uint16) [][32]byte {
	ret := _m.Called(tree, assetId)

	var r0 [][32]byte
	if rf, ok := ret.Get(0).(func([][][]byte, uint16) [][32]byte); ok {
		r0 = rf(tree, assetId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([][32]byte)
		}
	}

	return r0
}

type mockConstructorTestingTNewMerkleTreeInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewMerkleTreeInterface creates a new instance of MerkleTreeInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMerkleTreeInterface(t mockConstructorTestingTNewMerkleTreeInterface) *MerkleTreeInterface {
	mock := &MerkleTreeInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
