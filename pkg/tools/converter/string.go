package converter

import (
	"github.com/google/uuid"
)

func StringToUUID(str string) uuid.UUID {
	if len(str) == 0 {
		return uuid.Nil
	}

	value, err := uuid.Parse(str)
	if err != nil {
		return uuid.Nil
	}

	return value
}
