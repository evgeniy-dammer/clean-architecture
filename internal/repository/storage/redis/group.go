package redis

import (
	"github.com/evgeniy-dammer/clean-architecture/internal/domain/group"
	"github.com/evgeniy-dammer/clean-architecture/pkg/type/context"
	"github.com/evgeniy-dammer/clean-architecture/pkg/type/queryparameter"
	"github.com/google/uuid"
)

func (r *Repository) CreateGroup(ctx context.Context, group *group.Group) (*group.Group, error) {
	panic("implement me")
}

func (r *Repository) UpdateGroup(ctx context.Context, groupID uuid.UUID, updateFn func(group *group.Group) (*group.Group, error)) (*group.Group, error) { //nolint:lll
	panic("implement me")
}

func (r *Repository) DeleteGroup(ctx context.Context, groupID uuid.UUID) error {
	panic("implement me")
}

func (r *Repository) GetListGroup(ctx context.Context, parameter queryparameter.QueryParameter) ([]*group.Group, error) { //nolint:lll
	panic("implement me")
}

func (r *Repository) GetGroupByID(ctx context.Context, groupID uuid.UUID) (*group.Group, error) {
	panic("implement me")
}

func (r *Repository) CountGroup(ctx context.Context) (uint64, error) {
	panic("implement me")
}
