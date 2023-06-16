// Code generated by mockery v2.28.1. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	grpc "google.golang.org/grpc"
)

// InitServers is an autogenerated mock type for the InitServers type
type InitServers struct {
	mock.Mock
}

// Execute provides a mock function with given fields: s
func (_m *InitServers) Execute(s *grpc.Server) {
	_m.Called(s)
}

type mockConstructorTestingTNewInitServers interface {
	mock.TestingT
	Cleanup(func())
}

// NewInitServers creates a new instance of InitServers. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewInitServers(t mockConstructorTestingTNewInitServers) *InitServers {
	mock := &InitServers{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}