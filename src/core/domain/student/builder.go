package student

import (
	"eletronic_point/src/core/domain/campus"
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/institution"
	"eletronic_point/src/core/domain/internship"
	"eletronic_point/src/core/domain/person"
	"eletronic_point/src/core/domain/timeRecord"
	"eletronic_point/src/core/messages"
	"strings"
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

func (builder *builder) WithPerson(_person person.Person) *builder {
	builder.student.Person = _person
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

func (builder *builder) WithInstitution(_institution institution.Institution) *builder {
	if _institution == nil {
		builder.fields = append(builder.fields, messages.StudentInstitution)
		builder.errorMessages = append(builder.errorMessages, messages.StudentInstitutionErrorMessage)
		return builder
	}
	builder.student._institution = _institution
	return builder
}

func (builder *builder) WithCampus(_campus campus.Campus) *builder {
	if _campus == nil {
		builder.fields = append(builder.fields, messages.StudentCampus)
		builder.errorMessages = append(builder.errorMessages, messages.StudentCampusErrorMessage)
		return builder
	}
	builder.student._campus = _campus
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

func (builder *builder) WithCurrentInternship(currentInternship internship.Internship) *builder {
	if currentInternship == nil {
		builder.fields = append(builder.fields, messages.Internship)
		builder.errorMessages = append(builder.errorMessages, messages.InternshipErrorMessage)
		return builder
	}
	builder.student.currentInternship = currentInternship
	return builder
}

func (builder *builder) WithInternshipHistory(internshipHistory []internship.Internship) *builder {
	builder.student.internshipHistory = internshipHistory
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
