package infra

import (
	"errors"
	"fmt"
)

type Error interface {
	Err() string
	IsInternal() bool
	Native() error
}

type sourceErr struct {
	err        error
	isInternal bool
}

func NewSourceErr(err error) Error {
	return &sourceErr{err: err}
}

func NewSourceErrFromStr(message string) Error {
	err := errors.New(message)
	return &sourceErr{err: err}
}

func NewInternalSourceErr(err error) Error {
	return &sourceErr{err, true}
}

func NewUnexpectedSourceErr() Error {
	err := errors.New(fmt.Sprintf("An unexpected error occurred. Please contact the support."))
	return &sourceErr{err, true}
}

func (instance *sourceErr) Err() string {
	return instance.err.Error()
}

func (instance *sourceErr) IsInternal() bool {
	return instance.isInternal
}

func (instance *sourceErr) Native() error {
	return instance.err
}
