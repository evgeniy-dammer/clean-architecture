package contact

import (
	"github.com/evgeniy-dammer/clean-architecture/internal/domain/contact"
	"github.com/evgeniy-dammer/clean-architecture/pkg/type/queryparameter"
	"github.com/google/uuid"
)

func (uc *UseCase) CreateContact(contacts ...*contact.Contact) ([]*contact.Contact, error) {
	panic("implement me")
}

func (uc *UseCase) UpdateContact(contactUpdate *contact.Contact) (*contact.Contact, error) {
	panic("implement me")
}

func (uc *UseCase) DeleteContact(contactID uuid.UUID) error {
	panic("implement me")
}

func (uc *UseCase) GetListContact(parameter queryparameter.QueryParameter) ([]*contact.Contact, error) {
	panic("implement me")
}

func (uc *UseCase) GetContactByID(contactID uuid.UUID) (response *contact.Contact, err error) {
	panic("implement me")
}

func (uc *UseCase) CountContact(parameter queryparameter.QueryParameter) (uint64, error) {
	panic("implement me")
}
