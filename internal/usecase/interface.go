package usecase

import (
	"github.com/evgeniy-dammer/clean-architecture/internal/domain/contact"
	"github.com/evgeniy-dammer/clean-architecture/internal/domain/group"
	"github.com/evgeniy-dammer/clean-architecture/pkg/type/queryparameter"
	"github.com/google/uuid"
)

type Contact interface {
	CreateContact(contacts ...*contact.Contact) ([]*contact.Contact, error)
	UpdateContact(contactUpdate *contact.Contact) (*contact.Contact, error)
	DeleteContact(ID uuid.UUID /*Тут можно передавать фильтр*/) error

	ContactReader
}

type ContactReader interface {
	GetListContact(parameter queryparameter.QueryParameter) ([]*contact.Contact, error)
	GetContactByID(ID uuid.UUID) (response *contact.Contact, err error)
	CountContact(parameter queryparameter.QueryParameter) (uint64, error)
}

type Group interface {
	CreateGroup(groupCreate *group.Group) (*group.Group, error)
	UpdateGroup(groupUpdate *group.Group) (*group.Group, error)
	DeleteGroup(ID uuid.UUID /*Тут можно передавать фильтр*/) error

	GroupReader
	ContactInGroup
}

type GroupReader interface {
	GetListGroup(parameter queryparameter.QueryParameter) ([]*group.Group, error)
	GetGroupByID(ID uuid.UUID) (*group.Group, error)
	CountGroup(parameter queryparameter.QueryParameter) (uint64, error)
}

type ContactInGroup interface {
	CreateContactIntoGroup(groupID uuid.UUID, contacts ...*contact.Contact) ([]*contact.Contact, error)
	AddContactToGroup(groupID uuid.UUID, contactIDs ...uuid.UUID) error
	DeleteContactFromGroup(groupID, contactID uuid.UUID) error
}
