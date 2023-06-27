// Code generated by mockery v2.20.0. DO NOT EDIT.

package resource

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	storage "github.com/hashicorp/consul/internal/storage"
	pbresource "github.com/hashicorp/consul/proto-public/pbresource"
)

// MockBackend is an autogenerated mock type for the Backend type
type MockBackend struct {
	mock.Mock
}

// DeleteCAS provides a mock function with given fields: ctx, id, version
func (_m *MockBackend) DeleteCAS(ctx context.Context, id *pbresource.ID, version string) error {
	ret := _m.Called(ctx, id, version)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *pbresource.ID, string) error); ok {
		r0 = rf(ctx, id, version)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// List provides a mock function with given fields: ctx, consistency, resType, tenancy, namePrefix
func (_m *MockBackend) List(ctx context.Context, consistency storage.ReadConsistency, resType storage.UnversionedType, tenancy *pbresource.Tenancy, namePrefix string) ([]*pbresource.Resource, error) {
	ret := _m.Called(ctx, consistency, resType, tenancy, namePrefix)

	var r0 []*pbresource.Resource
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, storage.ReadConsistency, storage.UnversionedType, *pbresource.Tenancy, string) ([]*pbresource.Resource, error)); ok {
		return rf(ctx, consistency, resType, tenancy, namePrefix)
	}
	if rf, ok := ret.Get(0).(func(context.Context, storage.ReadConsistency, storage.UnversionedType, *pbresource.Tenancy, string) []*pbresource.Resource); ok {
		r0 = rf(ctx, consistency, resType, tenancy, namePrefix)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*pbresource.Resource)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, storage.ReadConsistency, storage.UnversionedType, *pbresource.Tenancy, string) error); ok {
		r1 = rf(ctx, consistency, resType, tenancy, namePrefix)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListByOwner provides a mock function with given fields: ctx, id
func (_m *MockBackend) ListByOwner(ctx context.Context, id *pbresource.ID) ([]*pbresource.Resource, error) {
	ret := _m.Called(ctx, id)

	var r0 []*pbresource.Resource
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pbresource.ID) ([]*pbresource.Resource, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pbresource.ID) []*pbresource.Resource); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*pbresource.Resource)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pbresource.ID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Read provides a mock function with given fields: ctx, consistency, id
func (_m *MockBackend) Read(ctx context.Context, consistency storage.ReadConsistency, id *pbresource.ID) (*pbresource.Resource, error) {
	ret := _m.Called(ctx, consistency, id)

	var r0 *pbresource.Resource
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, storage.ReadConsistency, *pbresource.ID) (*pbresource.Resource, error)); ok {
		return rf(ctx, consistency, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, storage.ReadConsistency, *pbresource.ID) *pbresource.Resource); ok {
		r0 = rf(ctx, consistency, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pbresource.Resource)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, storage.ReadConsistency, *pbresource.ID) error); ok {
		r1 = rf(ctx, consistency, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WatchList provides a mock function with given fields: ctx, resType, tenancy, namePrefix
func (_m *MockBackend) WatchList(ctx context.Context, resType storage.UnversionedType, tenancy *pbresource.Tenancy, namePrefix string) (storage.Watch, error) {
	ret := _m.Called(ctx, resType, tenancy, namePrefix)

	var r0 storage.Watch
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, storage.UnversionedType, *pbresource.Tenancy, string) (storage.Watch, error)); ok {
		return rf(ctx, resType, tenancy, namePrefix)
	}
	if rf, ok := ret.Get(0).(func(context.Context, storage.UnversionedType, *pbresource.Tenancy, string) storage.Watch); ok {
		r0 = rf(ctx, resType, tenancy, namePrefix)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(storage.Watch)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, storage.UnversionedType, *pbresource.Tenancy, string) error); ok {
		r1 = rf(ctx, resType, tenancy, namePrefix)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WriteCAS provides a mock function with given fields: ctx, res
func (_m *MockBackend) WriteCAS(ctx context.Context, res *pbresource.Resource) (*pbresource.Resource, error) {
	ret := _m.Called(ctx, res)

	var r0 *pbresource.Resource
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *pbresource.Resource) (*pbresource.Resource, error)); ok {
		return rf(ctx, res)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *pbresource.Resource) *pbresource.Resource); ok {
		r0 = rf(ctx, res)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*pbresource.Resource)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *pbresource.Resource) error); ok {
		r1 = rf(ctx, res)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewMockBackend interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockBackend creates a new instance of MockBackend. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockBackend(t mockConstructorTestingTNewMockBackend) *MockBackend {
	mock := &MockBackend{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
