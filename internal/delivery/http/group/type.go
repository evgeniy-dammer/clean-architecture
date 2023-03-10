package group

import "time"

// ResponseGroup
// Response group structure.
type ResponseGroup struct {
	// Group ID
	ID string `json:"id" binding:"required,uuid" example:"00000000-0000-0000-0000-000000000000" format:"uuid"`
	// Create date
	CreatedAt time.Time `json:"createdAt"  binding:"required"`
	// Update date
	ModifiedAt time.Time `json:"modifiedAt"  binding:"required"`
	Group
}

// Group
// Object for group structure.
type Group struct {
	ShortGroup
	// Contacts count in group
	ContactsAmount uint64 `json:"contactsAmount" default:"10" binding:"min=0" minimum:"0"`
}

// ShortGroup
// Short group structure.
type ShortGroup struct {
	// Group name
	Name string `json:"name" binding:"required,max=100" example:"Название группы" maxLength:"100"`
	// Group description
	Description string `json:"description" example:"Описание группы" binding:"max=1000" maxLength:"1000"`
}

// ListGroup
// Object for group list structure.
type ListGroup struct {
	// List of groups
	List []*ResponseGroup `json:"list" binding:"min=0" minimum:"0"`
	// Total groups in system
	Total uint64 `json:"total" example:"10" default:"0" binding:"min=0" minimum:"0"`
	// Limit of groups in request
	Limit uint64 `json:"limit"  example:"10" default:"10" binding:"min=0" minimum:"0"`
	// Offset by groups
	Offset uint64 `json:"offset" example:"20" default:"0" binding:"min=0" minimum:"0"`
}

type ID struct {
	// Group ID
	Value string `json:"id" uri:"id" binding:"required,uuid" example:"00000000-0000-0000-0000-000000000000" format:"uuid"`
}

type ContactID struct {
	// Contact ID
	Value string `json:"id" uri:"contactId" binding:"required,uuid" example:"00000000-0000-0000-0000-000000000000" format:"uuid"` //nolint:lll
}
