package course

import (
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/messages"
	"strings"

	"github.com/google/uuid"
)

type builder struct {
	fields        []string
	errorMessages []string
	course        *course
}

func NewBuilder() *builder {
	return &builder{
		fields:        []string{},
		errorMessages: []string{},
		course:        &course{},
	}
}

func (this *builder) WithID(id uuid.UUID) *builder {
	if id == uuid.Nil {
		this.fields = append(this.fields, messages.CourseID)
		this.errorMessages = append(this.errorMessages, messages.CourseIDErrorMessage)
		return this
	}
	this.course.id = id
	return this
}

func (this *builder) WithName(name string) *builder {
	name = strings.TrimSpace(name)
	if len(name) == 0 {
		this.fields = append(this.fields, messages.CourseName)
		this.errorMessages = append(this.errorMessages, messages.CourseNameErrorMessage)
		return this
	}
	this.course.name = name
	return this
}

func (this *builder) Build() (*course, errors.Error) {
	if len(this.errorMessages) > 0 {
		return nil, errors.NewValidationWithMetadata(this.errorMessages, map[string]interface{}{
			messages.Fields: this.fields,
		})
	}
	return this.course, nil
}
