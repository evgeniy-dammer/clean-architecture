package redis

import (
	"github.com/evgeniy-dammer/clean-architecture/internal/domain/group"
	"github.com/evgeniy-dammer/clean-architecture/pkg/type/queryparameter"
	"github.com/google/uuid"
)

func (r *Repository) CreateGroup(group *group.Group) (*group.Group, error) {
	panic("implement me")
}

func (r *Repository) UpdateGroup(groupID uuid.UUID, updateFn func(group *group.Group) (*group.Group, error)) (*group.Group, error) {
	panic("implement me")
}

func (r *Repository) DeleteGroup(groupID uuid.UUID) error {
	panic("implement me")
}

func (r *Repository) ListGroup(parameter queryparameter.QueryParameter) ([]*group.Group, error) {
	panic("implement me")
}

func (r *Repository) ReadGroupByID(groupID uuid.UUID) (*group.Group, error) {
	panic("implement me")
}

func (r *Repository) CountGroup() (uint64, error) {
	panic("implement me")
}
