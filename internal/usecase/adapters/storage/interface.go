package storage

import (
	"github.com/evgeniy-dammer/clean-architecture/internal/domain/contact"
	"github.com/evgeniy-dammer/clean-architecture/internal/domain/group"
	"github.com/evgeniy-dammer/clean-architecture/pkg/type/queryparameter"
	"github.com/google/uuid"
)

type Storage interface {
	Contact
	Group
}

type Contact interface {
	CreateContact(contacts ...*contact.Contact) ([]*contact.Contact, error)
	UpdateContact(contactID uuid.UUID, updateFn func(contact *contact.Contact) (*contact.Contact, error)) (*contact.Contact, error)
	DeleteContact(contactID uuid.UUID) error

	ContactReader
}

type ContactReader interface {
	GetListContact(parameter queryparameter.QueryParameter) ([]*contact.Contact, error)
	GetContactByID(contactID uuid.UUID) (*contact.Contact, error)
	CountContact(parameter queryparameter.QueryParameter) (uint64, error)
}

type Group interface {
	CreateGroup(group *group.Group) (*group.Group, error)
	UpdateGroup(groupID uuid.UUID, updateFn func(group *group.Group) (*group.Group, error)) (*group.Group, error)
	DeleteGroup(groupID uuid.UUID) error

	GroupReader
	ContactInGroup
}

type GroupReader interface {
	GetListGroup(parameter queryparameter.QueryParameter) ([]*group.Group, error)
	GetGroupByID(groupID uuid.UUID) (*group.Group, error)
	CountGroup(parameter queryparameter.QueryParameter) (uint64, error)
}

type ContactInGroup interface {
	CreateContactIntoGroup(groupID uuid.UUID, contacts ...*contact.Contact) ([]*contact.Contact, error)
	AddContactToGroup(groupID uuid.UUID, contactIDs ...uuid.UUID) error
	DeleteContactFromGroup(groupID, contactID uuid.UUID) error
}
