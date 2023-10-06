package util

import (
	"github.com/google/uuid"
	"github.com/kx-boutique/pkg/errors"
)

func ParseUUID(s string) (uuid.UUID, error) {
	parsed, err := uuid.Parse(s)
	if err != nil {
		return uuid.Nil, errors.ErrValidationFailed(err.Error())
	}
	return parsed, nil
}
