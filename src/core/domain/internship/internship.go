package internship

import (
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/internshipLocation"
	"eletronic_point/src/core/messages"
	"time"
)

var _ Internship = &internship{}

type internship struct {
	startedIn time.Time
	endedIn   *time.Time
	location  internshipLocation.InternshipLocation
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
