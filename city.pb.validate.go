// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: city.proto

package arb_protos

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

// Validate checks the field values on City with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *City) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on City with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in CityMultiError, or nil if none found.
func (m *City) ValidateAll() error {
	return m.validate(true)
}

func (m *City) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Lat

	// no validation rules for Lng

	if len(errors) > 0 {
		return CityMultiError(errors)
	}

	return nil
}

// CityMultiError is an error wrapping multiple validation errors returned by
// City.ValidateAll() if the designated constraints aren't met.
type CityMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CityMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CityMultiError) AllErrors() []error { return m }

// CityValidationError is the validation error returned by City.Validate if the
// designated constraints aren't met.
type CityValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CityValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CityValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CityValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CityValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CityValidationError) ErrorName() string { return "CityValidationError" }

// Error satisfies the builtin error interface
func (e CityValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCity.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CityValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CityValidationError{}

// Validate checks the field values on CityRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CityRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CityRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CityRequestMultiError, or
// nil if none found.
func (m *CityRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CityRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for Lat

	// no validation rules for Lng

	if len(errors) > 0 {
		return CityRequestMultiError(errors)
	}

	return nil
}

// CityRequestMultiError is an error wrapping multiple validation errors
// returned by CityRequest.ValidateAll() if the designated constraints aren't met.
type CityRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CityRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CityRequestMultiError) AllErrors() []error { return m }

// CityRequestValidationError is the validation error returned by
// CityRequest.Validate if the designated constraints aren't met.
type CityRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CityRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CityRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CityRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CityRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CityRequestValidationError) ErrorName() string { return "CityRequestValidationError" }

// Error satisfies the builtin error interface
func (e CityRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCityRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CityRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CityRequestValidationError{}

// Validate checks the field values on CityResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CityResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CityResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CityResponseMultiError, or
// nil if none found.
func (m *CityResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CityResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Success

	if all {
		switch v := interface{}(m.GetData()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CityResponseValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CityResponseValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CityResponseValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CityResponseMultiError(errors)
	}

	return nil
}

// CityResponseMultiError is an error wrapping multiple validation errors
// returned by CityResponse.ValidateAll() if the designated constraints aren't met.
type CityResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CityResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CityResponseMultiError) AllErrors() []error { return m }

// CityResponseValidationError is the validation error returned by
// CityResponse.Validate if the designated constraints aren't met.
type CityResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CityResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CityResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CityResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CityResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CityResponseValidationError) ErrorName() string { return "CityResponseValidationError" }

// Error satisfies the builtin error interface
func (e CityResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCityResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CityResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CityResponseValidationError{}

// Validate checks the field values on Empty with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Empty) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Empty with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in EmptyMultiError, or nil if none found.
func (m *Empty) ValidateAll() error {
	return m.validate(true)
}

func (m *Empty) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return EmptyMultiError(errors)
	}

	return nil
}

// EmptyMultiError is an error wrapping multiple validation errors returned by
// Empty.ValidateAll() if the designated constraints aren't met.
type EmptyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m EmptyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m EmptyMultiError) AllErrors() []error { return m }

// EmptyValidationError is the validation error returned by Empty.Validate if
// the designated constraints aren't met.
type EmptyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EmptyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EmptyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EmptyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EmptyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EmptyValidationError) ErrorName() string { return "EmptyValidationError" }

// Error satisfies the builtin error interface
func (e EmptyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEmpty.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EmptyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EmptyValidationError{}

// Validate checks the field values on DeleteCityResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *DeleteCityResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DeleteCityResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// DeleteCityResponseMultiError, or nil if none found.
func (m *DeleteCityResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *DeleteCityResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Success

	if len(errors) > 0 {
		return DeleteCityResponseMultiError(errors)
	}

	return nil
}

// DeleteCityResponseMultiError is an error wrapping multiple validation errors
// returned by DeleteCityResponse.ValidateAll() if the designated constraints
// aren't met.
type DeleteCityResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DeleteCityResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DeleteCityResponseMultiError) AllErrors() []error { return m }

