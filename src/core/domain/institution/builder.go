package institution

import (
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/messages"
	"strings"

	"github.com/google/uuid"
)

type builder struct {
	fields        []string
	errorMessages []string
	institution   *institution
}

func NewBuilder() *builder {
	return &builder{
		fields:        []string{},
		errorMessages: []string{},
		institution:   &institution{},
	}
}

func (this *builder) WithID(id uuid.UUID) *builder {
	if id == uuid.Nil {
		this.fields = append(this.fields, messages.InstitutionID)
		this.errorMessages = append(this.errorMessages, messages.InstitutionIDErrorMessage)
		return this
	}
	this.institution.id = id
	return this
}

func (this *builder) WithName(name string) *builder {
	name = strings.TrimSpace(name)
	if len(name) == 0 {
		this.fields = append(this.fields, messages.InstitutionName)
		this.errorMessages = append(this.errorMessages, messages.InstitutionNameErrorMessage)
		return this
	}
	this.institution.name = name
	return this
}

func (this *builder) Build() (*institution, errors.Error) {
	if len(this.errorMessages) > 0 {
		return nil, errors.NewValidationWithMetadata(this.errorMessages, map[string]interface{}{
			messages.Fields: this.fields,
		})
	}
	return this.institution, nil
}
