package rules

import (
	"backend_template/src/core/domain/errors"
	"fmt"
	"reflect"
)

func NewRequiredRule(argument string) *Rule {
	return &Rule{
		Type:        REQUIRED,
		validator:   validateExists,
		Description: "verify if a value exists",
		argument:    argument,
	}
}

func validateExists(value interface{}) bool {
	fieldType := reflect.TypeOf(value)
	if fieldType == nil {
		return false
	}
	switch fieldType.Kind() {
	case reflect.Slice:
		return reflect.ValueOf(value).Len() > 0
	case reflect.Map:
		return false
	case reflect.String:
		return value.(string) != ""
	}
	return true
}

func NewRequiredError(fieldName, fieldType string) errors.Error {
	message := fmt.Sprintf("'%s' field of type '%s' is missing or empty", fieldName, fieldType)
	return newError(message)
}
