// Code generated by mockery v2.20.0. DO NOT EDIT.

package mockStorage

import (
	contact "github.com/evgeniy-dammer/clean-architecture/internal/domain/contact"
	context "github.com/evgeniy-dammer/clean-architecture/pkg/type/context"

	group "github.com/evgeniy-dammer/clean-architecture/internal/domain/group"

	mock "github.com/stretchr/testify/mock"

	queryparameter "github.com/evgeniy-dammer/clean-architecture/pkg/type/queryparameter"

	uuid "github.com/google/uuid"
)

// Storage is an autogenerated mock type for the Storage type
type Storage struct {
	mock.Mock
}

// AddContactToGroup provides a mock function with given fields: ctx, groupID, contactIDs
func (_m *Storage) AddContactToGroup(ctx context.Context, groupID uuid.UUID, contactIDs ...uuid.UUID) error {
	_va := make([]interface{}, len(contactIDs))
	for _i := range contactIDs {
		_va[_i] = contactIDs[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, groupID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, ...uuid.UUID) error); ok {
		r0 = rf(ctx, groupID, contactIDs...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CountContact provides a mock function with given fields: ctx, parameter
func (_m *Storage) CountContact(ctx context.Context, parameter queryparameter.QueryParameter) (uint64, error) {
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

// CountGroup provides a mock function with given fields: ctx, parameter
func (_m *Storage) CountGroup(ctx context.Context, parameter queryparameter.QueryParameter) (uint64, error) {
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
func (_m *Storage) CreateContact(ctx context.Context, contacts ...*contact.Contact) ([]*contact.Contact, error) {
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

// CreateContactIntoGroup provides a mock function with given fields: ctx, groupID, contacts
func (_m *Storage) CreateContactIntoGroup(ctx context.Context, groupID uuid.UUID, contacts ...*contact.Contact) ([]*contact.Contact, error) {
	_va := make([]interface{}, len(contacts))
	for _i := range contacts {
		_va[_i] = contacts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, groupID)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []*contact.Contact
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, ...*contact.Contact) ([]*contact.Contact, error)); ok {
		return rf(ctx, groupID, contacts...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, ...*contact.Contact) []*contact.Contact); ok {
		r0 = rf(ctx, groupID, contacts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*contact.Contact)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID, ...*contact.Contact) error); ok {
		r1 = rf(ctx, groupID, contacts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateGroup provides a mock function with given fields: ctx, _a1
func (_m *Storage) CreateGroup(ctx context.Context, _a1 *group.Group) (*group.Group, error) {
	ret := _m.Called(ctx, _a1)

	var r0 *group.Group
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *group.Group) (*group.Group, error)); ok {
		return rf(ctx, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *group.Group) *group.Group); ok {
		r0 = rf(ctx, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*group.Group)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *group.Group) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteContact provides a mock function with given fields: ctx, contactID
func (_m *Storage) DeleteContact(ctx context.Context, contactID uuid.UUID) error {
	ret := _m.Called(ctx, contactID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) error); ok {
		r0 = rf(ctx, contactID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteContactFromGroup provides a mock function with given fields: ctx, groupID, contactID
func (_m *Storage) DeleteContactFromGroup(ctx context.Context, groupID uuid.UUID, contactID uuid.UUID) error {
	ret := _m.Called(ctx, groupID, contactID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, uuid.UUID) error); ok {
		r0 = rf(ctx, groupID, contactID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteGroup provides a mock function with given fields: ctx, groupID
func (_m *Storage) DeleteGroup(ctx context.Context, groupID uuid.UUID) error {
	ret := _m.Called(ctx, groupID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) error); ok {
		r0 = rf(ctx, groupID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetContactByID provides a mock function with given fields: ctx, contactID
func (_m *Storage) GetContactByID(ctx context.Context, contactID uuid.UUID) (*contact.Contact, error) {
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

// GetGroupByID provides a mock function with given fields: ctx, groupID
func (_m *Storage) GetGroupByID(ctx context.Context, groupID uuid.UUID) (*group.Group, error) {
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

// GetListContact provides a mock function with given fields: ctx, parameter
func (_m *Storage) GetListContact(ctx context.Context, parameter queryparameter.QueryParameter) ([]*contact.Contact, error) {
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

// GetListGroup provides a mock function with given fields: ctx, parameter
func (_m *Storage) GetListGroup(ctx context.Context, parameter queryparameter.QueryParameter) ([]*group.Group, error) {
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

// UpdateContact provides a mock function with given fields: ctx, contactID, updateFn
func (_m *Storage) UpdateContact(ctx context.Context, contactID uuid.UUID, updateFn func(*contact.Contact) (*contact.Contact, error)) (*contact.Contact, error) {
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

// UpdateGroup provides a mock function with given fields: ctx, groupID, updateFn
func (_m *Storage) UpdateGroup(ctx context.Context, groupID uuid.UUID, updateFn func(*group.Group) (*group.Group, error)) (*group.Group, error) {
	ret := _m.Called(ctx, groupID, updateFn)

	var r0 *group.Group
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, func(*group.Group) (*group.Group, error)) (*group.Group, error)); ok {
		return rf(ctx, groupID, updateFn)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, func(*group.Group) (*group.Group, error)) *group.Group); ok {
		r0 = rf(ctx, groupID, updateFn)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*group.Group)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID, func(*group.Group) (*group.Group, error)) error); ok {
		r1 = rf(ctx, groupID, updateFn)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewStorage interface {
	mock.TestingT
	Cleanup(func())
}

// NewStorage creates a new instance of Storage. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewStorage(t mockConstructorTestingTNewStorage) *Storage {
	mock := &Storage{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
