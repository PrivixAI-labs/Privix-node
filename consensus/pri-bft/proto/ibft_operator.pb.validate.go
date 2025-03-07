// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: consensus/ibft/proto/ibft_operator.proto

package proto

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

// Validate checks the field values on IbftStatusResp with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *IbftStatusResp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IbftStatusResp with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in IbftStatusRespMultiError,
// or nil if none found.
func (m *IbftStatusResp) ValidateAll() error {
	return m.validate(true)
}

func (m *IbftStatusResp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Key

	if len(errors) > 0 {
		return IbftStatusRespMultiError(errors)
	}

	return nil
}

// IbftStatusRespMultiError is an error wrapping multiple validation errors
// returned by IbftStatusResp.ValidateAll() if the designated constraints
// aren't met.
type IbftStatusRespMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IbftStatusRespMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IbftStatusRespMultiError) AllErrors() []error { return m }

// IbftStatusRespValidationError is the validation error returned by
// IbftStatusResp.Validate if the designated constraints aren't met.
type IbftStatusRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IbftStatusRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IbftStatusRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IbftStatusRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IbftStatusRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IbftStatusRespValidationError) ErrorName() string { return "IbftStatusRespValidationError" }

// Error satisfies the builtin error interface
func (e IbftStatusRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIbftStatusResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IbftStatusRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IbftStatusRespValidationError{}

// Validate checks the field values on SnapshotReq with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SnapshotReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SnapshotReq with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SnapshotReqMultiError, or
// nil if none found.
func (m *SnapshotReq) ValidateAll() error {
	return m.validate(true)
}

func (m *SnapshotReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Latest

	// no validation rules for Number

	if len(errors) > 0 {
		return SnapshotReqMultiError(errors)
	}

	return nil
}

// SnapshotReqMultiError is an error wrapping multiple validation errors
// returned by SnapshotReq.ValidateAll() if the designated constraints aren't met.
type SnapshotReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SnapshotReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SnapshotReqMultiError) AllErrors() []error { return m }

// SnapshotReqValidationError is the validation error returned by
// SnapshotReq.Validate if the designated constraints aren't met.
type SnapshotReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SnapshotReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SnapshotReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SnapshotReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SnapshotReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SnapshotReqValidationError) ErrorName() string { return "SnapshotReqValidationError" }

