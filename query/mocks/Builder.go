// Code generated by mockery v2.4.0-beta. DO NOT EDIT.

package mocks

import (
	logger "github.com/patrickascher/gofer/logger"
	query "github.com/patrickascher/gofer/query"
	mock "github.com/stretchr/testify/mock"
)

// Builder is an autogenerated mock type for the Builder type
type Builder struct {
	mock.Mock
}

// Config provides a mock function with given fields:
func (_m *Builder) Config() query.Config {
	ret := _m.Called()

	var r0 query.Config
	if rf, ok := ret.Get(0).(func() query.Config); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(query.Config)
	}

	return r0
}

// Query provides a mock function with given fields: _a0
func (_m *Builder) Query(_a0 ...query.Tx) query.Query {
	_va := make([]interface{}, len(_a0))
	for _i := range _a0 {
		_va[_i] = _a0[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 query.Query
	if rf, ok := ret.Get(0).(func(...query.Tx) query.Query); ok {
		r0 = rf(_a0...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(query.Query)
		}
	}

	return r0
}

// QuoteIdentifier provides a mock function with given fields: _a0
func (_m *Builder) QuoteIdentifier(_a0 string) string {
	ret := _m.Called(_a0)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// SetLogger provides a mock function with given fields: _a0
func (_m *Builder) SetLogger(_a0 logger.Manager) {
	_m.Called(_a0)
}
