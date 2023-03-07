package postgres

import (
	"context"
	"database/sql"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/evgeniy-dammer/clean-architecture/internal/domain/group"
	"github.com/evgeniy-dammer/clean-architecture/internal/repository/storage/postgres/dao"
	"github.com/evgeniy-dammer/clean-architecture/pkg/tools/converter"
	"github.com/evgeniy-dammer/clean-architecture/pkg/tools/transaction"
	"github.com/evgeniy-dammer/clean-architecture/pkg/type/columncode"
	"github.com/evgeniy-dammer/clean-architecture/pkg/type/queryparameter"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
)

var mappingSortGroup = map[columncode.ColumnCode]string{
	"id":          "id",
	"name":        "name",
	"description": "description",
}

func (r *Repository) CreateGroup(group *group.Group) (*group.Group, error) {
	query, args, err := r.genSQL.Insert("clean.group").
		Columns("id", "name", "description", "created_at", "modified_at").
		Values(group.ID(), group.Name().Value(), group.Description().Value(), group.CreatedAt(), group.ModifiedAt()).
		ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "unable to build a query string")
	}

	ctx := context.Background()

	if _, err = r.db.Exec(ctx, query, args...); err != nil {
		return nil, errors.Wrap(err, "unable to execute query")
	}

	return group, nil
}

func (r *Repository) UpdateGroup(groupID uuid.UUID, updateFn func(group *group.Group) (*group.Group, error)) (*group.Group, error) { //nolint:lll
	ctx := context.Background()

	trx, err := r.db.Begin(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "unable to begin transaction")
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.Finish(ctx, t, err)
	}(ctx, trx)

	upGroup, err := r.oneGroupTx(ctx, trx, groupID)
	if err != nil {
		return nil, errors.Wrap(err, "unable to select group")
	}

	groupForUpdate, err := updateFn(upGroup)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create group for update")
	}

	query, args, err := r.genSQL.Update("clean.group").
		Set("name", groupForUpdate.Name().Value()).
		Set("description", groupForUpdate.Description().Value()).
		Set("modified_at", groupForUpdate.ModifiedAt()).
		Where(squirrel.And{
			squirrel.Eq{
				"id":          groupID,
				"is_archived": false,
			},
		}).
		Suffix(`RETURNING id, name, description, created_at, modified_at`).
		ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "unable to build a query")
	}

	rows, err := trx.Query(ctx, query, args...)
	if err != nil {
		return nil, errors.Wrap(err, "unable to execute query")
	}

	var daoGroup []*dao.Group

	if err = pgxscan.ScanAll(&daoGroup, rows); err != nil {
		return nil, errors.Wrap(err, "unable to scan")
	}

	return groupForUpdate, nil
}

func (r *Repository) DeleteGroup(groupID uuid.UUID) error {
	ctx := context.Background()

	trx, err := r.db.Begin(ctx)
	if err != nil {
		return errors.Wrap(err, "unable to begin transaction")
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.Finish(ctx, t, err)
	}(ctx, trx)

	if err = r.deleteGroupTx(ctx, trx, groupID); err != nil {
		return errors.Wrap(err, "unable to delete group")
	}

	return nil
}

func (r *Repository) deleteGroupTx(ctx context.Context, trx pgx.Tx, groupID uuid.UUID) error {
	query, args, err := r.genSQL.Update("clean.group").
		Set("is_archived", true).
		Set("modified_at", time.Now().UTC()).
		Where(squirrel.Eq{
			"id":          groupID,
			"is_archived": false,
		}).ToSql()
	if err != nil {
		return errors.Wrap(err, "unable to build a query string")
	}

	if _, errEx := trx.Exec(ctx, query, args...); errEx != nil {
		return errors.Wrap(err, "unable to execute query")
	}

	if err = r.clearGroupTx(ctx, trx, groupID); err != nil {
		return errors.Wrap(err, "unable to clear group")
	}

	return nil
}

func (r *Repository) clearGroupTx(ctx context.Context, trx pgx.Tx, groupID uuid.UUID) error {
	query, args, err := r.genSQL.
		Delete("clean.contact_in_group").
		Where(squirrel.Eq{"group_id": groupID}).
		ToSql()
	if err != nil {
		return errors.Wrap(err, "unable to build a query string")
	}

	if _, err = trx.Exec(ctx, query, args...); err != nil {
		return errors.Wrap(err, "unable to execute query")
	}

	if err = r.updateGroupContactCount(ctx, trx, groupID); err != nil {
		return errors.Wrap(err, "unable to update group contact count")
	}

	return nil
}

func (r *Repository) GetListGroup(parameter queryparameter.QueryParameter) ([]*group.Group, error) {
	ctx := context.Background()

	trx, err := r.db.Begin(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "unable to begin transaction")
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.Finish(ctx, t, err)
	}(ctx, trx)

	response, err := r.listGroupTx(ctx, trx, parameter)
	if err != nil {
		return nil, errors.Wrap(err, "unable to select group list")
	}

	return response, nil
}

