package group

import (
	"github.com/evgeniy-dammer/clean-architecture/internal/domain/contact"
	"github.com/evgeniy-dammer/clean-architecture/internal/domain/contact/age"
	cname "github.com/evgeniy-dammer/clean-architecture/internal/domain/contact/name"
	"github.com/evgeniy-dammer/clean-architecture/internal/domain/contact/patronymic"
	"github.com/evgeniy-dammer/clean-architecture/internal/domain/contact/surname"
	"github.com/evgeniy-dammer/clean-architecture/internal/domain/group"
	"github.com/evgeniy-dammer/clean-architecture/internal/domain/group/description"
	"github.com/evgeniy-dammer/clean-architecture/internal/domain/group/name"
	mockStorage "github.com/evgeniy-dammer/clean-architecture/internal/repository/storage/mock"
	"github.com/evgeniy-dammer/clean-architecture/internal/usecase"
	"github.com/evgeniy-dammer/clean-architecture/pkg/type/context"
	"github.com/evgeniy-dammer/clean-architecture/pkg/type/email"
	"github.com/evgeniy-dammer/clean-architecture/pkg/type/gender"
	"github.com/evgeniy-dammer/clean-architecture/pkg/type/phone"
	"github.com/evgeniy-dammer/clean-architecture/pkg/type/queryparameter"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"os"
	"testing"
	"time"
)

var (
	storageRepository = new(mockStorage.Group)
	ucDialog          *UseCase
	data              = make(map[uuid.UUID]*group.Group)
	createGroups      []*group.Group
	updateGroups      []*group.Group
	
	storageRepositoryC = new(mockStorage.Contact)
	dataC              = make(map[uuid.UUID]*contact.Contact)
	createContacts     []*contact.Contact
	updateContacts     []*contact.Contact
	
	dataGC = make(map[uuid.UUID][]*contact.Contact)
)

func TestMain(m *testing.M) {
	groupName, _ := name.New("Name")
	groupDescription, _ := description.New("Description")
	
	createGroup := group.New(
		groupName,
		groupDescription,
	)
	
	updateGroup, _ := group.NewWithID(
		uuid.New(),
		time.Now().UTC(),
		time.Now().UTC(),
		groupName,
		groupDescription,
		0,
	)
	
	contactAge, _ := age.New(42)
	contactName, _ := cname.New("Иван")
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
	
	createGroups = append(createGroups, createGroup)
	updateGroups = append(updateGroups, updateGroup)
	
	createContacts = append(createContacts, createContact)
	updateContacts = append(updateContacts, updateContact)
	
	os.Exit(m.Run())
}

