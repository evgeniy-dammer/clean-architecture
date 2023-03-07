package query

import (
	"github.com/evgeniy-dammer/clean-architecture/pkg/type/sort"
	"github.com/gin-gonic/gin"
)

type Query struct {
	Sorts  sort.Sorts
	Limit  uint64
	Offset uint64
}

type SortOptions struct{}

type Options struct {
	// Тут можно добавить фильтр
	Sorts SortsOptions
}

type SortsOptions map[string]SortOptions // map[front_key]FilterOptions

var (
	keyForSort        = "sort"
	defaultKeyForSort = ""
	keyForLimit       = "limit"
	keyForOffset      = "offset"
)

func ParseQuery(ctx *gin.Context, options Options) (*Query, error) {
	sorts, err := parseSorts(ctx.DefaultQuery(keyForSort, defaultKeyForSort), options.Sorts)
	if err != nil {
		return nil, err
	}

	return &Query{
		Sorts:  sorts,
		Limit:  parseLimit(ctx.Query(keyForLimit)),
		Offset: parseOffset(ctx.Query(keyForOffset)),
	}, nil
}

func ParseSorts(c *gin.Context, options SortsOptions) (sort.Sorts, error) {
	return parseSorts(c.DefaultQuery(keyForSort, defaultKeyForSort), options)
}

func ParseLimit(c *gin.Context) uint64 {
	return parseLimit(c.Query(keyForLimit))
}

func ParseOffset(c *gin.Context) uint64 {
	return parseOffset(c.Query(keyForOffset))
}
