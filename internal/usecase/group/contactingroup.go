package group

import (
	"github.com/evgeniy-dammer/clean-architecture/internal/domain/contact"
	"github.com/google/uuid"
)

func (uc *UseCase) CreateContactIntoGroup(groupID uuid.UUID, contacts ...*contact.Contact) ([]*contact.Contact, error) {
	panic("implement me")
}

func (uc *UseCase) AddContactToGroup(groupID, contactID uuid.UUID) error {
	panic("implement me")
}

func (uc *UseCase) DeleteContactFromGroup(groupID, contactID uuid.UUID) error {
	panic("implement me")
}
