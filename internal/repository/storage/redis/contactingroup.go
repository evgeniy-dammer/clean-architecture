package redis

import (
	"github.com/evgeniy-dammer/clean-architecture/internal/domain/contact"
	"github.com/evgeniy-dammer/clean-architecture/pkg/type/context"
	"github.com/google/uuid"
)

func (r *Repository) CreateContactIntoGroup(ctx context.Context, groupID uuid.UUID, in ...*contact.Contact) ([]*contact.Contact, error) { //nolint:lll
	panic("implement me")
}

func (r *Repository) DeleteContactFromGroup(ctx context.Context, groupID, contactID uuid.UUID) error {
	panic("implement me")
}

func (r *Repository) AddContactToGroup(ctx context.Context, groupID uuid.UUID, contactIDs ...uuid.UUID) error {
	panic("implement me")
}
