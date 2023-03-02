package usecase

import (
	"github.com/evgeniy-dammer/clean-architecture/internal/domain/contact"
	"github.com/evgeniy-dammer/clean-architecture/internal/domain/group"
	"github.com/evgeniy-dammer/clean-architecture/pkg/type/queryparameter"
	"github.com/google/uuid"
)

type Contact interface {
	Create(contacts ...*contact.Contact) ([]*contact.Contact, error)
	Update(contact *contact.Contact) (*contact.Contact, error)
	Delete(contactID uuid.UUID) error

	ContactReader
}

type ContactReader interface {
	GetList(parameter queryparameter.QueryParameter) ([]*contact.Contact, error)
	GetByID(contactID uuid.UUID) (*contact.Contact, error)
	Count(parameter queryparameter.QueryParameter) (uint64, error)
}

type Group interface {
	Create(group *group.Group) (*group.Group, error)
	Update(group *group.Group) (*group.Group, error)
	Delete(groupID uuid.UUID) error

	GroupReader
	ContactInGroup
}

type GroupReader interface {
	GetList(parameter queryparameter.QueryParameter) ([]*group.Group, error)
	GetByID(groupID uuid.UUID) (*group.Group, error)
	Count(parameter queryparameter.QueryParameter) (uint64, error)
}

type ContactInGroup interface {
	CreateContactInGroup(groupID uuid.UUID, contacts ...*contact.Contact) ([]*contact.Contact, error)
	AddContactToGroup(groupID, contactID uuid.UUID) error
	DeleteContactFromGroup(groupID, contactID uuid.UUID) error
}