// DeleteCityResponseValidationError is the validation error returned by
// DeleteCityResponse.Validate if the designated constraints aren't met.
type DeleteCityResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeleteCityResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeleteCityResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeleteCityResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeleteCityResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeleteCityResponseValidationError) ErrorName() string {
	return "DeleteCityResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DeleteCityResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeleteCityResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeleteCityResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeleteCityResponseValidationError{}

// Validate checks the field values on GetCitiesResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetCitiesResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetCitiesResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetCitiesResponseMultiError, or nil if none found.
func (m *GetCitiesResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetCitiesResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Success

	for idx, item := range m.GetData() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GetCitiesResponseValidationError{
						field:  fmt.Sprintf("Data[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetCitiesResponseValidationError{
						field:  fmt.Sprintf("Data[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetCitiesResponseValidationError{
					field:  fmt.Sprintf("Data[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return GetCitiesResponseMultiError(errors)
	}

	return nil
}

// GetCitiesResponseMultiError is an error wrapping multiple validation errors
// returned by GetCitiesResponse.ValidateAll() if the designated constraints
// aren't met.
type GetCitiesResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetCitiesResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetCitiesResponseMultiError) AllErrors() []error { return m }

// GetCitiesResponseValidationError is the validation error returned by
// GetCitiesResponse.Validate if the designated constraints aren't met.
type GetCitiesResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetCitiesResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetCitiesResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetCitiesResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetCitiesResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetCitiesResponseValidationError) ErrorName() string {
	return "GetCitiesResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetCitiesResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetCitiesResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetCitiesResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetCitiesResponseValidationError{}

// Validate checks the field values on GetCityByIdRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetCityByIdRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetCityByIdRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetCityByIdRequestMultiError, or nil if none found.
func (m *GetCityByIdRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetCityByIdRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return GetCityByIdRequestMultiError(errors)
	}

	return nil
}

// GetCityByIdRequestMultiError is an error wrapping multiple validation errors
// returned by GetCityByIdRequest.ValidateAll() if the designated constraints
// aren't met.
type GetCityByIdRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetCityByIdRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetCityByIdRequestMultiError) AllErrors() []error { return m }

// GetCityByIdRequestValidationError is the validation error returned by
// GetCityByIdRequest.Validate if the designated constraints aren't met.
type GetCityByIdRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetCityByIdRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetCityByIdRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetCityByIdRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetCityByIdRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetCityByIdRequestValidationError) ErrorName() string {
	return "GetCityByIdRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetCityByIdRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetCityByIdRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetCityByIdRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetCityByIdRequestValidationError{}

// Validate checks the field values on GetCityByIdResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetCityByIdResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetCityByIdResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetCityByIdResponseMultiError, or nil if none found.
func (m *GetCityByIdResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetCityByIdResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Success

	if all {
		switch v := interface{}(m.GetData()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetCityByIdResponseValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetCityByIdResponseValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetCityByIdResponseValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetCityByIdResponseMultiError(errors)
	}

	return nil
}

// GetCityByIdResponseMultiError is an error wrapping multiple validation
// errors returned by GetCityByIdResponse.ValidateAll() if the designated
// constraints aren't met.
type GetCityByIdResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetCityByIdResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetCityByIdResponseMultiError) AllErrors() []error { return m }

// GetCityByIdResponseValidationError is the validation error returned by
// GetCityByIdResponse.Validate if the designated constraints aren't met.
type GetCityByIdResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetCityByIdResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetCityByIdResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetCityByIdResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetCityByIdResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetCityByIdResponseValidationError) ErrorName() string {
	return "GetCityByIdResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetCityByIdResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetCityByIdResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetCityByIdResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetCityByIdResponseValidationError{}

// Validate checks the field values on GetCityByNameRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetCityByNameRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetCityByNameRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetCityByNameRequestMultiError, or nil if none found.
func (m *GetCityByNameRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetCityByNameRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	if len(errors) > 0 {
		return GetCityByNameRequestMultiError(errors)
	}

	return nil
}

// GetCityByNameRequestMultiError is an error wrapping multiple validation
// errors returned by GetCityByNameRequest.ValidateAll() if the designated
// constraints aren't met.
type GetCityByNameRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetCityByNameRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetCityByNameRequestMultiError) AllErrors() []error { return m }

// GetCityByNameRequestValidationError is the validation error returned by
// GetCityByNameRequest.Validate if the designated constraints aren't met.
type GetCityByNameRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetCityByNameRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetCityByNameRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetCityByNameRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetCityByNameRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetCityByNameRequestValidationError) ErrorName() string {
	return "GetCityByNameRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetCityByNameRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetCityByNameRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetCityByNameRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetCityByNameRequestValidationError{}

// Validate checks the field values on GetCityByNameResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetCityByNameResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetCityByNameResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetCityByNameResponseMultiError, or nil if none found.
func (m *GetCityByNameResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetCityByNameResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Success

	if all {
		switch v := interface{}(m.GetData()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetCityByNameResponseValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetCityByNameResponseValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetCityByNameResponseValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetCityByNameResponseMultiError(errors)
	}

	return nil
}

// GetCityByNameResponseMultiError is an error wrapping multiple validation
// errors returned by GetCityByNameResponse.ValidateAll() if the designated
// constraints aren't met.
type GetCityByNameResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetCityByNameResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetCityByNameResponseMultiError) AllErrors() []error { return m }

// GetCityByNameResponseValidationError is the validation error returned by
// GetCityByNameResponse.Validate if the designated constraints aren't met.
type GetCityByNameResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetCityByNameResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetCityByNameResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetCityByNameResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetCityByNameResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetCityByNameResponseValidationError) ErrorName() string {
	return "GetCityByNameResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetCityByNameResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetCityByNameResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetCityByNameResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetCityByNameResponseValidationError{}