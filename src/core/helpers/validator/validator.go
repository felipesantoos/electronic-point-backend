package validator

import (
	"dit_backend/src/core"
	"dit_backend/src/core/domain/errors"
	"dit_backend/src/core/helpers/validator/rules"
	"dit_backend/src/core/utils"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type ValidateFunc func(interface{}) bool

var logger = core.Logger()

type Field struct {
	Name                  string
	Value                 interface{}
	Type                  string
	Rules                 []*rules.Rule
	ValidateIfValueExists bool
}

func (instance *Field) IsValid() ([]errors.Error, bool) {
	var errs []errors.Error
	for _, rule := range instance.Rules {
		if !rule.IsValid(instance.Value) {
			errs = append(errs, rule.Error(instance.Name))
			break
		}
	}
	return errs, len(errs) == 0
}

const (
	fieldDelimiter = "."
	ifExistsRule   = "ifExists"
)

// Rule Compilers
var lengthRuleCompiler = regexp.MustCompile(`len=(\d+)`)
var minLengthRuleCompiler = regexp.MustCompile(`minlen=(\d+)`)
var maxLengthRuleCompiler = regexp.MustCompile(`maxlen=(\d+)`)

// TODO: add struct array validation
func ValidateDTO[T interface{}](data interface{}) ([]string, bool) {
	var instance T
	var validators []*Field = buildValidatorsByStruct(instance)
	var errs []errors.Error = tryValidators(data, validators)
	var groupedErrs = groupErrors(errs)
	if len(groupedErrs) > 0 {
		logger.Error().Msg("Validation Error: " + strings.Join(groupedErrs, ", "))
	}
	return groupedErrs, len(groupedErrs) == 0
}

func buildValidatorsByStruct(instance interface{}) []*Field {
	var reflection = reflect.ValueOf(instance)
	if reflection.Kind() == reflect.Ptr {
		reflection = reflection.Elem()
	}
	var fields []*Field
	for i := 0; i < reflection.NumField(); i++ {
		fieldValue := reflection.Field(i)
		fieldType := reflection.Type().Field(i)
		if fieldProps := newFieldValidatorByField(fieldType, fieldValue); fieldProps != nil {
			fieldProps.ValidateIfValueExists = strings.Contains(fieldType.Tag.Get("validate"), ifExistsRule)
			fields = append(fields, fieldProps)
		}
		if fieldValue.Kind() == reflect.Struct {
			fields = append(fields, buildNestedFields(fieldType, fieldValue)...)
		}
	}
	return fields
}

func newFieldValidatorByField(fieldType reflect.StructField, fieldValue reflect.Value) *Field {
	if fieldType.Tag.Get("novalidate") == "true" || fieldType.Type.Kind() == reflect.Struct {
		return nil
	}
	typeName := fieldType.Type.Kind().String()
	if fieldType.Type.Kind() == reflect.Slice {
		typeName = fmt.Sprintf("[]%s", fieldType.Type.Elem().Name())
	}

	return &Field{
		Name:  extractFieldName(fieldType),
		Type:  typeName,
		Rules: getValidatorsByFieldTag(fieldType, fieldValue),
	}
}

func buildNestedFields(fieldType reflect.StructField, fieldValue reflect.Value) []*Field {
	fieldName := extractFieldName(fieldType)
	nestedValidators := buildValidatorsByStruct(fieldValue.Interface())
	validNestedFields := strings.Split(fieldType.Tag.Get("nestedProps"), ",")
	validNestedValidators := []*Field{}
	validateIfHasValue := strings.Contains(fieldType.Tag.Get("validate"), ifExistsRule)
	if fieldType.Tag.Get("nestedProps") != "" && len(validNestedFields) > 0 {
		for _, prop := range validNestedFields {
			for _, validator := range nestedValidators {
				if validator.Name == prop {
					validator.ValidateIfValueExists = validateIfHasValue
					validNestedValidators = append(validNestedValidators, validator)
					break
				}
			}
		}
	} else {
		validNestedValidators = nestedValidators
	}
	mustHideParentName := fieldType.Tag.Get("hideParentName") == "true"
	for _, validator := range validNestedValidators {
		if mustHideParentName {
			continue
		}
		delimiter := fieldDelimiter
		for _, fieldName := range validNestedFields {
			if validator.Name == fieldName {
				delimiter = "_"
				break
			}
		}
		validator.Name = fmt.Sprintf("%s%s%s", strings.ToLower(fieldName), delimiter, validator.Name)
	}
	return validNestedValidators
}

func extractFieldName(fieldType reflect.StructField) string {
	return strings.Split(fieldType.Tag.Get("json"), ",")[0]
}

func getValidatorsByFieldTag(fieldType reflect.StructField, fieldValue reflect.Value) []*rules.Rule {
	typeName := fieldType.Type.Kind().String()
	if strings.Contains(fieldValue.String(), "uuid") {
		typeName = "string UUID"
	}
	hints := strings.Split(fieldType.Tag.Get("validate"), ",")
	var validators []*rules.Rule
	if !strings.Contains(fieldType.Tag.Get("json"), "omitempty") {
		validators = append(validators, rules.NewRequiredRule(typeName))
	}
	if validator := rules.GetConvertionValidator(typeName); validator != nil {
		validators = append(validators, validator)
	}
	for _, hint := range hints {
		validator := getRuleByHint(hint)
		if validator != nil {
			validators = append(validators, validator)
		}
	}
	return validators
}

func getRuleByHint(hint string) *rules.Rule {
	var ruleBuilder func(length int) *rules.Rule
	var matchValue string
	if lengthRuleCompiler.Match([]byte(hint)) {
		ruleBuilder = rules.NewLengthRule
		matchValue = lengthRuleCompiler.FindStringSubmatch(hint)[1]
	} else if minLengthRuleCompiler.Match([]byte(hint)) {
		ruleBuilder = rules.NewMinLengthRule
		matchValue = minLengthRuleCompiler.FindStringSubmatch(hint)[1]
	} else if maxLengthRuleCompiler.Match([]byte(hint)) {
		ruleBuilder = rules.NewMaxLengthRule
		matchValue = maxLengthRuleCompiler.FindStringSubmatch(hint)[1]
	}
	if ruleBuilder == nil {
		return nil
	}
	value, err := strconv.Atoi(matchValue)
	if err != nil {
		value = 0
	}
	return ruleBuilder(value)
}

func tryValidators(data interface{}, fields []*Field) []errors.Error {
	var errs []errors.Error
	var formattedData map[string]interface{} = utils.FormatJSONData(data)
	for _, field := range fields {
		field.Value = getFieldValue(field.Name, formattedData)
		if field.ValidateIfValueExists && field.Value == nil {
			continue
		}
		if fieldErrs, ok := field.IsValid(); !ok {
			errs = append(errs, fieldErrs...)
		}
	}
	return errs
}

func getFieldValue(key string, data map[string]interface{}) interface{} {
	if !strings.Contains(key, fieldDelimiter) {
		return data[key]
	}
	var value interface{} = data
	for _, key := range strings.Split(key, fieldDelimiter) {
		value, exists := value.(map[string]interface{})[key]
		if !exists || value == nil {
			return nil
		}
	}
	return value
}

func groupErrors(errs []errors.Error) []string {
	var messages []string
	for _, err := range errs {
		messages = append(messages, err.Errors()...)
	}
	return messages
}
