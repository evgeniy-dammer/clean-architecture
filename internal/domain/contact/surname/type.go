package surname

import "github.com/pkg/errors"

const (
	MaxLength = 100
)

var ErrWrongLength = errors.Errorf("surname must be less than or equal to %d characters", MaxLength)

type Surname string

func New(surname string) (*Surname, error) {
	if len([]rune(surname)) > MaxLength {
		return nil, ErrWrongLength
	}

	s := Surname(surname)

	return &s, nil
}

func (s Surname) String() string {
	return string(s)
}
