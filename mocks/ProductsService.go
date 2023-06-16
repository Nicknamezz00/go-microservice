// Code generated by mockery v2.28.1. DO NOT EDIT.

package mocks

import (
	context "context"

	models "github.com/Nicknamezz00/go-microservice/internal/pkg/models"
	mock "github.com/stretchr/testify/mock"
)

// ProductsService is an autogenerated mock type for the ProductsService type
type ProductsService struct {
	mock.Mock
}

// Get provides a mock function with given fields: c, ID
func (_m *ProductsService) Get(c context.Context, ID uint64) (*models.Product, error) {
	ret := _m.Called(c, ID)

	var r0 *models.Product
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) (*models.Product, error)); ok {
		return rf(c, ID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64) *models.Product); ok {
		r0 = rf(c, ID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Product)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(c, ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewProductsService interface {
	mock.TestingT
	Cleanup(func())
}

// NewProductsService creates a new instance of ProductsService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewProductsService(t mockConstructorTestingTNewProductsService) *ProductsService {
	mock := &ProductsService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}