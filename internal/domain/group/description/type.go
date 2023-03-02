package description

import "github.com/pkg/errors"

const (
	maxLength = 1000
)

var ErrWrongLength = errors.Errorf("description must be less than or equal to %d characters", maxLength)

type Description struct {
	value string
}

func New(description string) (Description, error) {
	if len([]rune(description)) > maxLength {
		return Description{}, ErrWrongLength
	}

	return Description{value: description}, nil
}

func (d Description) Value() string {
	return d.value
}
