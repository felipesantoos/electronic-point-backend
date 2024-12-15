package internship

import (
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/internshipLocation"
	"time"

	"github.com/google/uuid"
)

type Internship interface {
	ID() uuid.UUID
	StartedIn() time.Time
	EndedIn() *time.Time
	Location() internshipLocation.InternshipLocation
	StudentID() uuid.UUID

	SetID(uuid.UUID) errors.Error
	SetStartedIn(time.Time) errors.Error
	SetEndedIn(*time.Time) errors.Error
	SetLocation(internshipLocation.InternshipLocation) errors.Error
	SetStudentID(uuid.UUID) errors.Error
}
