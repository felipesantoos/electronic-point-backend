package internship

import (
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/internshipLocation"
	"eletronic_point/src/core/domain/simplifiedStudent"
	"time"

	"github.com/google/uuid"
)

type Internship interface {
	ID() uuid.UUID
	StartedIn() time.Time
	EndedIn() *time.Time
	Location() internshipLocation.InternshipLocation
	Student() simplifiedStudent.SimplifiedStudent

	SetID(uuid.UUID) errors.Error
	SetStartedIn(time.Time) errors.Error
	SetEndedIn(*time.Time) errors.Error
	SetLocation(internshipLocation.InternshipLocation) errors.Error
	SetStudent(simplifiedStudent.SimplifiedStudent) errors.Error
}
