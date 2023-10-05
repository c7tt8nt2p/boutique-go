// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: v1/cart.proto

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

// Validate checks the field values on AddItemReq with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *AddItemReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AddItemReq with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in AddItemReqMultiError, or
// nil if none found.
func (m *AddItemReq) ValidateAll() error {
	return m.validate(true)
}

func (m *AddItemReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for UserId

	if len(errors) > 0 {
		return AddItemReqMultiError(errors)
	}

	return nil
}

// AddItemReqMultiError is an error wrapping multiple validation errors
// returned by AddItemReq.ValidateAll() if the designated constraints aren't met.
type AddItemReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AddItemReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AddItemReqMultiError) AllErrors() []error { return m }

// AddItemReqValidationError is the validation error returned by
// AddItemReq.Validate if the designated constraints aren't met.
type AddItemReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddItemReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddItemReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddItemReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddItemReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddItemReqValidationError) ErrorName() string { return "AddItemReqValidationError" }

// Error satisfies the builtin error interface
func (e AddItemReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddItemReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddItemReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddItemReqValidationError{}

// Validate checks the field values on AddItemResp with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *AddItemResp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AddItemResp with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in AddItemRespMultiError, or
// nil if none found.
func (m *AddItemResp) ValidateAll() error {
	return m.validate(true)
}

func (m *AddItemResp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for UserId

	if len(errors) > 0 {
		return AddItemRespMultiError(errors)
	}

	return nil
}

// AddItemRespMultiError is an error wrapping multiple validation errors
// returned by AddItemResp.ValidateAll() if the designated constraints aren't met.
type AddItemRespMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AddItemRespMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AddItemRespMultiError) AllErrors() []error { return m }

// AddItemRespValidationError is the validation error returned by
// AddItemResp.Validate if the designated constraints aren't met.
type AddItemRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AddItemRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AddItemRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AddItemRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AddItemRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AddItemRespValidationError) ErrorName() string { return "AddItemRespValidationError" }

// Error satisfies the builtin error interface
func (e AddItemRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAddItemResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AddItemRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AddItemRespValidationError{}

// Validate checks the field values on ViewCartReq with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ViewCartReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ViewCartReq with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ViewCartReqMultiError, or
// nil if none found.
func (m *ViewCartReq) ValidateAll() error {
	return m.validate(true)
}

func (m *ViewCartReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for UserId

	if len(errors) > 0 {
		return ViewCartReqMultiError(errors)
	}

	return nil
}

// ViewCartReqMultiError is an error wrapping multiple validation errors
// returned by ViewCartReq.ValidateAll() if the designated constraints aren't met.
type ViewCartReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ViewCartReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ViewCartReqMultiError) AllErrors() []error { return m }

// ViewCartReqValidationError is the validation error returned by
// ViewCartReq.Validate if the designated constraints aren't met.
type ViewCartReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ViewCartReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ViewCartReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ViewCartReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ViewCartReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ViewCartReqValidationError) ErrorName() string { return "ViewCartReqValidationError" }

// Error satisfies the builtin error interface
func (e ViewCartReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sViewCartReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ViewCartReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ViewCartReqValidationError{}

// Validate checks the field values on ViewCartResp with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ViewCartResp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ViewCartResp with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ViewCartRespMultiError, or
// nil if none found.
func (m *ViewCartResp) ValidateAll() error {
	return m.validate(true)
}

func (m *ViewCartResp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetItems() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ViewCartRespValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ViewCartRespValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ViewCartRespValidationError{
					field:  fmt.Sprintf("Items[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ViewCartRespMultiError(errors)
	}

	return nil
}

// ViewCartRespMultiError is an error wrapping multiple validation errors
// returned by ViewCartResp.ValidateAll() if the designated constraints aren't met.
type ViewCartRespMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ViewCartRespMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ViewCartRespMultiError) AllErrors() []error { return m }

// ViewCartRespValidationError is the validation error returned by
// ViewCartResp.Validate if the designated constraints aren't met.
type ViewCartRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ViewCartRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ViewCartRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ViewCartRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ViewCartRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ViewCartRespValidationError) ErrorName() string { return "ViewCartRespValidationError" }

// Error satisfies the builtin error interface
func (e ViewCartRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sViewCartResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ViewCartRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ViewCartRespValidationError{}

// Validate checks the field values on ViewCartResp_Item with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ViewCartResp_Item) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ViewCartResp_Item with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ViewCartResp_ItemMultiError, or nil if none found.
func (m *ViewCartResp_Item) ValidateAll() error {
	return m.validate(true)
}

func (m *ViewCartResp_Item) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ItemId

	// no validation rules for Quantity

	if len(errors) > 0 {
		return ViewCartResp_ItemMultiError(errors)
	}

	return nil
}

// ViewCartResp_ItemMultiError is an error wrapping multiple validation errors
// returned by ViewCartResp_Item.ValidateAll() if the designated constraints
// aren't met.
type ViewCartResp_ItemMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ViewCartResp_ItemMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ViewCartResp_ItemMultiError) AllErrors() []error { return m }

// ViewCartResp_ItemValidationError is the validation error returned by
// ViewCartResp_Item.Validate if the designated constraints aren't met.
type ViewCartResp_ItemValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ViewCartResp_ItemValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ViewCartResp_ItemValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ViewCartResp_ItemValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ViewCartResp_ItemValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ViewCartResp_ItemValidationError) ErrorName() string {
	return "ViewCartResp_ItemValidationError"
}

// Error satisfies the builtin error interface
func (e ViewCartResp_ItemValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sViewCartResp_Item.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ViewCartResp_ItemValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ViewCartResp_ItemValidationError{}
