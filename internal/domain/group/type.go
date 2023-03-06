package group

import (
	"time"

	"github.com/evgeniy-dammer/clean-architecture/internal/domain/group/description"
	"github.com/evgeniy-dammer/clean-architecture/internal/domain/group/name"
	"github.com/google/uuid"
)

type Group struct {
	createdAt    time.Time
	modifiedAt   time.Time
	name         name.Name
	description  description.Description
	contactCount uint64
	id           uuid.UUID
}

func NewWithID(
	id uuid.UUID,
	createdAt time.Time,
	modifiedAt time.Time,
	name name.Name,
	description description.Description,
	contactCount uint64,
) (*Group, error) {
	return &Group{
		id:           id,
		createdAt:    createdAt.UTC(),
		modifiedAt:   modifiedAt.UTC(),
		name:         name,
		description:  description,
		contactCount: contactCount,
	}, nil
}

func New(name name.Name, description description.Description) *Group {
	timeNow := time.Now().UTC()

	return &Group{
		id:           uuid.New(),
		name:         name,
		description:  description,
		createdAt:    timeNow,
		modifiedAt:   timeNow,
		contactCount: 0,
	}
}

func (g Group) ContactCount() uint64 {
	return g.contactCount
}

func (g Group) ID() uuid.UUID {
	return g.id
}

func (g Group) CreatedAt() time.Time {
	return g.createdAt
}

func (g Group) ModifiedAt() time.Time {
	return g.modifiedAt
}

func (g Group) Name() name.Name {
	return g.name
}

func (g Group) Description() description.Description {
	return g.description
}
