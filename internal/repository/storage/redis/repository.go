package redis

import (
	"github.com/go-redis/cache/v8"
)

type Repository struct {
	options Options
	cache   *cache.Cache
}

type Options struct{}

func New(cache *cache.Cache, options Options) *Repository {
	r := &Repository{
		cache: cache,
	}

	r.SetOptions(options)

	return r
}

func (r *Repository) SetOptions(options Options) {
	if r.options != options {
		r.options = options
	}
}
