package contact

import (
	"time"

	"github.com/evgeniy-dammer/clean-architecture/pkg/type/email"
	"github.com/evgeniy-dammer/clean-architecture/pkg/type/gender"
)

type ID struct {
	Value string `json:"id" uri:"id" binding:"required,uuid" example:"00000000-0000-0000-0000-000000000000" format:"uuid"`
}

// ResponseContact
// Contact response structure
type ResponseContact struct {
	// Contact ID
	ID string `json:"id" binding:"required,uuid" example:"00000000-0000-0000-0000-000000000000" format:"uuid"`
	// Create date
	CreatedAt time.Time `json:"createdAt"  binding:"required"`
	// Update date
	ModifiedAt time.Time `json:"modifiedAt"  binding:"required"`
	ShortContact
}

// ShortContact
// Short contact structure
type ShortContact struct {
	// Phone
	PhoneNumber string `json:"phoneNumber" binding:"required,max=50" maxLength:"50" example:"78002002020"`
	// Email
	Email email.Email `json:"email" binding:"omitempty,max=250,email" maxLength:"250" example:"example@gmail.com" format:"email" swaggertype:"string"` //nolint:lll
	// Name
	Name string `json:"name" binding:"max=50" maxLength:"50" example:"Иван"`
	// Surname
	Surname string `json:"surname" binding:"max=100" maxLength:"100" example:"Иванов"`
	// Patronymic
	Patronymic string `json:"patronymic" binding:"max=100" maxLength:"100" example:"Иванович"`
	// Gender
	Gender gender.Gender `json:"gender" example:"1" enums:"1,2" swaggertype:"integer"`
	// Age
	Age uint8 `json:"age" binding:"min=0,max=200" minimum:"0" maximum:"200" default:"0" example:"42"`
}

// ListContact
// Contact list structure
type ListContact struct {
	// List of contacts
	List []*ResponseContact `json:"list"`
	// Total contacts in system
	Total uint64 `json:"total" example:"10" default:"0" binding:"min=0" minimum:"0"`
	// Limit of contacts in request
	Limit uint64 `json:"limit"  example:"10" default:"10" binding:"min=0" minimum:"0"`
	// Offset by contacts
	Offset uint64 `json:"offset" example:"20" default:"0" binding:"min=0" minimum:"0"`
}
