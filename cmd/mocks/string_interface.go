// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// StringInterface is an autogenerated mock type for the StringInterface type
type StringInterface struct {
	mock.Mock
}

// ParseBool provides a mock function with given fields: str
func (_m *StringInterface) ParseBool(str string) (bool, error) {
	ret := _m.Called(str)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(str)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(str)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParseChainId provides a mock function with given fields: str
func (_m *StringInterface) ParseChainId(str string) (int64, error) {
	ret := _m.Called(str)

	var r0 int64
	if rf, ok := ret.Get(0).(func(string) int64); ok {
		r0 = rf(str)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(str)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParseFloat provides a mock function with given fields: str
func (_m *StringInterface) ParseFloat(str string) (float64, error) {
	ret := _m.Called(str)

	var r0 float64
	if rf, ok := ret.Get(0).(func(string) float64); ok {
		r0 = rf(str)
	} else {
		r0 = ret.Get(0).(float64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(str)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParseInt provides a mock function with given fields: str
func (_m *StringInterface) ParseInt(str string) (int, error) {
	ret := _m.Called(str)

	var r0 int
	if rf, ok := ret.Get(0).(func(string) int); ok {
		r0 = rf(str)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(str)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParseInt64 provides a mock function with given fields: str
func (_m *StringInterface) ParseInt64(str string) (int64, error) {
	ret := _m.Called(str)

	var r0 int64
	if rf, ok := ret.Get(0).(func(string) int64); ok {
		r0 = rf(str)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(str)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewStringInterface interface {
	mock.TestingT
	Cleanup(func())
}

// NewStringInterface creates a new instance of StringInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewStringInterface(t mockConstructorTestingTNewStringInterface) *StringInterface {
	mock := &StringInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
