package postgres

import (
	"context"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/evgeniy-dammer/clean-architecture/internal/domain/contact"
	"github.com/evgeniy-dammer/clean-architecture/internal/repository/storage/postgres/dao"
	"github.com/evgeniy-dammer/clean-architecture/pkg/tools/transaction"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
)

func (r *Repository) CreateContactIntoGroup(groupID uuid.UUID, contacts ...*contact.Contact) ([]*contact.Contact, error) { //nolint:lll
	ctx := context.Background()

	trx, err := r.db.Begin(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "unable to begin transaction")
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.Finish(ctx, t, err)
	}(ctx, trx)

	response, err := r.createContactTx(ctx, trx, contacts...)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create contact")
	}

	contactIDs := make([]uuid.UUID, len(response))

	for i, c := range response {
		contactIDs[i] = c.ID()
	}

	if err = r.fillGroupTx(ctx, trx, groupID, contactIDs...); err != nil {
		return nil, errors.Wrap(err, "unable to fill group")
	}

	return response, nil
}

func (r *Repository) DeleteContactFromGroup(groupID, contactID uuid.UUID) error {
	ctx := context.Background()

	trx, err := r.db.Begin(ctx)
	if err != nil {
		return errors.Wrap(err, "unable to begin transaction")
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.Finish(ctx, t, err)
	}(ctx, trx)

	query, args, err := r.genSQL.
		Delete("clean.contact_in_group").
		Where(squirrel.Eq{"contact_id": contactID, "group_id": groupID}).
		ToSql()
	if err != nil {
		return errors.Wrap(err, "unable to build a query string")
	}

	_, err = trx.Exec(ctx, query, args...)
	if err != nil {
		return errors.Wrap(err, "unable to execute query")
	}

	if err = r.updateGroupContactCount(ctx, trx, groupID); err != nil {
		return errors.Wrap(err, "unable to update group contact count")
	}

	return nil
}

func (r *Repository) AddContactToGroup(groupID uuid.UUID, contactIDs ...uuid.UUID) error {
	ctx := context.Background()

	trx, err := r.db.Begin(ctx)
	if err != nil {
		return errors.Wrap(err, "unable to begin transaction")
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.Finish(ctx, t, err)
	}(ctx, trx)

	if err = r.fillGroupTx(ctx, trx, groupID, contactIDs...); err != nil {
		return errors.Wrap(err, "unable to fill group")
	}

	return nil
}

func (r *Repository) fillGroupTx(ctx context.Context, trx pgx.Tx, groupID uuid.UUID, contactIDs ...uuid.UUID) error {
	_, mapExist, err := r.checkExistContactInGroup(ctx, trx, groupID, contactIDs...)
	if err != nil {
		return errors.Wrap(err, "unable to check if contact exists in group")
	}

	for i := 0; i < len(contactIDs); {
		contactID := contactIDs[i]
		if exist := mapExist[contactID]; exist {
			contactIDs[i] = contactIDs[len(contactIDs)-1]
			contactIDs = contactIDs[:len(contactIDs)-1]

			continue
		}
		i++
	}

	if len(contactIDs) == 0 {
		return nil
	}

	rows := make([][]interface{}, 0, len(contactIDs))

	timeNow := time.Now().UTC()

	for _, contactID := range contactIDs {
		rows = append(rows, []interface{}{
			timeNow,
			timeNow,
			groupID,
			contactID,
		})
	}

	_, err = trx.CopyFrom(
		ctx,
		pgx.Identifier{"clean", "contact_in_group"},
		dao.CreateColumnContactInGroup,
		pgx.CopyFromRows(rows),
	)

	if err != nil {
		return errors.Wrap(err, "unable to copy")
	}

	if err = r.updateGroupContactCount(ctx, trx, groupID); err != nil {
		return errors.Wrap(err, "unable to update group contact count")
	}

	return nil
}

// checkExistContactInGroup
// return listExist -- list existing contactID.
// return mapExist -- mapping contact ID how exist or not exist.
func (r *Repository) checkExistContactInGroup(ctx context.Context, trx pgx.Tx, groupID uuid.UUID, contactIDs ...uuid.UUID) (listExist []uuid.UUID, mapExist map[uuid.UUID]bool, err error) {
	listExist = make([]uuid.UUID, 0)
	mapExist = make(map[uuid.UUID]bool)

	if len(contactIDs) == 0 {
		return listExist, mapExist, nil
	}

	query, args, err := r.genSQL.
		Select("contact_id").
		From("clean.contact_in_group").
		Where(squirrel.Eq{"contact_id": contactIDs, "group_id": groupID}).ToSql()
	if err != nil {
		return nil, nil, errors.Wrap(err, "unable to build a query string")
	}

	rows, err := trx.Query(ctx, query, args...)
	if err != nil {
		return nil, nil, errors.Wrap(err, "unable to execute query")
	}

	for rows.Next() {
		contactID := uuid.UUID{}

		if err = rows.Scan(&contactID); err != nil {
			return nil, nil, errors.Wrap(err, "unable to scan ID")
		}

		listExist = append(listExist, contactID)
		mapExist[contactID] = true
	}

	for _, contactID := range contactIDs {
		if _, ok := mapExist[contactID]; !ok {
			mapExist[contactID] = false
		}
	}

	if err = rows.Err(); err != nil {
		return nil, nil, errors.Wrap(err, "errors occurred in rows")
	}

	return listExist, mapExist, nil
}
