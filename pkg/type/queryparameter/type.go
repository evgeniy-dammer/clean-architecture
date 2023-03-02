package queryparameter

import (
	"github.com/evgeniy-dammer/clean-architecture/pkg/type/pagination"
	"github.com/evgeniy-dammer/clean-architecture/pkg/type/sort"
)

type QueryParameter struct {
	Sorts      sort.Sorts
	Pagination pagination.Pagination
	/*Тут можно добавить фильтр*/
}
