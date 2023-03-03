package contact

import (
	"github.com/evgeniy-dammer/clean-architecture/internal/domain/contact"
	"github.com/evgeniy-dammer/clean-architecture/pkg/type/queryparameter"
	"github.com/google/uuid"
)

func (uc *UseCase) Create(contacts ...*contact.Contact) ([]*contact.Contact, error) {
	panic("implement me")
}

func (uc *UseCase) Update(contactUpdate contact.Contact) (*contact.Contact, error) {
	panic("implement me")
}

func (uc *UseCase) Delete(contactID uuid.UUID) error {
	panic("implement me")
}

func (uc *UseCase) List(parameter queryparameter.QueryParameter) ([]*contact.Contact, error) {
	panic("implement me")
}

func (uc *UseCase) ReadByID(contactID uuid.UUID) (response *contact.Contact, err error) {
	panic("implement me")
}

func (uc *UseCase) Count() (uint64, error) {
	panic("implement me")
}
