package internship

import (
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/internshipLocation"
	"time"
)

type Internship interface {
	StartedIn() time.Time
	EndedIn() *time.Time
	Location() internshipLocation.InternshipLocation

	SetStartedIn(time.Time) errors.Error
	SetEndedIn(*time.Time) errors.Error
	SetLocation(internshipLocation.InternshipLocation) errors.Error
}
