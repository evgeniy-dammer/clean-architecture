package group

import (
	"github.com/evgeniy-dammer/clean-architecture/internal/domain/group"
	"github.com/evgeniy-dammer/clean-architecture/pkg/type/queryparameter"
	"github.com/google/uuid"
)

func (uc *UseCase) Create(groupCreate *group.Group) (*group.Group, error) {
	panic("implement me")
}

func (uc *UseCase) Update(groupUpdate *group.Group) (*group.Group, error) {
	panic("implement me")
}

func (uc *UseCase) Delete(groupID uuid.UUID) error {
	panic("implement me")
}

func (uc *UseCase) List(parameter queryparameter.QueryParameter) ([]*group.Group, error) {
	panic("implement me")
}

func (uc *UseCase) ReadByID(groupID uuid.UUID) (*group.Group, error) {
	panic("implement me")
}

func (uc *UseCase) Count() (uint64, error) {
	panic("implement me")
}
