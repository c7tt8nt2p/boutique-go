package errors

import "github.com/go-kratos/kratos/v2/errors"

var ErrAuthFail = errors.New(401, "Authentication failed", "Missing token or token incorrect")

func ErrValidationFailed(msg string) *errors.Error {
	return errors.New(422, "Validation failed", msg)
}