func (r *Repository) listGroupTx(ctx context.Context, trx pgx.Tx, parameter queryparameter.QueryParameter) ([]*group.Group, error) { //nolint:lll
	var result []*group.Group

	builder := r.genSQL.Select("id", "name", "description", "created_at", "modified_at", "contact_count", "is_archived").
		From("clean.group")

	builder = builder.Where(squirrel.Eq{"is_archived": false})

	if len(parameter.Sorts) > 0 {
		builder = builder.OrderBy(parameter.Sorts.Parsing(mappingSortGroup)...)
	} else {
		builder = builder.OrderBy("created_at DESC")
	}

	if parameter.Pagination.Limit > 0 {
		builder = builder.Limit(parameter.Pagination.Limit)
	}

	if parameter.Pagination.Offset > 0 {
		builder = builder.Offset(parameter.Pagination.Offset)
	}

	query, args, err := builder.ToSql()

	rows, err := trx.Query(ctx, query, args...)
	if err != nil {
		return nil, errors.Wrap(err, "unable to execute query")
	}

	var groups []*dao.Group

	if err = pgxscan.ScanAll(&groups, rows); err != nil {
		return nil, errors.Wrap(err, "unable to scan")
	}

	for _, g := range groups {
		domainGroup, err := g.ToDomainGroup()
		if err != nil {
			return nil, errors.Wrap(err, "unable to create group")
		}

		result = append(result, domainGroup)
	}

	return result, nil
}

func (r *Repository) GetGroupByID(groupID uuid.UUID) (*group.Group, error) {
	ctx := context.Background()

	trx, err := r.db.Begin(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "unable to begin transaction")
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.Finish(ctx, t, err)
	}(ctx, trx)

	response, err := r.oneGroupTx(ctx, trx, groupID)
	if err != nil {
		return nil, errors.Wrap(err, "unable to select group")
	}

	return response, nil
}

func (r *Repository) oneGroupTx(ctx context.Context, trx pgx.Tx, groupID uuid.UUID) (*group.Group, error) {
	builder := r.genSQL.Select("id", "name", "description", "created_at", "modified_at", "contact_count", "is_archived").
		From("clean.group")

	builder = builder.Where(squirrel.Eq{"is_archived": false, "id": groupID})

	query, args, err := builder.ToSql()

	rows, err := trx.Query(ctx, query, args...)
	if err != nil {
		return nil, errors.Wrap(err, "unable to execute query")
	}

	var daoGroup []*dao.Group

	if err = pgxscan.ScanAll(&daoGroup, rows); err != nil {
		return nil, errors.Wrap(err, "unable to scan")
	}

	if len(daoGroup) == 0 {
		return nil, errors.New("group not found")
	}

	grp, err := daoGroup[0].ToDomainGroup()

	return grp, errors.Wrap(err, "unable to create new group")
}

func (r *Repository) CountGroup(parameter queryparameter.QueryParameter) (uint64, error) {
	ctx := context.Background()

	builder := r.genSQL.Select("COUNT(id)").From("clean.group")

	builder = builder.Where(squirrel.Eq{"is_archived": false})

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, errors.Wrap(err, "unable to build a query string")
	}

	row := r.db.QueryRow(ctx, query, args...)

	var total uint64

	if err = row.Scan(&total); err != nil {
		return 0, errors.Wrap(err, "unable to scan")
	}

	return total, nil
}

func (r *Repository) updateGroupsContactCountByFilters(ctx context.Context, trx pgx.Tx, groupID uuid.UUID) error {
	builder := r.genSQL.Select("contact_in_group.group_id").
		From("clean.contact_in_group").
		InnerJoin("clean.contact ON contact_in_group.contact_id = contact.id").
		GroupBy("contact_in_group.group_id")

	builder = builder.Where(squirrel.Eq{"contact_in_group.contact_id": groupID})

	query, args, err := builder.ToSql()
	if err != nil {
		return errors.Wrap(err, "unable to build a query string")
	}

	rows, err := trx.Query(ctx, query, args...)
	if err != nil {
		return errors.Wrap(err, "unable to execute query")
	}

	var groupIDs []uuid.UUID

	for rows.Next() {
		var groupID sql.NullString

		if err = rows.Scan(&groupID); err != nil {
			return errors.Wrap(err, "unable to scan")
		}

		groupIDs = append(groupIDs, converter.StringToUUID(groupID.String))
	}

	for _, groupID := range groupIDs {
		if err = r.updateGroupContactCount(ctx, trx, groupID); err != nil {
			return errors.Wrap(err, "unable to update group contact count")
		}
	}

	if err = rows.Err(); err != nil {
		return errors.Wrap(err, "errors occurred in rows")
	}

	return nil
}

func (r *Repository) updateGroupContactCount(ctx context.Context, trx pgx.Tx, groupID uuid.UUID) error {
	subSelect := r.genSQL.Select("count(contact_in_group.id)").
		From("clean.contact_in_group").
		InnerJoin("clean.contact ON contact_in_group.contact_id = contact.id").
		Where(squirrel.Eq{"group_id": groupID, "is_archived": false})

	query, _, err := r.genSQL.
		Update("clean.group").
		Set("contact_count", subSelect).
		Where(squirrel.Eq{"id": groupID}).
		ToSql()
	if err != nil {
		return errors.Wrap(err, "unable to build a query string")
	}

	args := []interface{}{groupID, false}

	if _, err = trx.Exec(ctx, query, args...); err != nil {
		return errors.Wrap(err, "unable to execute query")
	}

	return nil
}
