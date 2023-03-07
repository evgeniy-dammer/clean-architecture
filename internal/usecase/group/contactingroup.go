package group

import (
	"github.com/evgeniy-dammer/clean-architecture/internal/domain/contact"
	"github.com/evgeniy-dammer/clean-architecture/pkg/type/context"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func (uc *UseCase) CreateContactIntoGroup(ctx context.Context, groupID uuid.UUID, contacts ...*contact.Contact) ([]*contact.Contact, error) {
	list, err := uc.adapterStorage.CreateContactIntoGroup(ctx, groupID, contacts...)

	return list, errors.Wrap(err, "create contact in group use case error")
}

func (uc *UseCase) AddContactToGroup(ctx context.Context, groupID uuid.UUID, contactIDs ...uuid.UUID) error {
	err := uc.adapterStorage.AddContactToGroup(ctx, groupID, contactIDs...)

	return errors.Wrap(err, "add contact to group use case error")
}

func (uc *UseCase) DeleteContactFromGroup(ctx context.Context, groupID, contactID uuid.UUID) error {
	err := uc.adapterStorage.DeleteContactFromGroup(ctx, groupID, contactID)

	return errors.Wrap(err, "delete contact from group error")
}
