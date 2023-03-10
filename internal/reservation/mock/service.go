// Code generated by mockery v2.20.0. DO NOT EDIT.

package mock

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	domain "github.com/1nkh3art1/goods-reservation-service/internal/reservation/domain"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// GoodsAmount provides a mock function with given fields: ctx, warehouseID
func (_m *Service) GoodsAmount(ctx context.Context, warehouseID domain.WarehouseID) (*domain.Responce, error) {
	ret := _m.Called(ctx, warehouseID)

	var r0 *domain.Responce
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.WarehouseID) (*domain.Responce, error)); ok {
		return rf(ctx, warehouseID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.WarehouseID) *domain.Responce); ok {
		r0 = rf(ctx, warehouseID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Responce)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.WarehouseID) error); ok {
		r1 = rf(ctx, warehouseID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReleaseGoods provides a mock function with given fields: ctx, req
func (_m *Service) ReleaseGoods(ctx context.Context, req *domain.Request) (*domain.Responce, error) {
	ret := _m.Called(ctx, req)

	var r0 *domain.Responce
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Request) (*domain.Responce, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Request) *domain.Responce); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Responce)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *domain.Request) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReserveGoods provides a mock function with given fields: ctx, req
func (_m *Service) ReserveGoods(ctx context.Context, req *domain.Request) (*domain.Responce, error) {
	ret := _m.Called(ctx, req)

	var r0 *domain.Responce
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Request) (*domain.Responce, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Request) *domain.Responce); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Responce)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *domain.Request) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WarehouseAvailability provides a mock function with given fields: ctx, warehouseID
func (_m *Service) WarehouseAvailability(ctx context.Context, warehouseID domain.WarehouseID) (bool, error) {
	ret := _m.Called(ctx, warehouseID)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.WarehouseID) (bool, error)); ok {
		return rf(ctx, warehouseID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.WarehouseID) bool); ok {
		r0 = rf(ctx, warehouseID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.WarehouseID) error); ok {
		r1 = rf(ctx, warehouseID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewService interface {
	mock.TestingT
	Cleanup(func())
}

// NewService creates a new instance of Service. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewService(t mockConstructorTestingTNewService) *Service {
	mock := &Service{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
