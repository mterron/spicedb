// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: authzed/api/v0/developer.proto

package v0

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
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
)

// Validate checks the field values on FormatSchemaRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *FormatSchemaRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Schema

	return nil
}

// FormatSchemaRequestValidationError is the validation error returned by
// FormatSchemaRequest.Validate if the designated constraints aren't met.
type FormatSchemaRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FormatSchemaRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FormatSchemaRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FormatSchemaRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FormatSchemaRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FormatSchemaRequestValidationError) ErrorName() string {
	return "FormatSchemaRequestValidationError"
}

// Error satisfies the builtin error interface
func (e FormatSchemaRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFormatSchemaRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FormatSchemaRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FormatSchemaRequestValidationError{}

// Validate checks the field values on FormatSchemaResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *FormatSchemaResponse) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetError()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return FormatSchemaResponseValidationError{
				field:  "Error",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for FormattedSchema

	return nil
}

// FormatSchemaResponseValidationError is the validation error returned by
// FormatSchemaResponse.Validate if the designated constraints aren't met.
type FormatSchemaResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FormatSchemaResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FormatSchemaResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FormatSchemaResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FormatSchemaResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FormatSchemaResponseValidationError) ErrorName() string {
	return "FormatSchemaResponseValidationError"
}

// Error satisfies the builtin error interface
func (e FormatSchemaResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFormatSchemaResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FormatSchemaResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FormatSchemaResponseValidationError{}

// Validate checks the field values on UpgradeSchemaRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpgradeSchemaRequest) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// UpgradeSchemaRequestValidationError is the validation error returned by
// UpgradeSchemaRequest.Validate if the designated constraints aren't met.
type UpgradeSchemaRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpgradeSchemaRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpgradeSchemaRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpgradeSchemaRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpgradeSchemaRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpgradeSchemaRequestValidationError) ErrorName() string {
	return "UpgradeSchemaRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpgradeSchemaRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpgradeSchemaRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpgradeSchemaRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpgradeSchemaRequestValidationError{}

// Validate checks the field values on UpgradeSchemaResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpgradeSchemaResponse) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetError()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpgradeSchemaResponseValidationError{
				field:  "Error",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for UpgradedSchema

	return nil
}

// UpgradeSchemaResponseValidationError is the validation error returned by
// UpgradeSchemaResponse.Validate if the designated constraints aren't met.
type UpgradeSchemaResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpgradeSchemaResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpgradeSchemaResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpgradeSchemaResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpgradeSchemaResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpgradeSchemaResponseValidationError) ErrorName() string {
	return "UpgradeSchemaResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UpgradeSchemaResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpgradeSchemaResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpgradeSchemaResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpgradeSchemaResponseValidationError{}

// Validate checks the field values on ShareRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ShareRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Schema

	// no validation rules for RelationshipsYaml

	// no validation rules for ValidationYaml

	// no validation rules for AssertionsYaml

	return nil
}

// ShareRequestValidationError is the validation error returned by
// ShareRequest.Validate if the designated constraints aren't met.
type ShareRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ShareRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ShareRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ShareRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ShareRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ShareRequestValidationError) ErrorName() string { return "ShareRequestValidationError" }

// Error satisfies the builtin error interface
func (e ShareRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sShareRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ShareRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ShareRequestValidationError{}

// Validate checks the field values on ShareResponse with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ShareResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for ShareReference

	return nil
}

// ShareResponseValidationError is the validation error returned by
// ShareResponse.Validate if the designated constraints aren't met.
type ShareResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ShareResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ShareResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ShareResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ShareResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ShareResponseValidationError) ErrorName() string { return "ShareResponseValidationError" }

// Error satisfies the builtin error interface
func (e ShareResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sShareResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ShareResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ShareResponseValidationError{}

// Validate checks the field values on LookupShareRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *LookupShareRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for ShareReference

	return nil
}

// LookupShareRequestValidationError is the validation error returned by
// LookupShareRequest.Validate if the designated constraints aren't met.
type LookupShareRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LookupShareRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LookupShareRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LookupShareRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LookupShareRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LookupShareRequestValidationError) ErrorName() string {
	return "LookupShareRequestValidationError"
}

// Error satisfies the builtin error interface
func (e LookupShareRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLookupShareRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LookupShareRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LookupShareRequestValidationError{}

// Validate checks the field values on LookupShareResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *LookupShareResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Status

	// no validation rules for Schema

	// no validation rules for RelationshipsYaml

	// no validation rules for ValidationYaml

	// no validation rules for AssertionsYaml

	return nil
}

// LookupShareResponseValidationError is the validation error returned by
// LookupShareResponse.Validate if the designated constraints aren't met.
type LookupShareResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LookupShareResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LookupShareResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LookupShareResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LookupShareResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LookupShareResponseValidationError) ErrorName() string {
	return "LookupShareResponseValidationError"
}

