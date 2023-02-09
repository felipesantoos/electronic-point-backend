package response

import (
	"dit_backend/src/core/domain/errors"
	"regexp"
	"strings"
)

var validationErrorRegexCompiler = regexp.MustCompile(`^('.+?') (.*)`)

type ErrorMessage struct {
	StatusCode    int            `json:"status_code,omitempty"`
	Message       string         `json:"message"`
	InvalidFields []InvalidField `json:"invalid_fields,omitempty"`
	isInternal    bool
}

type InvalidField struct {
	FieldName   string `json:"field_name"`
	Description string `json:"description"`
}

func NewErrorFromCore(err errors.Error, statusCode int) *ErrorMessage {
	errorMessage := &ErrorMessage{
		StatusCode: statusCode,
		Message:    strings.Join(err.Errors(), ", "),
	}
	if err.CausedByValidation() {
		for _, message := range err.Errors() {
			matches := validationErrorRegexCompiler.FindStringSubmatch(message)
			if len(matches) != 3 {
				continue
			}
			errorMessage.InvalidFields = append(errorMessage.InvalidFields, InvalidField{
				FieldName:   strings.ReplaceAll(matches[1], "'", ""),
				Description: matches[2],
			})
		}
	}
	return errorMessage
}

func (instance *ErrorMessage) IsInternal() bool {
	return instance.isInternal
}
