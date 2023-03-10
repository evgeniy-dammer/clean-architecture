package group

import (
	"time"

	"github.com/evgeniy-dammer/clean-architecture/internal/domain/group"
	"github.com/evgeniy-dammer/clean-architecture/pkg/type/context"
	"github.com/evgeniy-dammer/clean-architecture/pkg/type/queryparameter"
	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
)

func (uc *UseCase) CreateGroup(ctx context.Context, groupCreate *group.Group) (*group.Group, error) {
	span, ctxt := opentracing.StartSpanFromContext(ctx, "CreateGroup")
	defer span.Finish()

	grp, err := uc.adapterStorage.CreateGroup(context.New(ctxt), groupCreate)

	return grp, errors.Wrap(err, "create group use case error")
}

func (uc *UseCase) UpdateGroup(ctx context.Context, groupUpdate *group.Group) (*group.Group, error) {
	span, ctxt := opentracing.StartSpanFromContext(ctx, "UpdateGroup")
	defer span.Finish()

	grp, err := uc.adapterStorage.UpdateGroup(
		context.New(ctxt),
		groupUpdate.ID(),
		func(oldGroup *group.Group) (*group.Group, error) {
			newGrp, err := group.NewWithID(
				oldGroup.ID(),
				oldGroup.CreatedAt(),
				time.Now().UTC(),
				groupUpdate.Name(),
				groupUpdate.Description(),
				groupUpdate.ContactCount(),
			)

			return newGrp, errors.Wrap(err, "unable to create new group with ID")
		})

	return grp, errors.Wrap(err, "update group use case error")
}

func (uc *UseCase) DeleteGroup(ctx context.Context, groupID uuid.UUID) error {
	span, ctxt := opentracing.StartSpanFromContext(ctx, "DeleteGroup")
	defer span.Finish()

	err := uc.adapterStorage.DeleteGroup(context.New(ctxt), groupID)

	return errors.Wrap(err, "delete group use case error")
}

func (uc *UseCase) GetListGroup(ctx context.Context, parameter queryparameter.QueryParameter) ([]*group.Group, error) {
	span, ctxt := opentracing.StartSpanFromContext(ctx, "GetListGroup")
	defer span.Finish()

	groups, err := uc.adapterStorage.GetListGroup(context.New(ctxt), parameter)

	return groups, errors.Wrap(err, "get group list use case error")
}

func (uc *UseCase) GetGroupByID(ctx context.Context, groupID uuid.UUID) (*group.Group, error) {
	span, ctxt := opentracing.StartSpanFromContext(ctx, "GetGroupByID")
	defer span.Finish()

	grp, err := uc.adapterStorage.GetGroupByID(context.New(ctxt), groupID)

	return grp, errors.Wrap(err, "get group by ID use case error")
}

func (uc *UseCase) CountGroup(ctx context.Context, parameter queryparameter.QueryParameter) (uint64, error) {
	span, ctxt := opentracing.StartSpanFromContext(ctx, "CountGroup")
	defer span.Finish()

	count, err := uc.adapterStorage.CountGroup(context.New(ctxt), parameter)

	return count, errors.Wrap(err, "count group use case error")
}
