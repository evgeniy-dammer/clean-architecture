package postgres

import (
	"github.com/evgeniy-dammer/clean-architecture/internal/domain/contact"
	"github.com/evgeniy-dammer/clean-architecture/internal/domain/contact/age"
	"github.com/evgeniy-dammer/clean-architecture/internal/domain/contact/name"
	"github.com/evgeniy-dammer/clean-architecture/internal/domain/contact/patronymic"
	"github.com/evgeniy-dammer/clean-architecture/internal/domain/contact/surname"
	"github.com/evgeniy-dammer/clean-architecture/internal/repository/storage/postgres/dao"
	"github.com/evgeniy-dammer/clean-architecture/pkg/type/email"
	"github.com/evgeniy-dammer/clean-architecture/pkg/type/gender"
	"github.com/evgeniy-dammer/clean-architecture/pkg/type/phone"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
)

func (r Repository) toCopyFromSource(contacts ...*contact.Contact) pgx.CopyFromSource {
	rows := make([][]interface{}, len(contacts))

	for i, val := range contacts {
		rows[i] = []interface{}{
			val.ID(),
			val.CreatedAt(),
			val.ModifiedAt(),
			val.PhoneNumber().String(),
			val.Email().String(),
			val.Name().String(),
			val.Surname().String(),
			val.Patronymic().String(),
			val.Age(),
			val.Gender(),
		}
	}

	return pgx.CopyFromRows(rows)
}

func (r Repository) toDomainContact(dao *dao.Contact) (*contact.Contact, error) {
	nameObject, err := name.New(dao.Name)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create a new name")
	}

	surnameObject, err := surname.New(dao.Surname)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create a new surname")
	}

	patronymicObject, err := patronymic.New(dao.Patronymic)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create a new patronymic")
	}

	ageObject, err := age.New(dao.Age)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create a new age")
	}

	localEmail, err := email.New(dao.Email)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create a new email")
	}

	result, err := contact.NewWithID(
		dao.ID,
		dao.CreatedAt,
		dao.ModifiedAt,
		*phone.New(dao.PhoneNumber),
		localEmail,
		*nameObject,
		*surnameObject,
		*patronymicObject,
		*ageObject,
		gender.Gender(dao.Gender),
	)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create a new contact with ID")
	}

	return result, nil
}

func (r Repository) toDomainContacts(dao []*dao.Contact) ([]*contact.Contact, error) {
	result := make([]*contact.Contact, len(dao))

	for index, value := range dao {
		c, err := r.toDomainContact(value)
		if err != nil {
			return nil, err
		}

		result[index] = c
	}

	return result, nil
}
