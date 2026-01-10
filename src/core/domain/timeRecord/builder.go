package timeRecord

import (
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/internship"
	"eletronic_point/src/core/domain/simplifiedStudent"
	"eletronic_point/src/core/domain/timeRecordStatus"
	"eletronic_point/src/core/messages"
	"eletronic_point/src/utils/validator"
	"strings"
	"time"

	"github.com/google/uuid"
)

type builder struct {
	fields        []string
	errorMessages []string
	timeRecord    *timeRecord
}

func NewBuilder() *builder {
	return &builder{
		fields:        []string{},
		errorMessages: []string{},
		timeRecord:    &timeRecord{},
	}
}

func (builder *builder) WithID(id uuid.UUID) *builder {
	if !validator.IsUUIDValid(id) {
		builder.fields = append(builder.fields, messages.TimeRecordID)
		builder.errorMessages = append(builder.errorMessages, messages.TimeRecordIDErrorMessage)
		return builder
	}

	builder.timeRecord.id = id
	return builder
}

func (builder *builder) WithDate(date time.Time) *builder {
	if date.IsZero() {
		builder.fields = append(builder.fields, messages.TimeRecordDate)
		builder.errorMessages = append(builder.errorMessages, messages.TimeRecordDateErrorMessage)
		return builder
	}

	builder.timeRecord.date = date
	return builder
}

func (builder *builder) WithEntryTime(entryTime time.Time) *builder {
	if entryTime.IsZero() {
		builder.fields = append(builder.fields, messages.TimeRecordEntryTime)
		builder.errorMessages = append(builder.errorMessages, messages.TimeRecordEntryTimeErrorMessage)
		return builder
	}

	builder.timeRecord.entryTime = entryTime
	return builder
}

func (builder *builder) WithExitTime(exitTime *time.Time) *builder {
	builder.timeRecord.exitTime = exitTime
	return builder
}

func (builder *builder) WithLocation(location string) *builder {
	location = strings.TrimSpace(location)
	if len(location) == 0 {
		builder.fields = append(builder.fields, messages.TimeRecordLocation)
		builder.errorMessages = append(builder.errorMessages, messages.TimeRecordLocationErrorMessage)
		return builder
	}

	builder.timeRecord.location = location
	return builder
}

func (builder *builder) WithIsOffSite(isOffSite bool) *builder {
	builder.timeRecord.isOffSite = isOffSite
	return builder
}

func (builder *builder) WithJustification(justification *string) *builder {
	builder.timeRecord.justification = justification
	return builder
}

func (builder *builder) WithStudentID(studentID uuid.UUID) *builder {
	if !validator.IsUUIDValid(studentID) {
		builder.fields = append(builder.fields, messages.TimeRecordStudentID)
		builder.errorMessages = append(builder.errorMessages, messages.TimeRecordStudentIDErrorMessage)
		return builder
	}

	builder.timeRecord.studentID = studentID
	return builder
}

func (builder *builder) WithInternshipID(internshipID uuid.UUID) *builder {
	if !validator.IsUUIDValid(internshipID) {
		builder.fields = append(builder.fields, messages.InternshipID)
		builder.errorMessages = append(builder.errorMessages, messages.InternshipIDErrorMessage)
		return builder
	}

	builder.timeRecord.internshipID = internshipID
	return builder
}

func (builder *builder) WithStudent(student simplifiedStudent.SimplifiedStudent) *builder {
	builder.timeRecord.student = student
	return builder
}

func (builder *builder) WithInternship(internship internship.Internship) *builder {
	builder.timeRecord.internship = internship
	return builder
}

func (builder *builder) WithTimeRecord(_timeRecordStatus timeRecordStatus.TimeRecordStatus) *builder {
	if _timeRecordStatus == nil {
		builder.fields = append(builder.fields, messages.TimeRecordStatus)
		builder.errorMessages = append(builder.errorMessages, messages.TimeRecordStatusErrorMessage)
		return builder
	}
	builder.timeRecord._timeRecordStatus = _timeRecordStatus
	return builder
}

func (builder *builder) Build() (TimeRecord, errors.Error) {
	if len(builder.errorMessages) > 0 {
		return nil, errors.NewValidationWithMetadata(builder.errorMessages, map[string]interface{}{
			messages.Fields: builder.fields})
	}
	return builder.timeRecord, nil
}
