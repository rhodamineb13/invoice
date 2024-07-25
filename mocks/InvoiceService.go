// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	context "context"
	dto "invoice/common/dto"

	mock "github.com/stretchr/testify/mock"
)

// InvoiceService is an autogenerated mock type for the InvoiceService type
type InvoiceService struct {
	mock.Mock
}

// AddInvoice provides a mock function with given fields: _a0, _a1
func (_m *InvoiceService) AddInvoice(_a0 context.Context, _a1 *dto.InvoiceDetailDTO) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for AddInvoice")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *dto.InvoiceDetailDTO) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EditInvoice provides a mock function with given fields: _a0, _a1
func (_m *InvoiceService) EditInvoice(_a0 context.Context, _a1 *dto.InvoiceDetailDTO) error {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for EditInvoice")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *dto.InvoiceDetailDTO) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InvoiceByID provides a mock function with given fields: _a0, _a1
func (_m *InvoiceService) InvoiceByID(_a0 context.Context, _a1 int) (*dto.InvoiceDetailDTO, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for InvoiceByID")
	}

	var r0 *dto.InvoiceDetailDTO
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) (*dto.InvoiceDetailDTO, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) *dto.InvoiceDetailDTO); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.InvoiceDetailDTO)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InvoiceIndex provides a mock function with given fields: _a0
func (_m *InvoiceService) InvoiceIndex(_a0 context.Context) ([]dto.InvoiceListsDTO, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for InvoiceIndex")
	}

	var r0 []dto.InvoiceListsDTO
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]dto.InvoiceListsDTO, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []dto.InvoiceListsDTO); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]dto.InvoiceListsDTO)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewInvoiceService creates a new instance of InvoiceService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewInvoiceService(t interface {
	mock.TestingT
	Cleanup(func())
}) *InvoiceService {
	mock := &InvoiceService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
