package group

import (
	"github.com/evgeniy-dammer/clean-architecture/internal/domain/contact"
	"github.com/evgeniy-dammer/clean-architecture/pkg/type/context"
	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
)

func (uc *UseCase) CreateContactIntoGroup(ctx context.Context, groupID uuid.UUID, contacts ...*contact.Contact) ([]*contact.Contact, error) { //nolint:lll
	span, ctxt := opentracing.StartSpanFromContext(ctx, "CreateContactIntoGroup")
	defer span.Finish()

	list, err := uc.adapterStorage.CreateContactIntoGroup(context.New(ctxt), groupID, contacts...)

	return list, errors.Wrap(err, "create contact in group use case error")
}

func (uc *UseCase) AddContactToGroup(ctx context.Context, groupID uuid.UUID, contactIDs ...uuid.UUID) error {
	span, ctxt := opentracing.StartSpanFromContext(ctx, "AddContactToGroup")
	defer span.Finish()

	err := uc.adapterStorage.AddContactToGroup(context.New(ctxt), groupID, contactIDs...)

	return errors.Wrap(err, "add contact to group use case error")
}

func (uc *UseCase) DeleteContactFromGroup(ctx context.Context, groupID, contactID uuid.UUID) error {
	span, ctxt := opentracing.StartSpanFromContext(ctx, "DeleteContactFromGroup")
	defer span.Finish()

	err := uc.adapterStorage.DeleteContactFromGroup(context.New(ctxt), groupID, contactID)

	return errors.Wrap(err, "delete contact from group error")
}
