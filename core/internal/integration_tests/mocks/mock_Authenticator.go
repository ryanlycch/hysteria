// Code generated by mockery v2.43.0. DO NOT EDIT.

package mocks

import (
	net "net"

	mock "github.com/stretchr/testify/mock"
)

// MockAuthenticator is an autogenerated mock type for the Authenticator type
type MockAuthenticator struct {
	mock.Mock
}

type MockAuthenticator_Expecter struct {
	mock *mock.Mock
}

func (_m *MockAuthenticator) EXPECT() *MockAuthenticator_Expecter {
	return &MockAuthenticator_Expecter{mock: &_m.Mock}
}

// Authenticate provides a mock function with given fields: addr, auth, tx
func (_m *MockAuthenticator) Authenticate(addr net.Addr, auth string, tx uint64) (bool, string) {
	ret := _m.Called(addr, auth, tx)

	if len(ret) == 0 {
		panic("no return value specified for Authenticate")
	}

	var r0 bool
	var r1 string
	if rf, ok := ret.Get(0).(func(net.Addr, string, uint64) (bool, string)); ok {
		return rf(addr, auth, tx)
	}
	if rf, ok := ret.Get(0).(func(net.Addr, string, uint64) bool); ok {
		r0 = rf(addr, auth, tx)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(net.Addr, string, uint64) string); ok {
		r1 = rf(addr, auth, tx)
	} else {
		r1 = ret.Get(1).(string)
	}

	return r0, r1
}

// MockAuthenticator_Authenticate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Authenticate'
type MockAuthenticator_Authenticate_Call struct {
	*mock.Call
}

// Authenticate is a helper method to define mock.On call
//   - addr net.Addr
//   - auth string
//   - tx uint64
func (_e *MockAuthenticator_Expecter) Authenticate(addr interface{}, auth interface{}, tx interface{}) *MockAuthenticator_Authenticate_Call {
	return &MockAuthenticator_Authenticate_Call{Call: _e.mock.On("Authenticate", addr, auth, tx)}
}

func (_c *MockAuthenticator_Authenticate_Call) Run(run func(addr net.Addr, auth string, tx uint64)) *MockAuthenticator_Authenticate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(net.Addr), args[1].(string), args[2].(uint64))
	})
	return _c
}

func (_c *MockAuthenticator_Authenticate_Call) Return(ok bool, id string) *MockAuthenticator_Authenticate_Call {
	_c.Call.Return(ok, id)
	return _c
}

func (_c *MockAuthenticator_Authenticate_Call) RunAndReturn(run func(net.Addr, string, uint64) (bool, string)) *MockAuthenticator_Authenticate_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockAuthenticator creates a new instance of MockAuthenticator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockAuthenticator(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockAuthenticator {
	mock := &MockAuthenticator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
