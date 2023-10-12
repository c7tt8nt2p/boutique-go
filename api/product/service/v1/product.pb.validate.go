// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: v1/product.proto

package v1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// define the regex for a UUID once up-front
var _product_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on CreateProductReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CreateProductReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateProductReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateProductReqMultiError, or nil if none found.
func (m *CreateProductReq) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateProductReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetName()) < 1 {
		err := CreateProductReqValidationError{
			field:  "Name",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetDescription()) < 5 {
		err := CreateProductReqValidationError{
			field:  "Description",
			reason: "value length must be at least 5 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetStock() <= 0 {
		err := CreateProductReqValidationError{
			field:  "Stock",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetUnitPrice() <= 0 {
		err := CreateProductReqValidationError{
			field:  "UnitPrice",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return CreateProductReqMultiError(errors)
	}

	return nil
}

// CreateProductReqMultiError is an error wrapping multiple validation errors
// returned by CreateProductReq.ValidateAll() if the designated constraints
// aren't met.
type CreateProductReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateProductReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateProductReqMultiError) AllErrors() []error { return m }

// CreateProductReqValidationError is the validation error returned by
// CreateProductReq.Validate if the designated constraints aren't met.
type CreateProductReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateProductReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateProductReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateProductReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateProductReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateProductReqValidationError) ErrorName() string { return "CreateProductReqValidationError" }

// Error satisfies the builtin error interface
func (e CreateProductReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateProductReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateProductReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateProductReqValidationError{}

// Validate checks the field values on CreateProductResp with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CreateProductResp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateProductResp with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateProductRespMultiError, or nil if none found.
func (m *CreateProductResp) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateProductResp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Description

	// no validation rules for Stock

	// no validation rules for UnitPrice

	if len(errors) > 0 {
		return CreateProductRespMultiError(errors)
	}

	return nil
}

// CreateProductRespMultiError is an error wrapping multiple validation errors
// returned by CreateProductResp.ValidateAll() if the designated constraints
// aren't met.
type CreateProductRespMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateProductRespMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateProductRespMultiError) AllErrors() []error { return m }

// CreateProductRespValidationError is the validation error returned by
// CreateProductResp.Validate if the designated constraints aren't met.
type CreateProductRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateProductRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateProductRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateProductRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateProductRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateProductRespValidationError) ErrorName() string {
	return "CreateProductRespValidationError"
}

