package contact

import (
	"github.com/evgeniy-dammer/clean-architecture/pkg/type/queryparameter"
	"os"
	"testing"
	"time"
	
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
	updateContacts    []*contact.Contact
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
	
	updateContact, _ := contact.NewWithID(
		uuid.New(),
		time.Now().UTC(),
		time.Now().UTC(),
		*phone.New("88002002020"),
		contactEmail,
		*contactName,
		*contactSurname,
		*contactPatronymic,
		*contactAge,
		gender.MALE,
	)
	
	createContacts = append(createContacts, createContact)
	updateContacts = append(updateContacts, updateContact)
	
	os.Exit(m.Run())
}

func initTestUseCaseContact(t *testing.T) {
	assertion := assert.New(t)
	
	storageRepository.On("SetOptions",
		mock.Anything).
		Return(func(options Options) {
			ucDialog.options = options
		})
	
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
	
	storageRepository.On("GetListContact",
		mock.Anything,
		mock.Anything).
		Return(func(ctx context.Context, parameter queryparameter.QueryParameter) []*contact.Contact {
			return createContacts
		}, func(ctx context.Context, parameter queryparameter.QueryParameter) error {
			return nil
		})
	
	storageRepository.On("UpdateContact",
		mock.Anything,
		mock.Anything,
		mock.Anything).
		Return(func(ctx context.Context, ID uuid.UUID, updateFn func(c *contact.Contact) (*contact.Contact, error)) *contact.Contact {
			return updateContacts[0]
		}, func(ctx context.Context, ID uuid.UUID, updateFn func(c *contact.Contact) (*contact.Contact, error)) error {
			return nil
		})
	
	storageRepository.On("DeleteContact",
		mock.Anything,
		mock.AnythingOfType("uuid.UUID")).
		Return(func(ctx context.Context, contactID uuid.UUID) error {
			delete(data, contactID)
			createContacts = createContacts[:len(createContacts)-1]
			
			if _, ok := data[contactID]; ok {
				return usecase.ErrContactNotFound
			}
			
			return nil
		})
	
	storageRepository.On("CountContact",
		mock.Anything,
		mock.Anything,
		mock.Anything).
		Return(func(ctx context.Context, parameter queryparameter.QueryParameter) uint64 {
			return uint64(len(data))
		}, func(ctx context.Context, parameter queryparameter.QueryParameter) error {
			return nil
		})
}

func TestContact(t *testing.T) {
	initTestUseCaseContact(t)
	
	ucDialog = New(storageRepository, Options{})
	assertion := assert.New(t)
	
	t.Run("new usecase", func(t *testing.T) {
		result := New(storageRepository, Options{})
		assertion.Equal(result, ucDialog)
	})
	
	t.Run("setup options", func(t *testing.T) {
		ucDialog.SetOptions(Options{})
		assertion.Equal(ucDialog.options, Options{})
	})
	
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
	
	t.Run("get contacts", func(t *testing.T) {
		ctx := context.Empty()
		
		result, err := ucDialog.GetListContact(ctx, queryparameter.QueryParameter{})
		assertion.NoError(err)
		assertion.Equal(result, createContacts)
	})
	
	t.Run("update contact", func(t *testing.T) {
		ctx := context.Empty()
		
		result, err := ucDialog.UpdateContact(ctx, updateContacts[0])
		assertion.NoError(err)
		assertion.Equal(result, updateContacts[0])
	})
	
	t.Run("delete contact", func(t *testing.T) {
		ctx := context.Empty()
		
		err := ucDialog.DeleteContact(ctx, createContacts[0].ID())
		assertion.NoError(err)
		assertion.Equal(err, nil)
	})
	
	t.Run("count contacts", func(t *testing.T) {
		ctx := context.Empty()
		
		result, err := ucDialog.CountContact(ctx, queryparameter.QueryParameter{})
		assertion.NoError(err)
		assertion.Equal(result, uint64(len(createContacts)))
	})
}
