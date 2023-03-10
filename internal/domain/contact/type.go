package contact

import (
	"fmt"
	"time"

	"github.com/evgeniy-dammer/clean-architecture/internal/domain/contact/age"
	"github.com/evgeniy-dammer/clean-architecture/internal/domain/contact/name"
	"github.com/evgeniy-dammer/clean-architecture/internal/domain/contact/patronymic"
	"github.com/evgeniy-dammer/clean-architecture/internal/domain/contact/surname"
	"github.com/evgeniy-dammer/clean-architecture/pkg/type/email"
	"github.com/evgeniy-dammer/clean-architecture/pkg/type/gender"
	"github.com/evgeniy-dammer/clean-architecture/pkg/type/phone"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

var ErrPhoneNumberRequired = errors.New("phone number is required")

type Contact struct {
	createdAt  time.Time
	modifiedAt time.Time
	phone      phone.Phone
	email      email.Email
	name       name.Name
	surname    surname.Surname
	patronymic patronymic.Patronymic
	id         uuid.UUID
	age        age.Age
	gender     gender.Gender
}

func NewWithID(
	id uuid.UUID,
	createdAt time.Time,
	modifiedAt time.Time,
	phone phone.Phone,
	email email.Email,
	name name.Name,
	surname surname.Surname,
	patronymic patronymic.Patronymic,
	age age.Age,
	gender gender.Gender,
) (*Contact, error) {
	if phone.IsEmpty() {
		return nil, ErrPhoneNumberRequired
	}

	if id == uuid.Nil {
		id = uuid.New()
	}

	return &Contact{
		id:         id,
		createdAt:  createdAt.UTC(),
		modifiedAt: modifiedAt.UTC(),
		phone:      phone,
		email:      email,
		name:       name,
		surname:    surname,
		patronymic: patronymic,
		age:        age,
		gender:     gender,
	}, nil
}

func New(
	phone phone.Phone,
	email email.Email,
	name name.Name,
	surname surname.Surname,
	patronymic patronymic.Patronymic,
	age age.Age,
	gender gender.Gender,
) (*Contact, error) {
	if phone.IsEmpty() {
		return nil, ErrPhoneNumberRequired
	}

	timeNow := time.Now().UTC()

	return &Contact{
		id:         uuid.New(),
		createdAt:  timeNow,
		modifiedAt: timeNow,
		phone:      phone,
		email:      email,
		name:       name,
		surname:    surname,
		patronymic: patronymic,
		age:        age,
		gender:     gender,
	}, nil
}

func (c *Contact) ID() uuid.UUID {
	return c.id
}

func (c *Contact) CreatedAt() time.Time {
	return c.createdAt
}

func (c *Contact) ModifiedAt() time.Time {
	return c.modifiedAt
}

func (c *Contact) Email() email.Email {
	return c.email
}

func (c *Contact) PhoneNumber() phone.Phone {
	return c.phone
}

func (c *Contact) Name() name.Name {
	return c.name
}

func (c *Contact) Surname() surname.Surname {
	return c.surname
}

func (c *Contact) Patronymic() patronymic.Patronymic {
	return c.patronymic
}

func (c *Contact) FullName() string {
	return fmt.Sprintf("%s %s %s", c.surname, c.name, c.patronymic)
}

func (c *Contact) Age() age.Age {
	return c.age
}

func (c *Contact) Gender() gender.Gender {
	return c.gender
}

func (c *Contact) Equal(contact Contact) bool {
	return c.id == contact.id
}
