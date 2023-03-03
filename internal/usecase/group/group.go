package group

import (
	"github.com/evgeniy-dammer/clean-architecture/internal/domain/group"
	"github.com/evgeniy-dammer/clean-architecture/pkg/type/queryparameter"
	"github.com/google/uuid"
)

func (uc *UseCase) CreateGroup(groupCreate *group.Group) (*group.Group, error) {
	panic("implement me")
}

func (uc *UseCase) UpdateGroup(groupUpdate *group.Group) (*group.Group, error) {
	panic("implement me")
}

func (uc *UseCase) DeleteGroup(groupID uuid.UUID) error {
	panic("implement me")
}

func (uc *UseCase) GetListGroup(parameter queryparameter.QueryParameter) ([]*group.Group, error) {
	panic("implement me")
}

func (uc *UseCase) GetGroupByID(groupID uuid.UUID) (*group.Group, error) {
	panic("implement me")
}

func (uc *UseCase) CountGroup(parameter queryparameter.QueryParameter) (uint64, error) {
	panic("implement me")
}
