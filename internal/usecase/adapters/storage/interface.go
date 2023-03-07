package storage

import (
	"github.com/evgeniy-dammer/clean-architecture/internal/domain/contact"
	"github.com/evgeniy-dammer/clean-architecture/internal/domain/group"
	"github.com/evgeniy-dammer/clean-architecture/pkg/type/context"
	"github.com/evgeniy-dammer/clean-architecture/pkg/type/queryparameter"
	"github.com/google/uuid"
)

type Storage interface {
	Contact
	Group
}

type Contact interface {
	CreateContact(ctx context.Context, contacts ...*contact.Contact) ([]*contact.Contact, error)
	UpdateContact(ctx context.Context, contactID uuid.UUID, updateFn func(contact *contact.Contact) (*contact.Contact, error)) (*contact.Contact, error) //nolint:lll
	DeleteContact(ctx context.Context, contactID uuid.UUID) error

	ContactReader
}

type ContactReader interface {
	GetListContact(ctx context.Context, parameter queryparameter.QueryParameter) ([]*contact.Contact, error)
	GetContactByID(ctx context.Context, contactID uuid.UUID) (*contact.Contact, error)
	CountContact(ctx context.Context, parameter queryparameter.QueryParameter) (uint64, error)
}

type Group interface {
	CreateGroup(ctx context.Context, group *group.Group) (*group.Group, error)
	UpdateGroup(ctx context.Context, groupID uuid.UUID, updateFn func(group *group.Group) (*group.Group, error)) (*group.Group, error) //nolint:lll
	DeleteGroup(ctx context.Context, groupID uuid.UUID) error

	GroupReader
	ContactInGroup
}

type GroupReader interface {
	GetListGroup(ctx context.Context, parameter queryparameter.QueryParameter) ([]*group.Group, error)
	GetGroupByID(ctx context.Context, groupID uuid.UUID) (*group.Group, error)
	CountGroup(ctx context.Context, parameter queryparameter.QueryParameter) (uint64, error)
}

type ContactInGroup interface {
	CreateContactIntoGroup(ctx context.Context, groupID uuid.UUID, contacts ...*contact.Contact) ([]*contact.Contact, error) //nolint:lll
	AddContactToGroup(ctx context.Context, groupID uuid.UUID, contactIDs ...uuid.UUID) error
	DeleteContactFromGroup(ctx context.Context, groupID, contactID uuid.UUID) error
}
