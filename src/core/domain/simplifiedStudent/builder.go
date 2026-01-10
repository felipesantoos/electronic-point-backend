package simplifiedStudent

import (
	"eletronic_point/src/core/domain/campus"
	"eletronic_point/src/core/domain/course"
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/institution"
	"eletronic_point/src/core/messages"
	"eletronic_point/src/utils/validator"
	"strings"

	"github.com/google/uuid"
)

type builder struct {
	fields            []string
	errorMessages     []string
	simplifiedStudent *simplifiedStudent
}

func NewBuilder() *builder {
	return &builder{
		fields:            []string{},
		errorMessages:     []string{},
		simplifiedStudent: &simplifiedStudent{},
	}
}

func (builder *builder) WithID(id uuid.UUID) *builder {
	if !validator.IsUUIDValid(id) {
		builder.fields = append(builder.fields, messages.StudentID)
		builder.errorMessages = append(builder.errorMessages, messages.StudentIDErrorMessage)
		return builder
	}

	builder.simplifiedStudent.id = &id
	return builder
}

func (builder *builder) WithName(name string) *builder {
	name = strings.TrimSpace(name)
	if len(name) == 0 || len(strings.Split(name, " ")) == 1 {
		builder.fields = append(builder.fields, messages.StudentName)
		builder.errorMessages = append(builder.errorMessages, messages.StudentNameErrorMessage)
		return builder
	}

	builder.simplifiedStudent.name = name
	return builder
}

func (builder *builder) WithProfilePicture(profilePicture *string) *builder {
	builder.simplifiedStudent.profilePicture = profilePicture
	return builder
}

func (builder *builder) WithInstitution(_institution institution.Institution) *builder {
	if _institution == nil {
		builder.fields = append(builder.fields, messages.StudentInstitution)
		builder.errorMessages = append(builder.errorMessages, messages.StudentInstitutionErrorMessage)
		return builder
	}
	builder.simplifiedStudent._institution = _institution
	return builder
}

func (builder *builder) WithCampus(_campus campus.Campus) *builder {
	if _campus == nil {
		builder.fields = append(builder.fields, messages.StudentCampus)
		builder.errorMessages = append(builder.errorMessages, messages.StudentCampusErrorMessage)
		return builder
	}
	builder.simplifiedStudent._campus = _campus
	return builder
}

func (builder *builder) WithCourse(_course course.Course) *builder {
	if _course == nil {
		builder.fields = append(builder.fields, messages.StudentCourse)
		builder.errorMessages = append(builder.errorMessages, messages.StudentCourseErrorMessage)
		return builder
	}
	builder.simplifiedStudent._course = _course
	return builder
}

func (builder *builder) WithTotalWorkload(totalWorkload int) *builder {
	builder.simplifiedStudent.totalWorkload = totalWorkload
	return builder
}

func (builder *builder) Build() (SimplifiedStudent, errors.Error) {
	if len(builder.errorMessages) > 0 {
		return nil, errors.NewValidationWithMetadata(builder.errorMessages, map[string]interface{}{
			messages.Fields: builder.fields})
	}
	return builder.simplifiedStudent, nil
}
