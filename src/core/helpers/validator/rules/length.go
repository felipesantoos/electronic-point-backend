package rules

import (
	"dit_backend/src/core/domain/errors"
	"fmt"
)

func NewMinLengthRule(length int) *Rule {
	return &Rule{
		Type:        MINLENGTH,
		Description: fmt.Sprintf("verify if a value has a minimum length of %d", length),
		validator:   getMinLengthValidator(length),
		argument:    fmt.Sprint(length),
	}
}

func NewMaxLengthRule(length int) *Rule {
	return &Rule{
		Type:        MAXLENGTH,
		Description: fmt.Sprintf("verify if a value has a maximum length of %d", length),
		validator:   getMaxLengthValidator(length),
		argument:    fmt.Sprint(length),
	}
}

func NewLengthRule(length int) *Rule {
	return &Rule{
		Type:        LENGTH,
		Description: fmt.Sprintf("verify if a value has length equals to %d", length),
		validator:   getLengthValidator(length),
		argument:    fmt.Sprint(length),
	}
}

func NewLengthError(fieldName string, length string) errors.Error {
	message := fmt.Sprintf("'%s' field must have '%s' length", fieldName, length)
	return newError(message)
}

func NewMinLengthError(fieldName string, length string) errors.Error {
	message := fmt.Sprintf("'%s' field must have '%s' min length", fieldName, length)
	return newError(message)
}

func NewMaxLengthError(fieldName string, length string) errors.Error {
	message := fmt.Sprintf("'%s' field must have '%s' max length", fieldName, length)
	return newError(message)
}

func getMinLengthValidator(length int) ValidatorFunc {
	return func(value interface{}) bool {
		valueStr := value.(string)
		return len(valueStr) >= length
	}
}

func getMaxLengthValidator(length int) ValidatorFunc {
	return func(value interface{}) bool {
		valueStr := value.(string)
		return len(valueStr) <= length
	}
}

func getLengthValidator(length int) ValidatorFunc {
	return func(value interface{}) bool {
		valueStr := value.(string)
		return len(valueStr) == length
	}
}
