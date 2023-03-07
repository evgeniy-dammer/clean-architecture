package transaction

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/pkg/errors"
)

func Finish(ctx context.Context, trx pgx.Tx, err error) error {
	if err != nil {
		if rollbackErr := trx.Rollback(ctx); rollbackErr != nil {
			return errors.Wrap(rollbackErr, "transaction rollback error")
		}

		return err
	}

	if commitErr := trx.Commit(ctx); commitErr != nil {
		return errors.Wrap(commitErr, "transaction commit error")
	}

	return nil
}
