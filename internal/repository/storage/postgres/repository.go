package postgres

import (
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Repository struct {
	options Options
	genSQL  squirrel.StatementBuilderType
	db      *pgxpool.Pool
}

type Options struct{}

func New(db *pgxpool.Pool, options Options) *Repository {
	r := &Repository{
		genSQL: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
		db:     db,
	}

	r.SetOptions(options)

	return r
}

func (r *Repository) SetOptions(options Options) {
	if r.options != options {
		r.options = options
	}
}