// Error satisfies the builtin error interface
func (e SnapshotReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSnapshotReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SnapshotReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SnapshotReqValidationError{}

// Validate checks the field values on Snapshot with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Snapshot) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Snapshot with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SnapshotMultiError, or nil
// if none found.
func (m *Snapshot) ValidateAll() error {
	return m.validate(true)
}

func (m *Snapshot) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetValidators() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, SnapshotValidationError{
						field:  fmt.Sprintf("Validators[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, SnapshotValidationError{
						field:  fmt.Sprintf("Validators[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SnapshotValidationError{
					field:  fmt.Sprintf("Validators[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Number

	// no validation rules for Hash

	for idx, item := range m.GetVotes() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, SnapshotValidationError{
						field:  fmt.Sprintf("Votes[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, SnapshotValidationError{
						field:  fmt.Sprintf("Votes[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SnapshotValidationError{
					field:  fmt.Sprintf("Votes[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return SnapshotMultiError(errors)
	}

	return nil
}

// SnapshotMultiError is an error wrapping multiple validation errors returned
// by Snapshot.ValidateAll() if the designated constraints aren't met.
type SnapshotMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SnapshotMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SnapshotMultiError) AllErrors() []error { return m }

// SnapshotValidationError is the validation error returned by
// Snapshot.Validate if the designated constraints aren't met.
type SnapshotValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SnapshotValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SnapshotValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SnapshotValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SnapshotValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SnapshotValidationError) ErrorName() string { return "SnapshotValidationError" }

// Error satisfies the builtin error interface
func (e SnapshotValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSnapshot.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SnapshotValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SnapshotValidationError{}

// Validate checks the field values on ProposeReq with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ProposeReq) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ProposeReq with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ProposeReqMultiError, or
// nil if none found.
func (m *ProposeReq) ValidateAll() error {
	return m.validate(true)
}

func (m *ProposeReq) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Address

	// no validation rules for Auth

	if len(errors) > 0 {
		return ProposeReqMultiError(errors)
	}

	return nil
}

// ProposeReqMultiError is an error wrapping multiple validation errors
// returned by ProposeReq.ValidateAll() if the designated constraints aren't met.
type ProposeReqMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ProposeReqMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ProposeReqMultiError) AllErrors() []error { return m }

// ProposeReqValidationError is the validation error returned by
// ProposeReq.Validate if the designated constraints aren't met.
type ProposeReqValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProposeReqValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProposeReqValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProposeReqValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProposeReqValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProposeReqValidationError) ErrorName() string { return "ProposeReqValidationError" }

// Error satisfies the builtin error interface
func (e ProposeReqValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProposeReq.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProposeReqValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProposeReqValidationError{}

// Validate checks the field values on CandidatesResp with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CandidatesResp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CandidatesResp with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CandidatesRespMultiError,
// or nil if none found.
func (m *CandidatesResp) ValidateAll() error {
	return m.validate(true)
}

func (m *CandidatesResp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetCandidates() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, CandidatesRespValidationError{
						field:  fmt.Sprintf("Candidates[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, CandidatesRespValidationError{
						field:  fmt.Sprintf("Candidates[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return CandidatesRespValidationError{
					field:  fmt.Sprintf("Candidates[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return CandidatesRespMultiError(errors)
	}

	return nil
}

// CandidatesRespMultiError is an error wrapping multiple validation errors
// returned by CandidatesResp.ValidateAll() if the designated constraints
// aren't met.
type CandidatesRespMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CandidatesRespMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CandidatesRespMultiError) AllErrors() []error { return m }

// CandidatesRespValidationError is the validation error returned by
// CandidatesResp.Validate if the designated constraints aren't met.
type CandidatesRespValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CandidatesRespValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CandidatesRespValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CandidatesRespValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CandidatesRespValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CandidatesRespValidationError) ErrorName() string { return "CandidatesRespValidationError" }

// Error satisfies the builtin error interface
func (e CandidatesRespValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCandidatesResp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CandidatesRespValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CandidatesRespValidationError{}

// Validate checks the field values on Candidate with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Candidate) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Candidate with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CandidateMultiError, or nil
// if none found.
func (m *Candidate) ValidateAll() error {
	return m.validate(true)
}

func (m *Candidate) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Address

	// no validation rules for BlsPubkey

	// no validation rules for Auth

	if len(errors) > 0 {
		return CandidateMultiError(errors)
	}

	return nil
}

// CandidateMultiError is an error wrapping multiple validation errors returned
// by Candidate.ValidateAll() if the designated constraints aren't met.
type CandidateMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CandidateMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CandidateMultiError) AllErrors() []error { return m }

// CandidateValidationError is the validation error returned by
// Candidate.Validate if the designated constraints aren't met.
type CandidateValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CandidateValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CandidateValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CandidateValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CandidateValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CandidateValidationError) ErrorName() string { return "CandidateValidationError" }

// Error satisfies the builtin error interface
func (e CandidateValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCandidate.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CandidateValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CandidateValidationError{}

// Validate checks the field values on Snapshot_Validator with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *Snapshot_Validator) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Snapshot_Validator with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// Snapshot_ValidatorMultiError, or nil if none found.
func (m *Snapshot_Validator) ValidateAll() error {
	return m.validate(true)
}

func (m *Snapshot_Validator) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Type

	// no validation rules for Address

	// no validation rules for Data

	if len(errors) > 0 {
		return Snapshot_ValidatorMultiError(errors)
	}

	return nil
}

// Snapshot_ValidatorMultiError is an error wrapping multiple validation errors
// returned by Snapshot_Validator.ValidateAll() if the designated constraints
// aren't met.
type Snapshot_ValidatorMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Snapshot_ValidatorMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Snapshot_ValidatorMultiError) AllErrors() []error { return m }

// Snapshot_ValidatorValidationError is the validation error returned by
// Snapshot_Validator.Validate if the designated constraints aren't met.
type Snapshot_ValidatorValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Snapshot_ValidatorValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Snapshot_ValidatorValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Snapshot_ValidatorValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Snapshot_ValidatorValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Snapshot_ValidatorValidationError) ErrorName() string {
	return "Snapshot_ValidatorValidationError"
}

// Error satisfies the builtin error interface
func (e Snapshot_ValidatorValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSnapshot_Validator.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Snapshot_ValidatorValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Snapshot_ValidatorValidationError{}

// Validate checks the field values on Snapshot_Vote with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Snapshot_Vote) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Snapshot_Vote with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in Snapshot_VoteMultiError, or
// nil if none found.
func (m *Snapshot_Vote) ValidateAll() error {
	return m.validate(true)
}

func (m *Snapshot_Vote) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Validator

	// no validation rules for Proposed

	// no validation rules for Auth

	if len(errors) > 0 {
		return Snapshot_VoteMultiError(errors)
	}

	return nil
}

// Snapshot_VoteMultiError is an error wrapping multiple validation errors
// returned by Snapshot_Vote.ValidateAll() if the designated constraints
// aren't met.
type Snapshot_VoteMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Snapshot_VoteMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Snapshot_VoteMultiError) AllErrors() []error { return m }

// Snapshot_VoteValidationError is the validation error returned by
// Snapshot_Vote.Validate if the designated constraints aren't met.
type Snapshot_VoteValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Snapshot_VoteValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Snapshot_VoteValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Snapshot_VoteValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Snapshot_VoteValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Snapshot_VoteValidationError) ErrorName() string { return "Snapshot_VoteValidationError" }

// Error satisfies the builtin error interface
func (e Snapshot_VoteValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSnapshot_Vote.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Snapshot_VoteValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Snapshot_VoteValidationError{}
