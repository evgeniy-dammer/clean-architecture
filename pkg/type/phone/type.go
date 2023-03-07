package phone

import (
	"strings"
)

type Phone struct {
	value string
}

func New(phone string) *Phone {
	return &Phone{
		value: getNumbers(phone),
	}
}

func (p Phone) String() string {
	return p.value
}

func (p *Phone) Equal(phoneNumber Phone) bool {
	return p.value == phoneNumber.value
}

func (p *Phone) IsEmpty() bool {
	return len(strings.TrimSpace(p.value)) == 0
}
