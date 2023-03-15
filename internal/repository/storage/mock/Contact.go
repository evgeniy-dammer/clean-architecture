// Code generated by mockery v2.20.0. DO NOT EDIT.

package mockStorage

import (
	contact "github.com/evgeniy-dammer/clean-architecture/internal/domain/contact"
	context "github.com/evgeniy-dammer/clean-architecture/pkg/type/context"

	mock "github.com/stretchr/testify/mock"

	queryparameter "github.com/evgeniy-dammer/clean-architecture/pkg/type/queryparameter"

	uuid "github.com/google/uuid"
)

// Contact is an autogenerated mock type for the Contact type
type Contact struct {
	mock.Mock
}

// CountContact provides a mock function with given fields: ctx, parameter
func (_m *Contact) CountContact(ctx context.Context, parameter queryparameter.QueryParameter) (uint64, error) {
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

// CreateContact provides a mock function with given fields: ctx, contacts
func (_m *Contact) CreateContact(ctx context.Context, contacts ...*contact.Contact) ([]*contact.Contact, error) {
	_va := make([]interface{}, len(contacts))
	for _i := range contacts {
		_va[_i] = contacts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []*contact.Contact
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ...*contact.Contact) ([]*contact.Contact, error)); ok {
		return rf(ctx, contacts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ...*contact.Contact) []*contact.Contact); ok {
		r0 = rf(ctx, contacts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*contact.Contact)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ...*contact.Contact) error); ok {
		r1 = rf(ctx, contacts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteContact provides a mock function with given fields: ctx, contactID
func (_m *Contact) DeleteContact(ctx context.Context, contactID uuid.UUID) error {
	ret := _m.Called(ctx, contactID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) error); ok {
		r0 = rf(ctx, contactID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetContactByID provides a mock function with given fields: ctx, contactID
func (_m *Contact) GetContactByID(ctx context.Context, contactID uuid.UUID) (*contact.Contact, error) {
	ret := _m.Called(ctx, contactID)

	var r0 *contact.Contact
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) (*contact.Contact, error)); ok {
		return rf(ctx, contactID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *contact.Contact); ok {
		r0 = rf(ctx, contactID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*contact.Contact)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, contactID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetListContact provides a mock function with given fields: ctx, parameter
func (_m *Contact) GetListContact(ctx context.Context, parameter queryparameter.QueryParameter) ([]*contact.Contact, error) {
	ret := _m.Called(ctx, parameter)

	var r0 []*contact.Contact
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, queryparameter.QueryParameter) ([]*contact.Contact, error)); ok {
		return rf(ctx, parameter)
	}
	if rf, ok := ret.Get(0).(func(context.Context, queryparameter.QueryParameter) []*contact.Contact); ok {
		r0 = rf(ctx, parameter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*contact.Contact)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, queryparameter.QueryParameter) error); ok {
		r1 = rf(ctx, parameter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateContact provides a mock function with given fields: ctx, contactID, updateFn
func (_m *Contact) UpdateContact(ctx context.Context, contactID uuid.UUID, updateFn func(*contact.Contact) (*contact.Contact, error)) (*contact.Contact, error) {
	ret := _m.Called(ctx, contactID, updateFn)

	var r0 *contact.Contact
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, func(*contact.Contact) (*contact.Contact, error)) (*contact.Contact, error)); ok {
		return rf(ctx, contactID, updateFn)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, func(*contact.Contact) (*contact.Contact, error)) *contact.Contact); ok {
		r0 = rf(ctx, contactID, updateFn)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*contact.Contact)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID, func(*contact.Contact) (*contact.Contact, error)) error); ok {
		r1 = rf(ctx, contactID, updateFn)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewContact interface {
	mock.TestingT
	Cleanup(func())
}

// NewContact creates a new instance of Contact. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewContact(t mockConstructorTestingTNewContact) *Contact {
	mock := &Contact{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}