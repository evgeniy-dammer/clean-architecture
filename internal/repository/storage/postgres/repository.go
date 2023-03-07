package postgres

import (
	"fmt"
	"os"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
	"github.com/pressly/goose"
)

type Repository struct {
	options Options
	genSQL  squirrel.StatementBuilderType
	db      *pgxpool.Pool
}

type Options struct {
	DefaultLimit  uint64
	DefaultOffset uint64
	Timeout       time.Duration
}

func New(database *pgxpool.Pool, options Options) (*Repository, error) {
	if err := migrations(database); err != nil {
		return nil, err
	}

	r := &Repository{
		genSQL: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
		db:     database,
	}

	r.SetOptions(options)

	return r, nil
}

func (r *Repository) SetOptions(options Options) {
	if r.options != options {
		r.options = options
	}
}

func migrations(pool *pgxpool.Pool) (err error) {
	database, err := goose.OpenDBWithDriver("postgres", pool.Config().ConnConfig.ConnString())
	if err != nil {
		return errors.Wrap(err, "unable to open database with driver")
	}

	defer func() {
		if errClose := database.Close(); errClose != nil {
			err = errClose

			return
		}
	}()

	goose.SetTableName("contact_version")

	if err = goose.Run("up", database, os.Getenv("MIGRATIONS_DIR")); err != nil {
		return fmt.Errorf("goose %s error : %w", "up", err)
	}

	return
}
