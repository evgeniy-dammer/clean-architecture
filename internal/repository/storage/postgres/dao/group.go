package dao

import (
	"time"

	"github.com/evgeniy-dammer/clean-architecture/internal/domain/group"
	"github.com/evgeniy-dammer/clean-architecture/internal/domain/group/description"
	"github.com/evgeniy-dammer/clean-architecture/internal/domain/group/name"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type Group struct {
	ID           uuid.UUID `db:"id"`
	Name         string    `db:"name"`
	Description  string    `db:"description"`
	CreatedAt    time.Time `db:"created_at"`
	ModifiedAt   time.Time `db:"modified_at"`
	ContactCount uint64    `db:"contact_count"`
	IsArchived   bool      `db:"is_archived"`
}

func (g *Group) ToDomainGroup() (*group.Group, error) {
	gName, err := name.New(g.Name)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create new name")
	}

	gDescription, err := description.New(g.Description)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create new description")
	}

	grp, err := group.NewWithID(g.ID, g.CreatedAt, g.ModifiedAt, gName, gDescription, g.ContactCount)

	return grp, errors.Wrap(err, "unable to create new group with ID")
}
