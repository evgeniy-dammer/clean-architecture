package group

import (
	"time"

	"github.com/evgeniy-dammer/clean-architecture/internal/domain/group"
	"github.com/evgeniy-dammer/clean-architecture/pkg/type/context"
	"github.com/evgeniy-dammer/clean-architecture/pkg/type/queryparameter"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func (uc *UseCase) CreateGroup(ctx context.Context, groupCreate *group.Group) (*group.Group, error) {
	group, err := uc.adapterStorage.CreateGroup(ctx, groupCreate)

	return group, errors.Wrap(err, "create group use case error")
}

func (uc *UseCase) UpdateGroup(ctx context.Context, groupUpdate *group.Group) (*group.Group, error) {
	group, err := uc.adapterStorage.UpdateGroup(
		ctx,
		groupUpdate.ID(),
		func(oldGroup *group.Group) (*group.Group, error) {
			group, err := group.NewWithID(
				oldGroup.ID(),
				oldGroup.CreatedAt(),
				time.Now().UTC(),
				groupUpdate.Name(),
				groupUpdate.Description(),
				groupUpdate.ContactCount(),
			)

			return group, errors.Wrap(err, "unable to create new group with ID")
		})

	return group, errors.Wrap(err, "update group use case error")
}

func (uc *UseCase) DeleteGroup(ctx context.Context, groupID uuid.UUID) error {
	err := uc.adapterStorage.DeleteGroup(ctx, groupID)

	return errors.Wrap(err, "delete group use case error")
}

func (uc *UseCase) GetListGroup(ctx context.Context, parameter queryparameter.QueryParameter) ([]*group.Group, error) {
	groups, err := uc.adapterStorage.GetListGroup(ctx, parameter)

	return groups, errors.Wrap(err, "get group list use case error")
}

func (uc *UseCase) GetGroupByID(ctx context.Context, groupID uuid.UUID) (*group.Group, error) {
	group, err := uc.adapterStorage.GetGroupByID(ctx, groupID)

	return group, errors.Wrap(err, "get group by ID use case error")
}

func (uc *UseCase) CountGroup(ctx context.Context, parameter queryparameter.QueryParameter) (uint64, error) {
	count, err := uc.adapterStorage.CountGroup(ctx, parameter)

	return count, errors.Wrap(err, "count group use case error")
}
