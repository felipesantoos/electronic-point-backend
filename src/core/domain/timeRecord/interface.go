package timeRecord

import (
	"backend_template/src/core/domain/errors"
	"time"

	"github.com/google/uuid"
)

type TimeRecord interface {
	ID() uuid.UUID
	Date() time.Time
	EntryTime() time.Time
	ExitTime() *time.Time
	Location() string
	IsOffSite() bool
	Justification() *string
	StudentID() uuid.UUID

	SetID(uuid.UUID) errors.Error
	SetDate(time.Time) errors.Error
	SetEntryTime(time.Time) errors.Error
	SetExitTime(*time.Time) errors.Error
	SetLocation(string) errors.Error
	SetIsOffSite(bool) errors.Error
	SetJustification(*string) errors.Error
	SetStudentID(uuid.UUID) errors.Error
}
