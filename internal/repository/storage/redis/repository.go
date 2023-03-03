package redis

import "github.com/evgeniy-dammer/clean-architecture/internal/usecase/adapters/storage"

type Repository struct {
	options Options
}

type Options struct{}

func New(storage storage.Storage, options Options) *Repository {
	r := &Repository{}

	r.SetOptions(options)

	return r
}

func (r *Repository) SetOptions(options Options) {
	if r.options != options {
		r.options = options
	}
}
