package timeRecord

import (
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/timeRecordStatus"
	"eletronic_point/src/core/messages"
	"time"

	"github.com/google/uuid"
)

var _ TimeRecord = &timeRecord{}

type timeRecord struct {
	id                uuid.UUID
	date              time.Time
	entryTime         time.Time
	exitTime          *time.Time
	location          string
	isOffSite         bool
	justification     *string
	studentID         uuid.UUID
	_timeRecordStatus timeRecordStatus.TimeRecordStatus
}

func (t *timeRecord) ID() uuid.UUID {
	return t.id
}

func (t *timeRecord) Date() time.Time {
	return t.date
}

func (t *timeRecord) EntryTime() time.Time {
	return t.entryTime
}

func (t *timeRecord) ExitTime() *time.Time {
	return t.exitTime
}

func (t *timeRecord) Location() string {
	return t.location
}

func (t *timeRecord) IsOffSite() bool {
	return t.isOffSite
}

func (t *timeRecord) Justification() *string {
	return t.justification
}

func (t *timeRecord) StudentID() uuid.UUID {
	return t.studentID
}

func (t *timeRecord) TimeRecordStatus() timeRecordStatus.TimeRecordStatus {
	return t._timeRecordStatus
}

func (t *timeRecord) SetID(id uuid.UUID) errors.Error {
	if id == uuid.Nil {
		return errors.NewFromString(messages.TimeRecordIDErrorMessage)
	}
	t.id = id
	return nil
}

func (t *timeRecord) SetDate(date time.Time) errors.Error {
	if date.IsZero() {
		return errors.NewFromString(messages.TimeRecordDateErrorMessage)
	}
	t.date = date
	return nil
}

func (t *timeRecord) SetEntryTime(entryTime time.Time) errors.Error {
	if entryTime.IsZero() {
		return errors.NewFromString(messages.TimeRecordEntryTimeErrorMessage)
	}
	t.entryTime = entryTime
	return nil
}

func (t *timeRecord) SetExitTime(exitTime *time.Time) errors.Error {
	t.exitTime = exitTime
	return nil
}

func (t *timeRecord) SetLocation(location string) errors.Error {
	if location == "" {
		return errors.NewFromString(messages.TimeRecordLocationErrorMessage)
	}
	t.location = location
	return nil
}

func (t *timeRecord) SetIsOffSite(isOffSite bool) errors.Error {
	t.isOffSite = isOffSite
	return nil
}

func (t *timeRecord) SetJustification(justification *string) errors.Error {
	t.justification = justification
	return nil
}

func (t *timeRecord) SetStudentID(studentID uuid.UUID) errors.Error {
	if studentID == uuid.Nil {
		return errors.NewFromString(messages.TimeRecordStudentIDErrorMessage)
	}
	t.studentID = studentID
	return nil
}

func (t *timeRecord) SetTimeRecordStatus(_timeRecordStatus timeRecordStatus.TimeRecordStatus) errors.Error {
	if _timeRecordStatus == nil {
		return errors.NewFromString(messages.TimeRecordStatusErrorMessage)
	}
	t._timeRecordStatus = _timeRecordStatus
	return nil
}
