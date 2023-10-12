package util

import (
	"github.com/google/uuid"
	"github.com/kx-boutique/pkg/errors"
)

func ParseUUID(s string) uuid.UUID {
	parsed, err := uuid.Parse(s)
	if err != nil {
		panic(errors.AppInternalErr(err.Error()))
	}
	return parsed
}
