package contact

import (
	"time"

	"github.com/evgeniy-dammer/clean-architecture/internal/domain/contact"
	"github.com/evgeniy-dammer/clean-architecture/pkg/type/context"
	"github.com/evgeniy-dammer/clean-architecture/pkg/type/queryparameter"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func (uc *UseCase) CreateContact(ctx context.Context, contacts ...*contact.Contact) ([]*contact.Contact, error) {
	list, err := uc.adapterStorage.CreateContact(ctx, contacts...)

	return list, errors.Wrap(err, "create contact use case error")
}

func (uc *UseCase) UpdateContact(ctx context.Context, contactUpdate *contact.Contact) (*contact.Contact, error) {
	contact, err := uc.adapterStorage.UpdateContact(
		ctx,
		contactUpdate.ID(),
		func(oldContact *contact.Contact) (*contact.Contact, error) {
			contact, err := contact.NewWithID(
				oldContact.ID(),
				oldContact.CreatedAt(),
				time.Now().UTC(),
				contactUpdate.PhoneNumber(),
				contactUpdate.Email(),
				contactUpdate.Name(),
				contactUpdate.Surname(),
				contactUpdate.Patronymic(),
				contactUpdate.Age(),
				contactUpdate.Gender(),
			)

			return contact, errors.Wrap(err, "unable to create new contact with ID")
		})

	return contact, errors.Wrap(err, "update contact use case error")
}

func (uc *UseCase) DeleteContact(ctx context.Context, contactID uuid.UUID) error {
	err := uc.adapterStorage.DeleteContact(ctx, contactID)

	return errors.Wrap(err, "delete contact use case error")
}

func (uc *UseCase) GetListContact(ctx context.Context, parameter queryparameter.QueryParameter) ([]*contact.Contact, error) {
	contacts, err := uc.adapterStorage.GetListContact(ctx, parameter)

	return contacts, errors.Wrap(err, "get contact list use case error")
}

func (uc *UseCase) GetContactByID(ctx context.Context, contactID uuid.UUID) (*contact.Contact, error) {
	response, err := uc.adapterStorage.GetContactByID(ctx, contactID)

	return response, errors.Wrap(err, "get contact by ID use case error")
}

func (uc *UseCase) CountContact(ctx context.Context, parameter queryparameter.QueryParameter) (uint64, error) {
	count, err := uc.adapterStorage.CountContact(ctx, parameter)

	return count, errors.Wrap(err, "count contacts use case error")
}
