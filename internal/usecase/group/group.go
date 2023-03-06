package group

import (
	"time"

	"github.com/evgeniy-dammer/clean-architecture/internal/domain/group"
	"github.com/evgeniy-dammer/clean-architecture/pkg/type/queryparameter"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func (uc *UseCase) CreateGroup(groupCreate *group.Group) (*group.Group, error) {
	group, err := uc.adapterStorage.CreateGroup(groupCreate)

	return group, errors.Wrap(err, "create group use case error")
}

func (uc *UseCase) UpdateGroup(groupUpdate *group.Group) (*group.Group, error) {
	group, err := uc.adapterStorage.UpdateGroup(
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

func (uc *UseCase) DeleteGroup(groupID uuid.UUID) error {
	err := uc.adapterStorage.DeleteGroup(groupID)

	return errors.Wrap(err, "delete group use case error")
}

func (uc *UseCase) GetListGroup(parameter queryparameter.QueryParameter) ([]*group.Group, error) {
	groups, err := uc.adapterStorage.GetListGroup(parameter)

	return groups, errors.Wrap(err, "get group list use case error")
}

func (uc *UseCase) GetGroupByID(groupID uuid.UUID) (*group.Group, error) {
	group, err := uc.adapterStorage.GetGroupByID(groupID)

	return group, errors.Wrap(err, "get group by ID use case error")
}

func (uc *UseCase) CountGroup(parameter queryparameter.QueryParameter) (uint64, error) {
	count, err := uc.adapterStorage.CountGroup(parameter)

	return count, errors.Wrap(err, "count group use case error")
}
