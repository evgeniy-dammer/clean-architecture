package contact

import (
	"os"
	"testing"
	
	"github.com/evgeniy-dammer/clean-architecture/internal/domain/contact"
	"github.com/evgeniy-dammer/clean-architecture/internal/domain/contact/age"
	"github.com/evgeniy-dammer/clean-architecture/internal/domain/contact/name"
	"github.com/evgeniy-dammer/clean-architecture/internal/domain/contact/patronymic"
	"github.com/evgeniy-dammer/clean-architecture/internal/domain/contact/surname"
	mockStorage "github.com/evgeniy-dammer/clean-architecture/internal/repository/storage/mock"
	"github.com/evgeniy-dammer/clean-architecture/internal/usecase"
	"github.com/evgeniy-dammer/clean-architecture/pkg/type/context"
	"github.com/evgeniy-dammer/clean-architecture/pkg/type/email"
	"github.com/evgeniy-dammer/clean-architecture/pkg/type/gender"
	"github.com/evgeniy-dammer/clean-architecture/pkg/type/phone"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	storageRepository = new(mockStorage.Contact)
	ucDialog          *UseCase
	data              = make(map[uuid.UUID]*contact.Contact)
	createContacts    []*contact.Contact
)

func TestMain(m *testing.M) {
	contactAge, _ := age.New(42)
	contactName, _ := name.New("Иван")
	contactSurname, _ := surname.New("Иванов")
	contactPatronymic, _ := patronymic.New("Иванович")
	contactEmail, _ := email.New("ivanii@gmail.com")
	createContact, _ := contact.New(
		*phone.New("88002002020"),
		contactEmail,
		*contactName,
		*contactSurname,
		*contactPatronymic,
		*contactAge,
		gender.MALE,
	)
	
	createContacts = append(createContacts, createContact)
	
	os.Exit(m.Run())
}

func initTestUseCaseContact(t *testing.T) {
	assertion := assert.New(t)
	storageRepository.On("CreateContact",
		mock.Anything,
		mock.Anything).
		Return(func(ctx context.Context, contacts ...*contact.Contact) []*contact.Contact {
			assertion.Equal(contacts, createContacts)
			
			for _, c := range contacts {
				data[c.ID()] = c
			}
			
			return contacts
		}, func(ctx context.Context, contacts ...*contact.Contact) error {
			return nil
		})
	
	storageRepository.On("GetContactByID",
		mock.Anything,
		mock.AnythingOfType("uuid.UUID")).
		Return(func(ctx context.Context, ID uuid.UUID) *contact.Contact {
			if c, ok := data[ID]; ok {
				return c
			}
			
			return nil
		}, func(ctx context.Context, ID uuid.UUID) error {
			if _, ok := data[ID]; !ok {
				return usecase.ErrContactNotFound
			}
			
			return nil
		})
	
	storageRepository.On("UpdateContact",
		mock.Anything,
		mock.Anything).
		Return(func(ctx context.Context, ID uuid.UUID, updateFn func(c *contact.Contact) (*contact.Contact, error)) *contact.Contact {
			return nil
		}, func(ctx context.Context, ID uuid.UUID, updateFn func(c *contact.Contact) (*contact.Contact, error)) error {
			return nil
		})
}

func TestContact(t *testing.T) {
	initTestUseCaseContact(t)
	
	ucDialog = New(storageRepository, Options{})
	
	assertion := assert.New(t)
	t.Run("create contact", func(t *testing.T) {
		ctx := context.Empty()
		
		result, err := ucDialog.CreateContact(ctx, createContacts...)
		assertion.NoError(err)
		assertion.Equal(result, createContacts)
	})
	
	t.Run("get contact", func(t *testing.T) {
		ctx := context.Empty()
		
		result, err := ucDialog.GetContactByID(ctx, createContacts[0].ID())
		assertion.NoError(err)
		assertion.Equal(result, createContacts[0])
	})
}
