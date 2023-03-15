// Code generated by mockery v2.20.0. DO NOT EDIT.

package mockStorage

import (
	group "github.com/evgeniy-dammer/clean-architecture/internal/domain/group"
	context "github.com/evgeniy-dammer/clean-architecture/pkg/type/context"

	mock "github.com/stretchr/testify/mock"

	queryparameter "github.com/evgeniy-dammer/clean-architecture/pkg/type/queryparameter"

	uuid "github.com/google/uuid"
)

// GroupReader is an autogenerated mock type for the GroupReader type
type GroupReader struct {
	mock.Mock
}

// CountGroup provides a mock function with given fields: ctx, parameter
func (_m *GroupReader) CountGroup(ctx context.Context, parameter queryparameter.QueryParameter) (uint64, error) {
	ret := _m.Called(ctx, parameter)

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, queryparameter.QueryParameter) (uint64, error)); ok {
		return rf(ctx, parameter)
	}
	if rf, ok := ret.Get(0).(func(context.Context, queryparameter.QueryParameter) uint64); ok {
		r0 = rf(ctx, parameter)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, queryparameter.QueryParameter) error); ok {
		r1 = rf(ctx, parameter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetGroupByID provides a mock function with given fields: ctx, groupID
func (_m *GroupReader) GetGroupByID(ctx context.Context, groupID uuid.UUID) (*group.Group, error) {
	ret := _m.Called(ctx, groupID)

	var r0 *group.Group
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (*group.Group, error)); ok {
		return rf(ctx, groupID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *group.Group); ok {
		r0 = rf(ctx, groupID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*group.Group)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, groupID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetListGroup provides a mock function with given fields: ctx, parameter
func (_m *GroupReader) GetListGroup(ctx context.Context, parameter queryparameter.QueryParameter) ([]*group.Group, error) {
	ret := _m.Called(ctx, parameter)

	var r0 []*group.Group
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, queryparameter.QueryParameter) ([]*group.Group, error)); ok {
		return rf(ctx, parameter)
	}
	if rf, ok := ret.Get(0).(func(context.Context, queryparameter.QueryParameter) []*group.Group); ok {
		r0 = rf(ctx, parameter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*group.Group)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, queryparameter.QueryParameter) error); ok {
		r1 = rf(ctx, parameter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewGroupReader interface {
	mock.TestingT
	Cleanup(func())
}

// NewGroupReader creates a new instance of GroupReader. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewGroupReader(t mockConstructorTestingTNewGroupReader) *GroupReader {
	mock := &GroupReader{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
