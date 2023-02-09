package errors

import (
	"dit_backend/src/infra"
	"errors"
)

const (
	unexpectedErrorMessage = "An unexpected error occurred. Please, contact the support."
	GENERIC                = iota
	INTERNAL
	VALIDATION
)

type Error interface {
	Errors() []string
	CausedInternally() bool
	CausedByValidation() bool
}

type errorImpl struct {
	err    []string
	origin int
}

func new(errs []string, errType int) Error {
	return &errorImpl{errs, errType}
}

func New(err error) Error {
	return new([]string{err.Error()}, GENERIC)
}

func NewFromString(message string) Error {
	return new([]string{message}, GENERIC)
}

func NewInternal(err error) Error {
	return new([]string{err.Error()}, INTERNAL)
}

func NewValidation(messages []string) Error {
	return new(messages, VALIDATION)
}

func NewFromInfra(err infra.Error) Error {
	if err.IsInternal() {
		return NewInternal(err.Native())
	}
	return New(err.Native())
}

func NewUnexpectedError() Error {
	return NewInternal(errors.New(unexpectedErrorMessage))
}

func (instance *errorImpl) Errors() []string {
	return instance.err
}

func (instance *errorImpl) CausedInternally() bool {
	return instance.origin == INTERNAL
}

func (instance *errorImpl) CausedByValidation() bool {
	return instance.origin == VALIDATION
}
