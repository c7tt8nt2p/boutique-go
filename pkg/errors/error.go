package errors

import "github.com/go-kratos/kratos/v2/errors"

type AppErr struct {
	Err *errors.Error
}

var ErrUInvalidCredentials = errors.Unauthorized("UNAUTHORIZED", "Invalid credentials")

func AppInvalidCredentialsErr() *AppErr {
	return &AppErr{Err: ErrUInvalidCredentials}
}

func AppInternalErr(msg string) *AppErr {
	err := errors.BadRequest("BAD REQUEST", msg)
	return &AppErr{Err: err}
}

func AppValidationErr(msg string) *AppErr {
	err := errors.New(422, "VALIDATION FAILED", msg)
	return &AppErr{Err: err}
}
