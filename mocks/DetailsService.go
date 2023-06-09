// Code generated by mockery v2.28.1. DO NOT EDIT.

package mocks

import (
	models "github.com/Nicknamezz00/go-microservice/internal/pkg/models"
	mock "github.com/stretchr/testify/mock"
)

// DetailsService is an autogenerated mock type for the DetailsService type
type DetailsService struct {
	mock.Mock
}

// Get provides a mock function with given fields: ID
func (_m *DetailsService) Get(ID uint64) (*models.Detail, error) {
	ret := _m.Called(ID)

	var r0 *models.Detail
	var r1 error
	if rf, ok := ret.Get(0).(func(uint64) (*models.Detail, error)); ok {
		return rf(ID)
	}
	if rf, ok := ret.Get(0).(func(uint64) *models.Detail); ok {
		r0 = rf(ID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Detail)
		}
	}

	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewDetailsService interface {
	mock.TestingT
	Cleanup(func())
}

// NewDetailsService creates a new instance of DetailsService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewDetailsService(t mockConstructorTestingTNewDetailsService) *DetailsService {
	mock := &DetailsService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