func initTestUseCaseGroup(t *testing.T) {
	assertion := assert.New(t)
	
	storageRepository.On("SetOptions",
		mock.Anything).
		Return(func(options Options) {
			ucDialog.options = options
		})
	
	storageRepository.On("CreateGroup",
		mock.Anything,
		mock.Anything).
		Return(func(ctx context.Context, group *group.Group) *group.Group {
			assertion.Equal(group, createGroups[0])
			
			data[group.ID()] = group
			
			return group
		}, func(ctx context.Context, group *group.Group) error {
			return nil
		})
	
	storageRepository.On("GetGroupByID",
		mock.Anything,
		mock.AnythingOfType("uuid.UUID")).
		Return(func(ctx context.Context, ID uuid.UUID) *group.Group {
			if c, ok := data[ID]; ok {
				return c
			}
			
			return nil
		}, func(ctx context.Context, ID uuid.UUID) error {
			if _, ok := data[ID]; !ok {
				return usecase.ErrGroupNotFound
			}
			
			return nil
		})
	
	storageRepository.On("GetListGroup",
		mock.Anything,
		mock.Anything).
		Return(func(ctx context.Context, parameter queryparameter.QueryParameter) []*group.Group {
			return createGroups
		}, func(ctx context.Context, parameter queryparameter.QueryParameter) error {
			return nil
		})
	
	storageRepository.On("UpdateGroup",
		mock.Anything,
		mock.Anything,
		mock.Anything).
		Return(func(ctx context.Context, ID uuid.UUID, updateFn func(c *group.Group) (*group.Group, error)) *group.Group {
			return updateGroups[0]
		}, func(ctx context.Context, ID uuid.UUID, updateFn func(c *group.Group) (*group.Group, error)) error {
			return nil
		})
	
	storageRepository.On("CreateContactIntoGroup",
		mock.Anything,
		mock.AnythingOfType("uuid.UUID"),
		mock.Anything).
		Return(func(ctx context.Context, groupID uuid.UUID, contacts ...*contact.Contact) []*contact.Contact {
			for _, c := range contacts {
				dataGC[groupID] = append(dataGC[groupID], c)
			}
			
			return contacts
		}, func(ctx context.Context, groupID uuid.UUID, contacts ...*contact.Contact) error {
			return nil
		})
	
	storageRepository.On("AddContactToGroup",
		mock.Anything,
		mock.AnythingOfType("uuid.UUID"),
		mock.Anything).
		Return(func(ctx context.Context, groupID uuid.UUID, contactIDs ...uuid.UUID) error {
			for _, c := range contactIDs {
				dataGC[groupID] = append(dataGC[groupID], dataC[c])
			}
			
			return nil
		})
	
	storageRepository.On("DeleteContactFromGroup",
		mock.Anything,
		mock.AnythingOfType("uuid.UUID"),
		mock.AnythingOfType("uuid.UUID")).
		Return(func(ctx context.Context, groupID, contactID uuid.UUID) error {
			for k, v := range dataGC[groupID] {
				if v.ID() == contactID {
					dataGC[groupID] = append(dataGC[groupID][:k], dataGC[groupID][k+1:]...)
				} else {
					return usecase.ErrContactNotFound
				}
			}
			
			return nil
		})
	
	storageRepository.On("DeleteGroup",
		mock.Anything,
		mock.AnythingOfType("uuid.UUID")).
		Return(func(ctx context.Context, groupID uuid.UUID) error {
			delete(data, groupID)
			createGroups = createGroups[:len(createGroups)-1]
			
			if _, ok := data[groupID]; ok {
				return usecase.ErrGroupNotFound
			}
			
			return nil
		})
	
	storageRepository.On("CountGroup",
		mock.Anything,
		mock.Anything,
		mock.Anything).
		Return(func(ctx context.Context, parameter queryparameter.QueryParameter) uint64 {
			return uint64(len(data))
		}, func(ctx context.Context, parameter queryparameter.QueryParameter) error {
			return nil
		})
	
}

func TestGroup(t *testing.T) {
	initTestUseCaseGroup(t)
	
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
	
	t.Run("create group", func(t *testing.T) {
		ctx := context.Empty()
		
		result, err := ucDialog.CreateGroup(ctx, createGroups[0])
		assertion.NoError(err)
		assertion.Equal(result, createGroups[0])
	})
	
	t.Run("get group", func(t *testing.T) {
		ctx := context.Empty()
		
		result, err := ucDialog.GetGroupByID(ctx, createGroups[0].ID())
		assertion.NoError(err)
		assertion.Equal(result, createGroups[0])
	})
	
	t.Run("get groups", func(t *testing.T) {
		ctx := context.Empty()
		
		result, err := ucDialog.GetListGroup(ctx, queryparameter.QueryParameter{})
		assertion.NoError(err)
		assertion.Equal(result, createGroups)
	})
	
	t.Run("update group", func(t *testing.T) {
		ctx := context.Empty()
		
		result, err := ucDialog.UpdateGroup(ctx, updateGroups[0])
		assertion.NoError(err)
		assertion.Equal(result, updateGroups[0])
	})
	
	t.Run("create contact into group", func(t *testing.T) {
		ctx := context.Empty()
		
		result, err := ucDialog.CreateContactIntoGroup(ctx, createGroups[0].ID(), createContacts...)
		assertion.NoError(err)
		assertion.Equal(result, createContacts)
	})
	
	t.Run("delete contact from group", func(t *testing.T) {
		ctx := context.Empty()
		
		err := ucDialog.DeleteContactFromGroup(ctx, createGroups[0].ID(), createContacts[0].ID())
		assertion.NoError(err)
		assertion.Equal(err, nil)
	})
	
	t.Run("add contact to group", func(t *testing.T) {
		ctx := context.Empty()
		
		err := ucDialog.AddContactToGroup(ctx, createGroups[0].ID(), createContacts[0].ID())
		assertion.NoError(err)
		assertion.Equal(err, nil)
	})
	
	t.Run("delete group", func(t *testing.T) {
		ctx := context.Empty()
		
		err := ucDialog.DeleteGroup(ctx, createGroups[0].ID())
		assertion.NoError(err)
		assertion.Equal(err, nil)
	})
	
	t.Run("count groups", func(t *testing.T) {
		ctx := context.Empty()
		
		result, err := ucDialog.CountGroup(ctx, queryparameter.QueryParameter{})
		assertion.NoError(err)
		assertion.Equal(result, uint64(len(createGroups)))
	})
	
}
