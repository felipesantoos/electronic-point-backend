package rules

import (
	"dit_backend/src/core/domain/errors"
	"fmt"
	"reflect"
)

var convertionValidator = map[string]*Rule{
	"int":      NewConvertionRule[float64](),
	"int32":    NewConvertionRule[float64](),
	"int64":    NewConvertionRule[float64](),
	"float32":  NewConvertionRule[float64](),
	"float64":  NewConvertionRule[float64](),
	"string":   NewConvertionRule[string](),
	"bool":     NewConvertionRule[bool](),
	"[]int":    NewSliceRule[int](),
	"[]string": NewSliceRule[string](),
}

func NewConvertionRuleWithMethod(method ValidatorFunc) *Rule {
	return &Rule{
		Type:        CONVERTION,
		Description: "verify if a value is convertable to int",
		validator:   method,
		argument:    "int",
	}
}

func GetType[T comparable]() string {
	var t T
	return fmt.Sprintf("%T", t)
}

func NewConvertionRule[T comparable]() *Rule {
	typeName := GetType[T]()
	return &Rule{
		Type:        CONVERTION,
		Description: fmt.Sprintf("verify if a value is convertable to %s", typeName),
		validator:   validateConvertion[T],
		argument:    typeName,
	}
}

func NewSliceRule[T comparable]() *Rule {
	typeName := GetType[T]()
	return &Rule{
		Type:        CONVERTION,
		Description: fmt.Sprintf("verify if a value is convertable to %s array", typeName),
		validator:   validateSliceConvertion[T],
		argument:    typeName,
	}
}

func GetConvertionValidator(key string) *Rule {
	return convertionValidator[key]
}

func NewWrongConvertionError(fieldName, fieldType string) errors.Error {
	if fieldType == "struct" {
		fieldType = "json"
	}
	message := fmt.Sprintf("'%s' field type must be '%s'", fieldName, fieldType)
	return newError(message)
}

func validateConvertion[T comparable](value interface{}) bool {
	var t T
	neededType := reflect.TypeOf(t)
	valueType := reflect.TypeOf(value)
	return neededType == valueType
}

func validateSliceConvertion[T comparable](value interface{}) bool {
	newValue, ok := value.([]interface{})
	for _, item := range newValue {
		_, ok = item.(T)
		if !ok {
			return false
		}
	}
	return ok
}
