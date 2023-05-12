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
)

type Error interface {
	String() string
	Messages() []string
	CausedInternally() bool
	CausedByValidation() bool
	Metadata() map[string]interface{}
	ValidationMessagesByMetadataFields(field []string) []string
}

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

func NewValidationWithMetadata(messages []string, metadata map[string]interface{}) Error {
	return new(messages, VALIDATION, metadata)
}

func NewValidationFromString(message string) Error {
	return new([]string{message}, VALIDATION, nil)
}

func NewUnexpected() Error {
	return NewInternal(errors.New(unexpectedErrorMessage))
}

func (instance *errorImpl) String() string {
	return strings.Join(instance.err, " & ")
}

func (instance *errorImpl) Messages() []string {
	return instance.err
}

func (instance *errorImpl) CausedInternally() bool {
	return instance.origin == INTERNAL
}

func (instance *errorImpl) CausedByValidation() bool {
	return instance.origin == VALIDATION
}

func (instance *errorImpl) Metadata() map[string]interface{} {
	return instance.metadata
}

func (instance *errorImpl) ValidationMessagesByMetadataFields(fields []string) []string {
	if instance.metadata == nil || instance.metadata["fields"] == nil {
		return nil
	}
	var metadataFields = instance.metadata["fields"].([]string)
	var messages = []string{}
	for index, f := range metadataFields {
		if slices.Contains(fields, f) {
			messages = append(messages, instance.Messages()[index])
		}
	}
	return messages
}