// Error satisfies the builtin error interface
func (e CreateProductRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateProductResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateProductRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateProductRespValidationError{}

// Validate checks the field values on GetProductByIdReq with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetProductByIdReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetProductByIdReq with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetProductByIdReqMultiError, or nil if none found.
func (m *GetProductByIdReq) ValidateAll() error {
	return m.validate(true)
}

func (m *GetProductByIdReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if err := m._validateUuid(m.GetId()); err != nil {
		err = GetProductByIdReqValidationError{
			field:  "Id",
			reason: "value must be a valid UUID",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GetProductByIdReqMultiError(errors)
	}

	return nil
}

func (m *GetProductByIdReq) _validateUuid(uuid string) error {
	if matched := _product_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// GetProductByIdReqMultiError is an error wrapping multiple validation errors
// returned by GetProductByIdReq.ValidateAll() if the designated constraints
// aren't met.
type GetProductByIdReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetProductByIdReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetProductByIdReqMultiError) AllErrors() []error { return m }

// GetProductByIdReqValidationError is the validation error returned by
// GetProductByIdReq.Validate if the designated constraints aren't met.
type GetProductByIdReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetProductByIdReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetProductByIdReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetProductByIdReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetProductByIdReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetProductByIdReqValidationError) ErrorName() string {
	return "GetProductByIdReqValidationError"
}

// Error satisfies the builtin error interface
func (e GetProductByIdReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetProductByIdReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetProductByIdReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetProductByIdReqValidationError{}

// Validate checks the field values on GetProductByIdResp with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetProductByIdResp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetProductByIdResp with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetProductByIdRespMultiError, or nil if none found.
func (m *GetProductByIdResp) ValidateAll() error {
	return m.validate(true)
}

func (m *GetProductByIdResp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Description

	// no validation rules for Stock

	// no validation rules for UnitPrice

	if all {
		switch v := interface{}(m.GetCreatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetProductByIdRespValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetProductByIdRespValidationError{
					field:  "CreatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCreatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetProductByIdRespValidationError{
				field:  "CreatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetUpdatedAt()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetProductByIdRespValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetProductByIdRespValidationError{
					field:  "UpdatedAt",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUpdatedAt()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetProductByIdRespValidationError{
				field:  "UpdatedAt",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetProductByIdRespMultiError(errors)
	}

	return nil
}

// GetProductByIdRespMultiError is an error wrapping multiple validation errors
// returned by GetProductByIdResp.ValidateAll() if the designated constraints
// aren't met.
type GetProductByIdRespMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetProductByIdRespMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetProductByIdRespMultiError) AllErrors() []error { return m }

// GetProductByIdRespValidationError is the validation error returned by
// GetProductByIdResp.Validate if the designated constraints aren't met.
type GetProductByIdRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetProductByIdRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetProductByIdRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetProductByIdRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetProductByIdRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetProductByIdRespValidationError) ErrorName() string {
	return "GetProductByIdRespValidationError"
}

// Error satisfies the builtin error interface
func (e GetProductByIdRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetProductByIdResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetProductByIdRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetProductByIdRespValidationError{}

// Validate checks the field values on ValidatePurchasableProductReq with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ValidatePurchasableProductReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ValidatePurchasableProductReq with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// ValidatePurchasableProductReqMultiError, or nil if none found.
func (m *ValidatePurchasableProductReq) ValidateAll() error {
	return m.validate(true)
}

func (m *ValidatePurchasableProductReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if err := m._validateUuid(m.GetId()); err != nil {
		err = ValidatePurchasableProductReqValidationError{
			field:  "Id",
			reason: "value must be a valid UUID",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if m.GetQty() <= 0 {
		err := ValidatePurchasableProductReqValidationError{
			field:  "Qty",
			reason: "value must be greater than 0",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return ValidatePurchasableProductReqMultiError(errors)
	}

	return nil
}

func (m *ValidatePurchasableProductReq) _validateUuid(uuid string) error {
	if matched := _product_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// ValidatePurchasableProductReqMultiError is an error wrapping multiple
// validation errors returned by ValidatePurchasableProductReq.ValidateAll()
// if the designated constraints aren't met.
type ValidatePurchasableProductReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ValidatePurchasableProductReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ValidatePurchasableProductReqMultiError) AllErrors() []error { return m }

// ValidatePurchasableProductReqValidationError is the validation error
// returned by ValidatePurchasableProductReq.Validate if the designated
// constraints aren't met.
type ValidatePurchasableProductReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ValidatePurchasableProductReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ValidatePurchasableProductReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ValidatePurchasableProductReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ValidatePurchasableProductReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ValidatePurchasableProductReqValidationError) ErrorName() string {
	return "ValidatePurchasableProductReqValidationError"
}

// Error satisfies the builtin error interface
func (e ValidatePurchasableProductReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sValidatePurchasableProductReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ValidatePurchasableProductReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ValidatePurchasableProductReqValidationError{}

// Validate checks the field values on ValidatePurchasableProductResp with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ValidatePurchasableProductResp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ValidatePurchasableProductResp with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// ValidatePurchasableProductRespMultiError, or nil if none found.
func (m *ValidatePurchasableProductResp) ValidateAll() error {
	return m.validate(true)
}

func (m *ValidatePurchasableProductResp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Validated

	// no validation rules for ValidationMessage

	if len(errors) > 0 {
		return ValidatePurchasableProductRespMultiError(errors)
	}

	return nil
}

// ValidatePurchasableProductRespMultiError is an error wrapping multiple
// validation errors returned by ValidatePurchasableProductResp.ValidateAll()
// if the designated constraints aren't met.
type ValidatePurchasableProductRespMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ValidatePurchasableProductRespMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ValidatePurchasableProductRespMultiError) AllErrors() []error { return m }

// ValidatePurchasableProductRespValidationError is the validation error
// returned by ValidatePurchasableProductResp.Validate if the designated
// constraints aren't met.
type ValidatePurchasableProductRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ValidatePurchasableProductRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ValidatePurchasableProductRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ValidatePurchasableProductRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ValidatePurchasableProductRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ValidatePurchasableProductRespValidationError) ErrorName() string {
	return "ValidatePurchasableProductRespValidationError"
}

// Error satisfies the builtin error interface
func (e ValidatePurchasableProductRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sValidatePurchasableProductResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ValidatePurchasableProductRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ValidatePurchasableProductRespValidationError{}
