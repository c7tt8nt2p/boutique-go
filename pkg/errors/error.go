package errors

import "github.com/go-kratos/kratos/v2/errors"

var ErrUInvalidCredentials = errors.Unauthorized("UNAUTHORIZED", "Invalid credentials")

func ErrValidationFailed(msg string) *errors.Error {
	return errors.New(422, "Validation failed", msg)
}
