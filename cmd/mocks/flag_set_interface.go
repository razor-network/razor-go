// Code generated by mockery v2.30.1. DO NOT EDIT.

package mocks

import (
	pflag "github.com/spf13/pflag"
	mock "github.com/stretchr/testify/mock"
)

// FlagSetInterface is an autogenerated mock type for the FlagSetInterface type
type FlagSetInterface struct {
	mock.Mock
}

// Changed provides a mock function with given fields: flagSet, flagName
func (_m *FlagSetInterface) Changed(flagSet *pflag.FlagSet, flagName string) bool {
	ret := _m.Called(flagSet, flagName)

	var r0 bool
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet, string) bool); ok {
		r0 = rf(flagSet, flagName)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// FetchFlagInput provides a mock function with given fields: flagSet, flagKeyword, dataType
func (_m *FlagSetInterface) FetchFlagInput(flagSet *pflag.FlagSet, flagKeyword string, dataType string) (interface{}, error) {
	ret := _m.Called(flagSet, flagKeyword, dataType)

	var r0 interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet, string, string) (interface{}, error)); ok {
		return rf(flagSet, flagKeyword, dataType)
	}
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet, string, string) interface{}); ok {
		r0 = rf(flagSet, flagKeyword, dataType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(*pflag.FlagSet, string, string) error); ok {
		r1 = rf(flagSet, flagKeyword, dataType)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FetchRootFlagInput provides a mock function with given fields: flagName, dataType
func (_m *FlagSetInterface) FetchRootFlagInput(flagName string, dataType string) (interface{}, error) {
	ret := _m.Called(flagName, dataType)

	var r0 interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (interface{}, error)); ok {
		return rf(flagName, dataType)
	}
	if rf, ok := ret.Get(0).(func(string, string) interface{}); ok {
		r0 = rf(flagName, dataType)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(flagName, dataType)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBoolRogue provides a mock function with given fields: flagSet
func (_m *FlagSetInterface) GetBoolRogue(flagSet *pflag.FlagSet) (bool, error) {
	ret := _m.Called(flagSet)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) (bool, error)); ok {
		return rf(flagSet)
	}
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) bool); ok {
		r0 = rf(flagSet)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(flagSet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBoolWeiRazor provides a mock function with given fields: flagSet
func (_m *FlagSetInterface) GetBoolWeiRazor(flagSet *pflag.FlagSet) (bool, error) {
	ret := _m.Called(flagSet)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) (bool, error)); ok {
		return rf(flagSet)
	}
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) bool); ok {
		r0 = rf(flagSet)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(flagSet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetInt8Power provides a mock function with given fields: flagSet
func (_m *FlagSetInterface) GetInt8Power(flagSet *pflag.FlagSet) (int8, error) {
	ret := _m.Called(flagSet)

	var r0 int8
	var r1 error
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) (int8, error)); ok {
		return rf(flagSet)
	}
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) int8); ok {
		r0 = rf(flagSet)
	} else {
		r0 = ret.Get(0).(int8)
	}

	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(flagSet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetIntLogFileMaxAge provides a mock function with given fields: flagSet
func (_m *FlagSetInterface) GetIntLogFileMaxAge(flagSet *pflag.FlagSet) (int, error) {
	ret := _m.Called(flagSet)

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) (int, error)); ok {
		return rf(flagSet)
	}
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) int); ok {
		r0 = rf(flagSet)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(flagSet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetIntLogFileMaxBackups provides a mock function with given fields: flagSet
func (_m *FlagSetInterface) GetIntLogFileMaxBackups(flagSet *pflag.FlagSet) (int, error) {
	ret := _m.Called(flagSet)

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) (int, error)); ok {
		return rf(flagSet)
	}
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) int); ok {
		r0 = rf(flagSet)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(flagSet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetIntLogFileMaxSize provides a mock function with given fields: flagSet
func (_m *FlagSetInterface) GetIntLogFileMaxSize(flagSet *pflag.FlagSet) (int, error) {
	ret := _m.Called(flagSet)

	var r0 int
	var r1 error
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) (int, error)); ok {
		return rf(flagSet)
	}
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) int); ok {
		r0 = rf(flagSet)
	} else {
		r0 = ret.Get(0).(int)
	}

	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(flagSet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStringAddress provides a mock function with given fields: flagSet
func (_m *FlagSetInterface) GetStringAddress(flagSet *pflag.FlagSet) (string, error) {
	ret := _m.Called(flagSet)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) (string, error)); ok {
		return rf(flagSet)
	}
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) string); ok {
		r0 = rf(flagSet)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(flagSet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStringCertFile provides a mock function with given fields: flagSet
func (_m *FlagSetInterface) GetStringCertFile(flagSet *pflag.FlagSet) (string, error) {
	ret := _m.Called(flagSet)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) (string, error)); ok {
		return rf(flagSet)
	}
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) string); ok {
		r0 = rf(flagSet)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(flagSet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStringCertKey provides a mock function with given fields: flagSet
func (_m *FlagSetInterface) GetStringCertKey(flagSet *pflag.FlagSet) (string, error) {
	ret := _m.Called(flagSet)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) (string, error)); ok {
		return rf(flagSet)
	}
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) string); ok {
		r0 = rf(flagSet)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(flagSet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStringExposeMetrics provides a mock function with given fields: flagSet
func (_m *FlagSetInterface) GetStringExposeMetrics(flagSet *pflag.FlagSet) (string, error) {
	ret := _m.Called(flagSet)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) (string, error)); ok {
		return rf(flagSet)
	}
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) string); ok {
		r0 = rf(flagSet)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(flagSet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStringFrom provides a mock function with given fields: flagSet
func (_m *FlagSetInterface) GetStringFrom(flagSet *pflag.FlagSet) (string, error) {
	ret := _m.Called(flagSet)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) (string, error)); ok {
		return rf(flagSet)
	}
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) string); ok {
		r0 = rf(flagSet)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(flagSet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStringName provides a mock function with given fields: flagSet
func (_m *FlagSetInterface) GetStringName(flagSet *pflag.FlagSet) (string, error) {
	ret := _m.Called(flagSet)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) (string, error)); ok {
		return rf(flagSet)
	}
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) string); ok {
		r0 = rf(flagSet)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(flagSet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStringSelector provides a mock function with given fields: flagSet
func (_m *FlagSetInterface) GetStringSelector(flagSet *pflag.FlagSet) (string, error) {
	ret := _m.Called(flagSet)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) (string, error)); ok {
		return rf(flagSet)
	}
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) string); ok {
		r0 = rf(flagSet)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(flagSet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStringSliceBackupNode provides a mock function with given fields: flagSet
func (_m *FlagSetInterface) GetStringSliceBackupNode(flagSet *pflag.FlagSet) ([]string, error) {
	ret := _m.Called(flagSet)

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) ([]string, error)); ok {
		return rf(flagSet)
	}
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) []string); ok {
		r0 = rf(flagSet)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(flagSet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStringSliceRogueMode provides a mock function with given fields: flagSet
func (_m *FlagSetInterface) GetStringSliceRogueMode(flagSet *pflag.FlagSet) ([]string, error) {
	ret := _m.Called(flagSet)

	var r0 []string
	var r1 error
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) ([]string, error)); ok {
		return rf(flagSet)
	}
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) []string); ok {
		r0 = rf(flagSet)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(flagSet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStringStatus provides a mock function with given fields: flagSet
func (_m *FlagSetInterface) GetStringStatus(flagSet *pflag.FlagSet) (string, error) {
	ret := _m.Called(flagSet)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) (string, error)); ok {
		return rf(flagSet)
	}
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) string); ok {
		r0 = rf(flagSet)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(flagSet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStringTo provides a mock function with given fields: flagSet
func (_m *FlagSetInterface) GetStringTo(flagSet *pflag.FlagSet) (string, error) {
	ret := _m.Called(flagSet)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) (string, error)); ok {
		return rf(flagSet)
	}
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) string); ok {
		r0 = rf(flagSet)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(flagSet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStringUrl provides a mock function with given fields: flagSet
func (_m *FlagSetInterface) GetStringUrl(flagSet *pflag.FlagSet) (string, error) {
	ret := _m.Called(flagSet)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) (string, error)); ok {
		return rf(flagSet)
	}
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) string); ok {
		r0 = rf(flagSet)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(flagSet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetStringValue provides a mock function with given fields: flagSet
func (_m *FlagSetInterface) GetStringValue(flagSet *pflag.FlagSet) (string, error) {
	ret := _m.Called(flagSet)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) (string, error)); ok {
		return rf(flagSet)
	}
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) string); ok {
		r0 = rf(flagSet)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(flagSet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUint16CollectionId provides a mock function with given fields: flagSet
func (_m *FlagSetInterface) GetUint16CollectionId(flagSet *pflag.FlagSet) (uint16, error) {
	ret := _m.Called(flagSet)

	var r0 uint16
	var r1 error
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) (uint16, error)); ok {
		return rf(flagSet)
	}
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) uint16); ok {
		r0 = rf(flagSet)
	} else {
		r0 = ret.Get(0).(uint16)
	}

	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(flagSet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUint16JobId provides a mock function with given fields: flagSet
func (_m *FlagSetInterface) GetUint16JobId(flagSet *pflag.FlagSet) (uint16, error) {
	ret := _m.Called(flagSet)

	var r0 uint16
	var r1 error
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) (uint16, error)); ok {
		return rf(flagSet)
	}
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) uint16); ok {
		r0 = rf(flagSet)
	} else {
		r0 = ret.Get(0).(uint16)
	}

	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(flagSet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUint32Aggregation provides a mock function with given fields: flagSet
func (_m *FlagSetInterface) GetUint32Aggregation(flagSet *pflag.FlagSet) (uint32, error) {
	ret := _m.Called(flagSet)

	var r0 uint32
	var r1 error
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) (uint32, error)); ok {
		return rf(flagSet)
	}
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) uint32); ok {
		r0 = rf(flagSet)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(flagSet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUint32BountyId provides a mock function with given fields: flagSet
func (_m *FlagSetInterface) GetUint32BountyId(flagSet *pflag.FlagSet) (uint32, error) {
	ret := _m.Called(flagSet)

	var r0 uint32
	var r1 error
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) (uint32, error)); ok {
		return rf(flagSet)
	}
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) uint32); ok {
		r0 = rf(flagSet)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(flagSet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUint32StakerId provides a mock function with given fields: flagSet
func (_m *FlagSetInterface) GetUint32StakerId(flagSet *pflag.FlagSet) (uint32, error) {
	ret := _m.Called(flagSet)

	var r0 uint32
	var r1 error
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) (uint32, error)); ok {
		return rf(flagSet)
	}
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) uint32); ok {
		r0 = rf(flagSet)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(flagSet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUint32Tolerance provides a mock function with given fields: flagSet
func (_m *FlagSetInterface) GetUint32Tolerance(flagSet *pflag.FlagSet) (uint32, error) {
	ret := _m.Called(flagSet)

	var r0 uint32
	var r1 error
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) (uint32, error)); ok {
		return rf(flagSet)
	}
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) uint32); ok {
		r0 = rf(flagSet)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(flagSet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUint8Commission provides a mock function with given fields: flagSet
func (_m *FlagSetInterface) GetUint8Commission(flagSet *pflag.FlagSet) (uint8, error) {
	ret := _m.Called(flagSet)

	var r0 uint8
	var r1 error
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) (uint8, error)); ok {
		return rf(flagSet)
	}
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) uint8); ok {
		r0 = rf(flagSet)
	} else {
		r0 = ret.Get(0).(uint8)
	}

	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(flagSet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUint8SelectorType provides a mock function with given fields: flagSet
func (_m *FlagSetInterface) GetUint8SelectorType(flagSet *pflag.FlagSet) (uint8, error) {
	ret := _m.Called(flagSet)

	var r0 uint8
	var r1 error
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) (uint8, error)); ok {
		return rf(flagSet)
	}
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) uint8); ok {
		r0 = rf(flagSet)
	} else {
		r0 = ret.Get(0).(uint8)
	}

	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(flagSet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUint8Weight provides a mock function with given fields: flagSet
func (_m *FlagSetInterface) GetUint8Weight(flagSet *pflag.FlagSet) (uint8, error) {
	ret := _m.Called(flagSet)

	var r0 uint8
	var r1 error
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) (uint8, error)); ok {
		return rf(flagSet)
	}
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) uint8); ok {
		r0 = rf(flagSet)
	} else {
		r0 = ret.Get(0).(uint8)
	}

	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(flagSet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUintSliceJobIds provides a mock function with given fields: flagSet
func (_m *FlagSetInterface) GetUintSliceJobIds(flagSet *pflag.FlagSet) ([]uint, error) {
	ret := _m.Called(flagSet)

	var r0 []uint
	var r1 error
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) ([]uint, error)); ok {
		return rf(flagSet)
	}
	if rf, ok := ret.Get(0).(func(*pflag.FlagSet) []uint); ok {
		r0 = rf(flagSet)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]uint)
		}
	}

	if rf, ok := ret.Get(1).(func(*pflag.FlagSet) error); ok {
		r1 = rf(flagSet)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewFlagSetInterface creates a new instance of FlagSetInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewFlagSetInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *FlagSetInterface {
	mock := &FlagSetInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}