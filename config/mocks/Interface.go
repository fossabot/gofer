// Code generated by mockery v2.4.0-beta. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Interface is an autogenerated mock type for the Interface type
type Interface struct {
	mock.Mock
}

// Parse provides a mock function with given fields: _a0, options
func (_m *Interface) Parse(_a0 interface{}, options interface{}) error {
	ret := _m.Called(_a0, options)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}, interface{}) error); ok {
		r0 = rf(_a0, options)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}