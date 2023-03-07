package query

import (
	"strconv"
	"strings"

	"github.com/evgeniy-dammer/clean-architecture/pkg/type/columncode"
	"github.com/evgeniy-dammer/clean-architecture/pkg/type/sort"
	"github.com/pkg/errors"
)

var (
	fieldSeparationCharacter        = ","
	sortTypeCharacters              = []string{"-", "+"}
	defaultValueForLimit     uint64 = 10
	maxValueForLimit         uint64 = 100
)

func parseSorts(strQuery string, options SortsOptions) (sort.Sorts, error) {
	result := make(sort.Sorts, 0)

	if len(strQuery) == 0 {
		return result, nil
	}

	// Получаем массив из значений через запятую "-name,+full_name,phone" => ["-name", "+full_name", "phone"]
	for _, field := range strings.Split(strQuery, fieldSeparationCharacter) {
		if len(field) < 2 {
			continue
		}

		name := field
		direction := sort.DirectionAsc

		if strings.HasPrefix(field, sortTypeCharacters[1]) {
			direction = sort.DirectionAsc
			name = field[len(sortTypeCharacters[1]):]
		}

		if strings.HasPrefix(field, sortTypeCharacters[0]) {
			direction = sort.DirectionDesc
			name = field[len(sortTypeCharacters[0]):]
		}

		key, err := columncode.New(name)
		if err != nil {
			return nil, errors.Wrap(err, "unable to create a new name")
		}

		if _, ok := options[key.String()]; !ok {
			continue
		}

		result = append(result, &sort.Sort{
			Key:       key,
			Direction: direction,
		})
	}

	return result, nil
}

func parseLimit(strLimit string) uint64 {
	limit, err := strconv.ParseUint(strLimit, 10, 64)

	if err != nil || limit == 0 {
		return defaultValueForLimit
	}

	if limit > maxValueForLimit {
		return maxValueForLimit
	}

	return limit
}

func parseOffset(strOffset string) uint64 {
	offset, err := strconv.ParseUint(strOffset, 10, 64)
	if err != nil {
		return 0
	}

	return offset
}
