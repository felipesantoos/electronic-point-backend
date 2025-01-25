package timeRecordStatus

import (
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/messages"
	"strings"

	"github.com/google/uuid"
)

type builder struct {
	fields           []string
	errorMessages    []string
	timeRecordStatus *timeRecordStatus
}

func NewBuilder() *builder {
	return &builder{
		fields:           []string{},
		errorMessages:    []string{},
		timeRecordStatus: &timeRecordStatus{},
	}
}

func (this *builder) WithID(id uuid.UUID) *builder {
	if id == uuid.Nil {
		this.fields = append(this.fields, messages.TimeRecordStatusID)
		this.errorMessages = append(this.errorMessages, messages.TimeRecordStatusIDErrorMessage)
		return this
	}
	this.timeRecordStatus.id = id
	return this
}

func (this *builder) WithName(name string) *builder {
	name = strings.TrimSpace(name)
	if len(name) == 0 {
		this.fields = append(this.fields, messages.TimeRecordStatusName)
		this.errorMessages = append(this.errorMessages, messages.TimeRecordStatusNameErrorMessage)
		return this
	}
	this.timeRecordStatus.name = name
	return this
}

func (this *builder) Build() (*timeRecordStatus, errors.Error) {
	if len(this.errorMessages) > 0 {
		return nil, errors.NewValidationWithMetadata(this.errorMessages, map[string]interface{}{
			messages.Fields: this.fields,
		})
	}
	return this.timeRecordStatus, nil
}
