package campus

import (
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/messages"
	"strings"

	"github.com/google/uuid"
)

type builder struct {
	fields        []string
	errorMessages []string
	campus        *campus
}

func NewBuilder() *builder {
	return &builder{
		fields:        []string{},
		errorMessages: []string{},
		campus:        &campus{},
	}
}

func (this *builder) WithID(id uuid.UUID) *builder {
	if id == uuid.Nil {
		this.fields = append(this.fields, messages.CampusID)
		this.errorMessages = append(this.errorMessages, messages.CampusIDErrorMessage)
		return this
	}
	this.campus.id = id
	return this
}

func (this *builder) WithName(name string) *builder {
	name = strings.TrimSpace(name)
	if len(name) == 0 {
		this.fields = append(this.fields, messages.CampusName)
		this.errorMessages = append(this.errorMessages, messages.CampusNameErrorMessage)
		return this
	}
	this.campus.name = name
	return this
}

func (this *builder) WithInstitutionID(institutionID uuid.UUID) *builder {
	if institutionID == uuid.Nil {
		this.fields = append(this.fields, messages.CampusInstitutionID)
		this.errorMessages = append(this.errorMessages, messages.CampusInstitutionIDErrorMessage)
		return this
	}
	this.campus.institutionID = institutionID
	return this
}

func (this *builder) Build() (*campus, errors.Error) {
	if len(this.errorMessages) > 0 {
		return nil, errors.NewValidationWithMetadata(this.errorMessages, map[string]interface{}{
			messages.Fields: this.fields,
		})
	}
	return this.campus, nil
}
