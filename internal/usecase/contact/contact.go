package contact

import (
	"time"
	
	"github.com/evgeniy-dammer/clean-architecture/internal/domain/contact"
	"github.com/evgeniy-dammer/clean-architecture/pkg/type/context"
	"github.com/evgeniy-dammer/clean-architecture/pkg/type/queryparameter"
	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
)

func (uc *UseCase) CreateContact(ctx context.Context, contacts ...*contact.Contact) ([]*contact.Contact, error) {
	span, ctxt := opentracing.StartSpanFromContext(ctx, "CreateContact")
	defer span.Finish()
	
	list, err := uc.adapterStorage.CreateContact(context.New(ctxt), contacts...)
	
	return list, errors.Wrap(err, "create contact use case error")
}

func (uc *UseCase) UpdateContact(ctx context.Context, contactUpdate *contact.Contact) (*contact.Contact, error) {
	span, ctxt := opentracing.StartSpanFromContext(ctx, "UpdateContact")
	defer span.Finish()
	
	cntct, err := uc.adapterStorage.UpdateContact(
		context.New(ctxt),
		contactUpdate.ID(),
		func(oldContact *contact.Contact) (*contact.Contact, error) {
			cnt, err := contact.NewWithID(
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
			
			return cnt, errors.Wrap(err, "unable to create new contact with ID")
		})
	
	return cntct, errors.Wrap(err, "update contact use case error")
}

func (uc *UseCase) DeleteContact(ctx context.Context, contactID uuid.UUID) error {
	span, ctxt := opentracing.StartSpanFromContext(ctx, "DeleteContact")
	defer span.Finish()
	
	err := uc.adapterStorage.DeleteContact(context.New(ctxt), contactID)
	
	return errors.Wrap(err, "delete contact use case error")
}

func (uc *UseCase) GetListContact(ctx context.Context, parameter queryparameter.QueryParameter) ([]*contact.Contact, error) { //nolint:lll
	span, ctxt := opentracing.StartSpanFromContext(ctx, "GetListContact")
	defer span.Finish()
	
	contacts, err := uc.adapterStorage.GetListContact(context.New(ctxt), parameter)
	
	return contacts, errors.Wrap(err, "get contact list use case error")
}

func (uc *UseCase) GetContactByID(ctx context.Context, contactID uuid.UUID) (*contact.Contact, error) {
	span, ctxt := opentracing.StartSpanFromContext(ctx, "GetContactByID")
	defer span.Finish()
	
	response, err := uc.adapterStorage.GetContactByID(context.New(ctxt), contactID)
	
	return response, errors.Wrap(err, "get contact by ID use case error")
}

func (uc *UseCase) CountContact(ctx context.Context, parameter queryparameter.QueryParameter) (uint64, error) {
	span, ctxt := opentracing.StartSpanFromContext(ctx, "CountContact")
	defer span.Finish()
	
	count, err := uc.adapterStorage.CountContact(context.New(ctxt), parameter)
	
	return count, errors.Wrap(err, "count contacts use case error")
}
