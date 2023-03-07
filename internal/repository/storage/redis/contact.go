package redis

import (
	"github.com/evgeniy-dammer/clean-architecture/internal/domain/contact"
	"github.com/evgeniy-dammer/clean-architecture/pkg/type/context"
	"github.com/evgeniy-dammer/clean-architecture/pkg/type/queryparameter"
	"github.com/google/uuid"
)

func (r *Repository) CreateContact(ctx context.Context, contacts ...*contact.Contact) ([]*contact.Contact, error) {
	panic("implement me")
}

func (r *Repository) UpdateContact(ctx context.Context, contactID uuid.UUID, updateFn func(c *contact.Contact) (*contact.Contact, error)) (*contact.Contact, error) { //nolint:lll
	panic("implement me")
}

func (r *Repository) DeleteContact(ctx context.Context, contactID uuid.UUID) error {
	panic("implement me")
}

func (r *Repository) GetListContact(ctx context.Context, parameter queryparameter.QueryParameter) ([]*contact.Contact, error) { //nolint:lll
	panic("implement me")
}

func (r *Repository) GetContactByID(ctx context.Context, contactID uuid.UUID) (*contact.Contact, error) {
	panic("implement me")
}

func (r *Repository) CountContact(ctx context.Context) (uint64, error) {
	panic("implement me")
}
