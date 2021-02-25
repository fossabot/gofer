// Code generated by mockery v2.4.0-beta. DO NOT EDIT.

package mocks

import (
	query "github.com/patrickascher/gofer/query"
	mock "github.com/stretchr/testify/mock"
)

// Information is an autogenerated mock type for the Information type
type Information struct {
	mock.Mock
}

// Describe provides a mock function with given fields: columns
func (_m *Information) Describe(columns ...string) ([]query.Column, error) {
	_va := make([]interface{}, len(columns))
	for _i := range columns {
		_va[_i] = columns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []query.Column
	if rf, ok := ret.Get(0).(func(...string) []query.Column); ok {
		r0 = rf(columns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]query.Column)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(...string) error); ok {
		r1 = rf(columns...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ForeignKey provides a mock function with given fields:
func (_m *Information) ForeignKey() ([]query.ForeignKey, error) {
	ret := _m.Called()

	var r0 []query.ForeignKey
	if rf, ok := ret.Get(0).(func() []query.ForeignKey); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]query.ForeignKey)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}