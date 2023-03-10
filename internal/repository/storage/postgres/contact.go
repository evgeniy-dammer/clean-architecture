package postgres

import (
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/evgeniy-dammer/clean-architecture/internal/domain/contact"
	"github.com/evgeniy-dammer/clean-architecture/internal/repository/storage/postgres/dao"
	"github.com/evgeniy-dammer/clean-architecture/pkg/tools/transaction"
	"github.com/evgeniy-dammer/clean-architecture/pkg/type/columncode"
	"github.com/evgeniy-dammer/clean-architecture/pkg/type/context"
	log "github.com/evgeniy-dammer/clean-architecture/pkg/type/logger"
	"github.com/evgeniy-dammer/clean-architecture/pkg/type/queryparameter"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
)

var mappingSortContact = map[columncode.ColumnCode]string{
	"id":          "id",
	"fullName":    "full_name",
	"phoneNumber": "phone_number",
	"name":        "name",
	"surname":     "surname",
	"patronymic":  "patronymic",
	"email":       "email",
	"gender":      "gender",
	"age":         "age",
}

func (r *Repository) CreateContact(ctxg context.Context, contacts ...*contact.Contact) ([]*contact.Contact, error) {
	ctx := ctxg.CopyWithTimeout(r.options.Timeout)
	defer ctx.Cancel()

	span, ctxt := opentracing.StartSpanFromContext(ctxg, "CreateContact")
	defer span.Finish()

	ctx = context.New(ctxt)

	trx, err := r.db.Begin(ctx)
	if err != nil {
		return nil, errors.Wrap(log.ErrorWithContext(ctx, err), "unable to create contact")
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.Finish(ctx, t, err)
	}(ctx, trx)

	response, err := r.createContactTx(ctx, trx, contacts...)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create contact")
	}

	return response, nil
}

func (r *Repository) createContactTx(ctx context.Context, tx pgx.Tx, contacts ...*contact.Contact) ([]*contact.Contact, error) { //nolint:lll
	if len(contacts) == 0 {
		return []*contact.Contact{}, nil
	}

	_, err := tx.CopyFrom(ctx, pgx.Identifier{"clean", "contact"}, dao.CreateColumnContact, r.toCopyFromSource(contacts...)) //nolint:lll
	if err != nil {
		return nil, errors.Wrap(log.ErrorWithContext(ctx, err), "unable to copy")
	}

	return contacts, nil
}

func (r *Repository) UpdateContact(ctxg context.Context, contactID uuid.UUID, updateFn func(c *contact.Contact) (*contact.Contact, error)) (*contact.Contact, error) { //nolint:lll
	ctx := ctxg.CopyWithTimeout(r.options.Timeout)
	defer ctx.Cancel()

	span, ctxt := opentracing.StartSpanFromContext(ctxg, "UpdateContact")
	defer span.Finish()

	ctx = context.New(ctxt)

	trx, err := r.db.Begin(ctx)
	if err != nil {
		return nil, errors.Wrap(log.ErrorWithContext(ctx, err), "unable to begin transaction")
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.Finish(ctx, t, err)
	}(ctx, trx)

	upContact, err := r.oneContactTx(ctx, trx, contactID)
	if err != nil {
		return nil, errors.Wrap(err, "unable to select contact")
	}

	in, err := updateFn(upContact)
	if err != nil {
		return nil, errors.Wrap(err, "unable to prepare update function")
	}

	return r.updateContactTx(ctx, trx, in)
}

func (r *Repository) updateContactTx(ctx context.Context, trx pgx.Tx, input *contact.Contact) (*contact.Contact, error) { //nolint:lll
	builder := r.genSQL.Update("clean.contact").
		Set("email", input.Email().String()).
		Set("phone_number", input.PhoneNumber().String()).
		Set("age", input.Age()).
		Set("gender", input.Gender()).
		Set("modified_at", input.ModifiedAt()).
		Set("name", input.Name().String()).
		Set("surname", input.Surname().String()).
		Set("patronymic", input.Patronymic().String()).
		Where(squirrel.And{
			squirrel.Eq{
				"id":          input.ID(),
				"is_archived": false,
			},
		}).
		Suffix(`RETURNING id, created_at, modified_at, phone_number, email, name, surname, patronymic, age, gender`)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, errors.Wrap(log.ErrorWithContext(ctx, err), "unable to build a query string")
	}

	rows, err := trx.Query(ctx, query, args...)
	if err != nil {
		return nil, errors.Wrap(log.ErrorWithContext(ctx, err), "unable to execute a query")
	}

	var daoContacts []*dao.Contact

	if err = pgxscan.ScanAll(&daoContacts, rows); err != nil {
		return nil, errors.Wrap(log.ErrorWithContext(ctx, err), "unable to scan")
	}

	return r.toDomainContact(daoContacts[0])
}

func (r *Repository) DeleteContact(ctxg context.Context, contactID uuid.UUID) error {
	ctx := ctxg.CopyWithTimeout(r.options.Timeout)
	defer ctx.Cancel()

	span, ctxt := opentracing.StartSpanFromContext(ctxg, "DeleteContact")
	defer span.Finish()

	ctx = context.New(ctxt)

	trx, err := r.db.Begin(ctx)
	if err != nil {
		return errors.Wrap(log.ErrorWithContext(ctx, err), "unable to begin transaction")
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.Finish(ctx, t, err)
	}(ctx, trx)

	if err = r.deleteContactTx(ctx, trx, contactID); err != nil {
		return errors.Wrap(err, "unable to delete contact")
	}

	return nil
}

