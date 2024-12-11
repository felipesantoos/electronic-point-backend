package errors

import (
	"errors"
	"strings"

	"golang.org/x/exp/slices"
)

const (
	unexpectedErrorMessage = "An unexpected error occurred. Please, contact the support."
	GENERIC                = iota
	INTERNAL
	VALIDATION
	CLIENT
	FORBIDDEN
	CONFLICT
)

type errorImpl struct {
	err      []string
	origin   int
	metadata map[string]interface{}
}

func new(errs []string, errType int, metadata map[string]interface{}) Error {
	return &errorImpl{errs, errType, metadata}
}

func New(err error) Error {
	return new([]string{err.Error()}, GENERIC, nil)
}

func NewWithMetadata(err error, metadata map[string]interface{}) Error {
	return new([]string{err.Error()}, GENERIC, metadata)
}

func NewFromString(message string) Error {
	return new([]string{message}, GENERIC, nil)
}

func NewInternal(err error) Error {
	return new([]string{err.Error()}, INTERNAL, nil)
}

func NewValidation(messages []string) Error {
	return new(messages, VALIDATION, nil)
}

func NewClient(messages ...string) Error {
	return new(messages, CLIENT, nil)
}

func NewForbidden(messages ...string) Error {
	return new(messages, FORBIDDEN, nil)
}

func NewConflict(messages ...string) Error {
	return new(messages, CONFLICT, nil)
}

func NewValidationWithMetadata(messages []string, metadata map[string]interface{}) Error {
	return new(messages, VALIDATION, metadata)
}

func NewValidationFromString(message string) Error {
	return new([]string{message}, VALIDATION, nil)
}

func NewConflictFromString(message string) Error {
	return new([]string{message}, CONFLICT, nil)
}

func NewUnexpected() Error {
	return NewInternal(errors.New(unexpectedErrorMessage))
}

func (e *errorImpl) String() string {
	return strings.Join(e.err, " & ")
}

func (e *errorImpl) Messages() []string {
	return e.err
}

func (e *errorImpl) CausedInternally() bool {
	return e.origin == INTERNAL
}

func (e *errorImpl) CausedByValidation() bool {
	return e.origin == VALIDATION
}

func (e *errorImpl) CausedByClient() bool {
	return e.origin == CLIENT
}

func (e *errorImpl) CausedByForbiddenAccess() bool {
	return e.origin == FORBIDDEN
}

func (e *errorImpl) CausedByConflict() bool {
	return e.origin == CONFLICT
}

func (e *errorImpl) Metadata() map[string]interface{} {
	return e.metadata
}

func (e *errorImpl) ValidationMessagesByMetadataFields(fields []string) []string {
	if e.metadata == nil || e.metadata["fields"] == nil {
		return nil
	}
	var metadataFields = e.metadata["fields"].([]string)
	var messages = []string{}
	for index, f := range metadataFields {
		if slices.Contains(fields, f) {
			messages = append(messages, e.Messages()[index])
		}
	}
	return messages
}
