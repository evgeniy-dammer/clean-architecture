package sort

import "github.com/evgeniy-dammer/clean-architecture/pkg/type/columnCode"

const (
	DirectionAsc  Direction = "ASC"
	DirectionDesc Direction = "DESC"
)

type Sort struct {
	Key columnCode.ColumnCode
	Direction
}

type Direction string

func (d Direction) String() string {
	return string(d)
}

func (s Sort) Parsing(mapping map[columnCode.ColumnCode]string) string {
	column, ok := mapping[s.Key]
	if !ok {
		return ""
	}

	return column + " " + s.Direction.String()
}

type Sorts []*Sort

func (s Sorts) Parsing(mapping map[columnCode.ColumnCode]string) []string {
	result := make([]string, len(mapping))

	for i, sort := range s {
		result[i] = sort.Parsing(mapping)
	}

	return result
}
