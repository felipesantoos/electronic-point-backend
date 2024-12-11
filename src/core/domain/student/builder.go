package student

import (
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/timeRecord"
	"eletronic_point/src/core/messages"
	"eletronic_point/src/utils/validator"
	"strings"

	"github.com/google/uuid"
)

type builder struct {
	fields        []string
	errorMessages []string
	student       *student
}

func NewBuilder() *builder {
	return &builder{
		fields:        []string{},
		errorMessages: []string{},
		student:       &student{},
	}
}

func (builder *builder) WithID(id uuid.UUID) *builder {
	if !validator.IsUUIDValid(id) {
		builder.fields = append(builder.fields, messages.StudentID)
		builder.errorMessages = append(builder.errorMessages, messages.StudentIDErrorMessage)
		return builder
	}

	builder.student.id = id
	return builder
}

func (builder *builder) WithName(name string) *builder {
	name = strings.TrimSpace(name)
	if len(name) == 0 {
		builder.fields = append(builder.fields, messages.StudentName)
		builder.errorMessages = append(builder.errorMessages, messages.StudentNameErrorMessage)
		return builder
	}

	builder.student.name = name
	return builder
}

func (builder *builder) WithRegistration(registration string) *builder {
	registration = strings.TrimSpace(registration)
	if len(registration) == 0 {
		builder.fields = append(builder.fields, messages.StudentRegistration)
		builder.errorMessages = append(builder.errorMessages, messages.StudentRegistrationErrorMessage)
		return builder
	}

	builder.student.registration = registration
	return builder
}

func (builder *builder) WithProfilePicture(profilePicture *string) *builder {
	builder.student.profilePicture = profilePicture
	return builder
}

func (builder *builder) WithInstitution(institution string) *builder {
	institution = strings.TrimSpace(institution)
	if len(institution) == 0 {
		builder.fields = append(builder.fields, messages.StudentInstitution)
		builder.errorMessages = append(builder.errorMessages, messages.StudentInstitutionErrorMessage)
		return builder
	}

	builder.student.institution = institution
	return builder
}

func (builder *builder) WithCourse(course string) *builder {
	course = strings.TrimSpace(course)
	if len(course) == 0 {
		builder.fields = append(builder.fields, messages.StudentCourse)
		builder.errorMessages = append(builder.errorMessages, messages.StudentCourseErrorMessage)
		return builder
	}

	builder.student.course = course
	return builder
}

func (builder *builder) WithInternshipLocationName(locationName string) *builder {
	locationName = strings.TrimSpace(locationName)
	if len(locationName) == 0 {
		builder.fields = append(builder.fields, messages.StudentInternshipLocationName)
		builder.errorMessages = append(builder.errorMessages, messages.StudentInternshipLocationNameErrorMessage)
		return builder
	}

	builder.student.internshipLocationName = locationName
	return builder
}

func (builder *builder) WithInternshipAddress(address string) *builder {
	address = strings.TrimSpace(address)
	if len(address) == 0 {
		builder.fields = append(builder.fields, messages.StudentInternshipAddress)
		builder.errorMessages = append(builder.errorMessages, messages.StudentInternshipAddressErrorMessage)
		return builder
	}

	builder.student.internshipAddress = address
	return builder
}

func (builder *builder) WithInternshipLocation(location string) *builder {
	location = strings.TrimSpace(location)
	if len(location) == 0 {
		builder.fields = append(builder.fields, messages.StudentInternshipLocation)
		builder.errorMessages = append(builder.errorMessages, messages.StudentInternshipLocationErrorMessage)
		return builder
	}

	builder.student.internshipLocation = location
	return builder
}

func (builder *builder) WithTotalWorkload(totalWorkload int) *builder {
	if totalWorkload < 0 {
		builder.fields = append(builder.fields, messages.StudentTotalWorkload)
		builder.errorMessages = append(builder.errorMessages, messages.StudentTotalWorkloadErrorMessage)
		return builder
	}

	builder.student.totalWorkload = totalWorkload
	return builder
}

func (builder *builder) WithWorkloadCompleted(workloadCompleted int) *builder {
	if workloadCompleted < 0 {
		builder.fields = append(builder.fields, messages.StudentWorkloadCompleted)
		builder.errorMessages = append(builder.errorMessages, messages.StudentWorkloadCompletedErrorMessage)
		return builder
	}

	builder.student.workloadCompleted = workloadCompleted
	return builder
}

func (builder *builder) WithPendingWorkload(pendingWorkload int) *builder {
	if pendingWorkload < 0 {
		builder.fields = append(builder.fields, messages.StudentPendingWorkload)
		builder.errorMessages = append(builder.errorMessages, messages.StudentPendingWorkloadErrorMessage)
		return builder
	}

	builder.student.pendingWorkload = pendingWorkload
	return builder
}

func (builder *builder) WithFrequencyHistory(frequencyHistory []timeRecord.TimeRecord) *builder {
	builder.student.frequencyHistory = frequencyHistory
	return builder
}

func (builder *builder) Build() (Student, errors.Error) {
	if len(builder.errorMessages) > 0 {
		return nil, errors.NewValidationWithMetadata(builder.errorMessages, map[string]interface{}{
			messages.Fields: builder.fields})
	}
	return builder.student, nil
}
