package rules

import (
	"dit_backend/src/core/domain/errors"
	builtinErrors "errors"
)

type ValidatorFunc func(value interface{}) bool

type Rule struct {
	Type        int
	Description string
	validator   ValidatorFunc
	argument    string
}

func (instance *Rule) IsValid(value interface{}) bool {
	return instance.validator(value)
}

func (instance *Rule) Error(fieldName string) errors.Error {
	return newErrorByType(instance.Type, fieldName, instance.argument)
}

const (
	REQUIRED = iota
	CONVERTION
	LENGTH
	MINLENGTH
	MAXLENGTH
)

func newErrorByType(t int, fieldName, argument string) errors.Error {
	switch t {
	case REQUIRED:
		return NewRequiredError(fieldName, argument)
	case CONVERTION:
		return NewWrongConvertionError(fieldName, argument)
	case LENGTH:
		return NewLengthError(fieldName, argument)
	case MINLENGTH:
		return NewMinLengthError(fieldName, argument)
	case MAXLENGTH:
		return NewMaxLengthError(fieldName, argument)
	}
	return nil
}

func newError(message string) errors.Error {
	return errors.New(builtinErrors.New(message))
}
