// Code generated by mockery v2.4.0-beta. DO NOT EDIT.

package mocks

import (
	http "net/http"

	router "github.com/patrickascher/gofer/router"
	mock "github.com/stretchr/testify/mock"
)

// Provider is an autogenerated mock type for the Provider type
type Provider struct {
	mock.Mock
}

// AddPublicDir provides a mock function with given fields: url, path
func (_m *Provider) AddPublicDir(url string, path string) error {
	ret := _m.Called(url, path)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(url, path)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AddPublicFile provides a mock function with given fields: url, path
func (_m *Provider) AddPublicFile(url string, path string) error {
	ret := _m.Called(url, path)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(url, path)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AddRoute provides a mock function with given fields: _a0
func (_m *Provider) AddRoute(_a0 router.Route) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(router.Route) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// HTTPHandler provides a mock function with given fields:
func (_m *Provider) HTTPHandler() http.Handler {
	ret := _m.Called()

	var r0 http.Handler
	if rf, ok := ret.Get(0).(func() http.Handler); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.Handler)
		}
	}

	return r0
}

// SetNotFound provides a mock function with given fields: _a0
func (_m *Provider) SetNotFound(_a0 http.Handler) {
	_m.Called(_a0)
}
