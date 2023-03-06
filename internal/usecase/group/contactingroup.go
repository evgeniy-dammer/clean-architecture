package group

import (
	"github.com/evgeniy-dammer/clean-architecture/internal/domain/contact"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func (uc *UseCase) CreateContactIntoGroup(groupID uuid.UUID, contacts ...*contact.Contact) ([]*contact.Contact, error) {
	list, err := uc.adapterStorage.CreateContactIntoGroup(groupID, contacts...)

	return list, errors.Wrap(err, "create contact in group use case error")
}

func (uc *UseCase) AddContactToGroup(groupID uuid.UUID, contactIDs ...uuid.UUID) error {
	err := uc.adapterStorage.AddContactToGroup(groupID, contactIDs...)

	return errors.Wrap(err, "add contact to group use case error")
}

func (uc *UseCase) DeleteContactFromGroup(groupID, contactID uuid.UUID) error {
	err := uc.adapterStorage.DeleteContactFromGroup(groupID, contactID)

	return errors.Wrap(err, "delete contact from group error")
}