// Error satisfies the builtin error interface
func (e LookupShareResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLookupShareResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LookupShareResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LookupShareResponseValidationError{}

// Validate checks the field values on RequestContext with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *RequestContext) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Schema

	for idx, item := range m.GetRelationships() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RequestContextValidationError{
					field:  fmt.Sprintf("Relationships[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// RequestContextValidationError is the validation error returned by
// RequestContext.Validate if the designated constraints aren't met.
type RequestContextValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RequestContextValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RequestContextValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RequestContextValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RequestContextValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RequestContextValidationError) ErrorName() string { return "RequestContextValidationError" }

// Error satisfies the builtin error interface
func (e RequestContextValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRequestContext.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RequestContextValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RequestContextValidationError{}

// Validate checks the field values on EditCheckRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *EditCheckRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetContext()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EditCheckRequestValidationError{
				field:  "Context",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetCheckRelationships() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return EditCheckRequestValidationError{
					field:  fmt.Sprintf("CheckRelationships[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// EditCheckRequestValidationError is the validation error returned by
// EditCheckRequest.Validate if the designated constraints aren't met.
type EditCheckRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EditCheckRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EditCheckRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EditCheckRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EditCheckRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EditCheckRequestValidationError) ErrorName() string { return "EditCheckRequestValidationError" }

// Error satisfies the builtin error interface
func (e EditCheckRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEditCheckRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EditCheckRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EditCheckRequestValidationError{}

// Validate checks the field values on EditCheckResult with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *EditCheckResult) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetRelationship()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EditCheckResultValidationError{
				field:  "Relationship",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for IsMember

	if v, ok := interface{}(m.GetError()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EditCheckResultValidationError{
				field:  "Error",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// EditCheckResultValidationError is the validation error returned by
// EditCheckResult.Validate if the designated constraints aren't met.
type EditCheckResultValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EditCheckResultValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EditCheckResultValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EditCheckResultValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EditCheckResultValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EditCheckResultValidationError) ErrorName() string { return "EditCheckResultValidationError" }

// Error satisfies the builtin error interface
func (e EditCheckResultValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEditCheckResult.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EditCheckResultValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EditCheckResultValidationError{}

// Validate checks the field values on EditCheckResponse with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *EditCheckResponse) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetRequestErrors() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return EditCheckResponseValidationError{
					field:  fmt.Sprintf("RequestErrors[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetCheckResults() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return EditCheckResponseValidationError{
					field:  fmt.Sprintf("CheckResults[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// EditCheckResponseValidationError is the validation error returned by
// EditCheckResponse.Validate if the designated constraints aren't met.
type EditCheckResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EditCheckResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EditCheckResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EditCheckResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EditCheckResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EditCheckResponseValidationError) ErrorName() string {
	return "EditCheckResponseValidationError"
}

// Error satisfies the builtin error interface
func (e EditCheckResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEditCheckResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EditCheckResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EditCheckResponseValidationError{}

// Validate checks the field values on ValidateRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ValidateRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetContext()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ValidateRequestValidationError{
				field:  "Context",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for ValidationYaml

	// no validation rules for UpdateValidationYaml

	// no validation rules for AssertionsYaml

	return nil
}

// ValidateRequestValidationError is the validation error returned by
// ValidateRequest.Validate if the designated constraints aren't met.
type ValidateRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ValidateRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ValidateRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ValidateRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ValidateRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ValidateRequestValidationError) ErrorName() string { return "ValidateRequestValidationError" }

// Error satisfies the builtin error interface
func (e ValidateRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sValidateRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ValidateRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ValidateRequestValidationError{}

// Validate checks the field values on ValidateResponse with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ValidateResponse) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetRequestErrors() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ValidateResponseValidationError{
					field:  fmt.Sprintf("RequestErrors[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetValidationErrors() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ValidateResponseValidationError{
					field:  fmt.Sprintf("ValidationErrors[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for UpdatedValidationYaml

	return nil
}

// ValidateResponseValidationError is the validation error returned by
// ValidateResponse.Validate if the designated constraints aren't met.
type ValidateResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ValidateResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ValidateResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ValidateResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ValidateResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ValidateResponseValidationError) ErrorName() string { return "ValidateResponseValidationError" }

// Error satisfies the builtin error interface
func (e ValidateResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sValidateResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ValidateResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ValidateResponseValidationError{}

// Validate checks the field values on DeveloperError with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *DeveloperError) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Message

	// no validation rules for Line

	// no validation rules for Column

	// no validation rules for Source

	// no validation rules for Kind

	// no validation rules for Context

	return nil
}

// DeveloperErrorValidationError is the validation error returned by
// DeveloperError.Validate if the designated constraints aren't met.
type DeveloperErrorValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DeveloperErrorValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DeveloperErrorValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DeveloperErrorValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DeveloperErrorValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DeveloperErrorValidationError) ErrorName() string { return "DeveloperErrorValidationError" }

// Error satisfies the builtin error interface
func (e DeveloperErrorValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDeveloperError.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DeveloperErrorValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DeveloperErrorValidationError{}
