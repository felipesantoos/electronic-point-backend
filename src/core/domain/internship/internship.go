package internship

import (
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/internshipLocation"
	"eletronic_point/src/core/messages"
	"eletronic_point/src/utils/validator"
	"time"

	"github.com/google/uuid"
)

var _ Internship = &internship{}

type internship struct {
	id        uuid.UUID
	startedIn time.Time
	endedIn   *time.Time
	location  internshipLocation.InternshipLocation
	studentID uuid.UUID
}

func (i *internship) ID() uuid.UUID {
	return i.id
}

func (i *internship) StartedIn() time.Time {
	return i.startedIn
}

func (i *internship) EndedIn() *time.Time {
	return i.endedIn
}

func (i *internship) Location() internshipLocation.InternshipLocation {
	return i.location
}

func (i *internship) StudentID() uuid.UUID {
	return i.studentID
}

func (i *internship) SetID(id uuid.UUID) errors.Error {
	if !validator.IsUUIDValid(id) {
		return errors.NewValidationFromString(messages.InternshipIDErrorMessage)
	}
	i.id = id
	return nil
}

func (i *internship) SetStartedIn(startedIn time.Time) errors.Error {
	if startedIn.IsZero() {
		return errors.NewValidationFromString(messages.InternshipStartedInErrorMessage)
	}
	i.startedIn = startedIn
	return nil
}

func (i *internship) SetEndedIn(endedIn *time.Time) errors.Error {
	i.endedIn = endedIn
	return nil
}

func (i *internship) SetLocation(location internshipLocation.InternshipLocation) errors.Error {
	if location == nil {
		return errors.NewValidationFromString(messages.InternshipLocationErrorMessage)
	}
	i.location = location
	return nil
}

func (i *internship) SetStudentID(studentID uuid.UUID) errors.Error {
	if !validator.IsUUIDValid(studentID) {
		return errors.NewValidationFromString(messages.StudentIDErrorMessage)
	}
	i.studentID = studentID
	return nil
}