func (r *Repository) deleteContactTx(ctx context.Context, trx pgx.Tx, contactID uuid.UUID) error {
	builder := r.genSQL.Update("clean.contact").
		Set("is_archived", true).
		Set("modified_at", time.Now().UTC()).
		Where(squirrel.Eq{"is_archived": false, "id": contactID})

	query, args, err := builder.ToSql()
	if err != nil {
		return errors.Wrap(log.ErrorWithContext(ctx, err), "unable to build a query string")
	}

	rows, err := trx.Query(ctx, query, args...)
	if err != nil {
		return errors.Wrap(log.ErrorWithContext(ctx, err), "unable to execute query")
	}

	var daoContacts []*dao.Contact

	if err = pgxscan.ScanAll(&daoContacts, rows); err != nil {
		return errors.Wrap(log.ErrorWithContext(ctx, err), "unable to scan")
	}

	if err = r.updateGroupsContactCountByFilters(ctx, trx, contactID); err != nil {
		return errors.Wrap(err, "unable to update group contact count by filters")
	}

	return nil
}

func (r *Repository) GetListContact(ctxg context.Context, parameter queryparameter.QueryParameter) ([]*contact.Contact, error) { //nolint:lll
	ctx := ctxg.CopyWithTimeout(r.options.Timeout)
	defer ctx.Cancel()

	span, ctxt := opentracing.StartSpanFromContext(ctxg, "GetListContact")
	defer span.Finish()

	ctx = context.New(ctxt)

	trx, err := r.db.Begin(ctx)
	if err != nil {
		return nil, errors.Wrap(log.ErrorWithContext(ctx, err), "unable to begin transaction")
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.Finish(ctx, t, err)
	}(ctx, trx)

	if parameter.Pagination.Limit == 0 {
		parameter.Pagination.Limit = r.options.DefaultLimit
	}

	contacts, err := r.listContactTx(ctx, trx, parameter)
	if err != nil {
		return nil, errors.Wrap(err, "unable to select contact list")
	}

	return contacts, nil
}

func (r *Repository) listContactTx(ctx context.Context, trx pgx.Tx, parameter queryparameter.QueryParameter) ([]*contact.Contact, error) { //nolint:lll
	builder := r.genSQL.Select("id", "created_at", "modified_at", "phone_number", "email", "name", "surname", "patronymic", "age", "gender"). //nolint:lll
																			From("clean.contact")

	builder = builder.Where(squirrel.Eq{"is_archived": false})

	if len(parameter.Sorts) > 0 {
		builder = builder.OrderBy(parameter.Sorts.Parsing(mappingSortContact)...)
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
	if err != nil {
		return nil, errors.Wrap(log.ErrorWithContext(ctx, err), "unable to build a query string")
	}

	rows, err := trx.Query(ctx, query, args...)
	if err != nil {
		return nil, errors.Wrap(log.ErrorWithContext(ctx, err), "unable to execute a query")
	}

	var daoContacts []*dao.Contact

	if err = pgxscan.ScanAll(&daoContacts, rows); err != nil {
		return nil, errors.Wrap(log.ErrorWithContext(ctx, err), "unable to scan")
	}

	return r.toDomainContacts(daoContacts)
}

func (r *Repository) GetContactByID(ctxg context.Context, contactID uuid.UUID) (*contact.Contact, error) {
	ctx := ctxg.CopyWithTimeout(r.options.Timeout)
	defer ctx.Cancel()

	span, ctxt := opentracing.StartSpanFromContext(ctxg, "GetContactByID")
	defer span.Finish()

	ctx = context.New(ctxt)

	trx, err := r.db.Begin(ctx)
	if err != nil {
		return nil, errors.Wrap(log.ErrorWithContext(ctx, err), "unable to begin transaction")
	}

	defer func(ctx context.Context, t pgx.Tx) {
		err = transaction.Finish(ctx, t, err)
	}(ctx, trx)

	response, err := r.oneContactTx(ctx, trx, contactID)
	if err != nil {
		return nil, errors.Wrap(err, "unable to select contact")
	}

	return response, nil
}

func (r *Repository) oneContactTx(ctx context.Context, trx pgx.Tx, contactID uuid.UUID) (*contact.Contact, error) {
	builder := r.genSQL.Select("id", "created_at", "modified_at", "phone_number", "email", "name", "surname", "patronymic", "age", "gender"). //nolint:lll
																			From("clean.contact")

	builder = builder.Where(squirrel.Eq{"is_archived": false, "id": contactID})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, errors.Wrap(log.ErrorWithContext(ctx, err), "unable to build a query string")
	}

	rows, err := trx.Query(ctx, query, args...)
	if err != nil {
		return nil, errors.Wrap(log.ErrorWithContext(ctx, err), "unable to execute query")
	}

	var daoContact []*dao.Contact

	if err = pgxscan.ScanAll(&daoContact, rows); err != nil {
		return nil, errors.Wrap(log.ErrorWithContext(ctx, err), "unable to scan")
	}

	if len(daoContact) == 0 {
		return nil, errors.New("contact not found")
	}

	return r.toDomainContact(daoContact[0])
}

func (r *Repository) CountContact(ctxg context.Context, parameter queryparameter.QueryParameter) (uint64, error) {
	ctx := ctxg.CopyWithTimeout(r.options.Timeout)
	defer ctx.Cancel()

	span, ctxt := opentracing.StartSpanFromContext(ctxg, "CountContact")
	defer span.Finish()

	ctx = context.New(ctxt)

	builder := r.genSQL.Select("COUNT(id)").From("clean.contact")

	builder = builder.Where(squirrel.Eq{"is_archived": false})

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, errors.Wrap(log.ErrorWithContext(ctx, err), "unable to build a query string")
	}

	row := r.db.QueryRow(ctx, query, args...)

	var total uint64

	if err = row.Scan(&total); err != nil {
		return 0, errors.Wrap(log.ErrorWithContext(ctx, err), "unable to scan")
	}

	return total, nil
}
