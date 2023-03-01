package patronymic

import "github.com/pkg/errors"

const (
	MaxLength = 100
)

var ErrWrongLength = errors.Errorf("patronymic must be less than or equal to %d characters", MaxLength)

type Patronymic string

func New(patronymic string) (*Patronymic, error) {
	if len([]rune(patronymic)) > MaxLength {
		return nil, ErrWrongLength
	}

	p := Patronymic(patronymic)

	return &p, nil
}

func (p *Patronymic) String() string {
	return string(*p)
}
