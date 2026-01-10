package services

import (
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/timeRecord"
	"eletronic_point/src/core/domain/timeRecordStatus"
	"eletronic_point/src/core/interfaces/primary"
	"eletronic_point/src/core/interfaces/secondary"
	"eletronic_point/src/core/messages"
	"eletronic_point/src/core/services/filters"
	"math"
	"time"

	"github.com/google/uuid"
)

type timeRecordServices struct {
	repository           secondary.TimeRecordPort
	internshipRepository secondary.InternshipPort
}

func NewTimeRecordServices(repository secondary.TimeRecordPort, internshipRepository secondary.InternshipPort) primary.TimeRecordPort {
	return &timeRecordServices{repository, internshipRepository}
}

func (this *timeRecordServices) Create(_timeRecord timeRecord.TimeRecord) (*uuid.UUID, errors.Error) {
	studentID := _timeRecord.StudentID()
	internshipID := _timeRecord.InternshipID()
	date := _timeRecord.Date()

	// 1. Get student's specific internship
	internship, err := this.internshipRepository.Get(internshipID)
	if err != nil {
		return nil, err
	}

	// 2. Validate ownership: does this internship belong to the student?
	if internship.Student().ID() != nil && *internship.Student().ID() != studentID {
		return nil, errors.NewValidationFromString(messages.InternshipNotFoundErrorMessage)
	}

	// Associate the internship object with the time record
	_timeRecord.SetInternship(internship)

	// 3. Tolerance check (30 min) using the specific internship
	scheduleEntry := _timeRecord.Internship().ScheduleEntryTime()
	if scheduleEntry != nil {
		entry := _timeRecord.EntryTime()
		diff := math.Abs(float64(entry.Hour()*60 + entry.Minute() - (scheduleEntry.Hour()*60 + scheduleEntry.Minute())))
		if diff > 30 {
			return nil, errors.NewValidationFromString(messages.TimeRecordToleranceErrorMessage)
		}
	}

	// 3. Daily cumulative limit (5h)
	existingRecords, err := this.repository.List(filters.TimeRecordFilters{
		StudentID: &studentID,
		StartDate: &date,
		EndDate:   &date,
	})
	if err != nil {
		return nil, err
	}

	var totalMinutes float64
	for _, record := range existingRecords {
		if record.ExitTime() != nil {
			totalMinutes += record.ExitTime().Sub(record.EntryTime()).Minutes()
		}
	}

	const maxDailyMinutes = 5 * 60.0
	if totalMinutes >= maxDailyMinutes {
		return nil, errors.NewValidationFromString(messages.TimeRecordDailyLimitErrorMessage)
	}

	if _timeRecord.ExitTime() != nil {
		newRecordMinutes := _timeRecord.ExitTime().Sub(_timeRecord.EntryTime()).Minutes()
		if totalMinutes+newRecordMinutes > maxDailyMinutes {
			// Cut the excess
			allowedMinutes := maxDailyMinutes - totalMinutes
			newExitTime := _timeRecord.EntryTime().Add(time.Duration(allowedMinutes) * time.Minute)
			_timeRecord.SetExitTime(&newExitTime)
		}
	}

	return this.repository.Create(_timeRecord)
}

func (this *timeRecordServices) Update(_timeRecord timeRecord.TimeRecord) errors.Error {
	studentID := _timeRecord.StudentID()
	internshipID := _timeRecord.InternshipID()
	date := _timeRecord.Date()

	// 1. Get student's specific internship
	internship, err := this.internshipRepository.Get(internshipID)
	if err != nil {
		return err
	}

	// 2. Validate ownership
	if internship.Student().ID() != nil && *internship.Student().ID() != studentID {
		return errors.NewValidationFromString(messages.InternshipNotFoundErrorMessage)
	}

	_timeRecord.SetInternship(internship)

	// Daily cumulative limit (5h) check on update as well
	existingRecords, err := this.repository.List(filters.TimeRecordFilters{
		StudentID: &studentID,
		StartDate: &date,
		EndDate:   &date,
	})
	if err != nil {
		return err
	}

	var totalMinutes float64
	for _, record := range existingRecords {
		if record.ID() == _timeRecord.ID() {
			continue // Skip current record
		}
		if record.ExitTime() != nil {
			totalMinutes += record.ExitTime().Sub(record.EntryTime()).Minutes()
		}
	}

	const maxDailyMinutes = 5 * 60.0
	if totalMinutes >= maxDailyMinutes {
		return errors.NewValidationFromString(messages.TimeRecordDailyLimitErrorMessage)
	}

	if _timeRecord.ExitTime() != nil {
		newRecordMinutes := _timeRecord.ExitTime().Sub(_timeRecord.EntryTime()).Minutes()
		if totalMinutes+newRecordMinutes > maxDailyMinutes {
			// Cut the excess
			allowedMinutes := maxDailyMinutes - totalMinutes
			newExitTime := _timeRecord.EntryTime().Add(time.Duration(allowedMinutes) * time.Minute)
			_timeRecord.SetExitTime(&newExitTime)
		}
	}

	return this.repository.Update(_timeRecord)
}

func (this *timeRecordServices) Delete(id uuid.UUID) errors.Error {
	return this.repository.Delete(id)
}

func (this *timeRecordServices) List(_filters filters.TimeRecordFilters) ([]timeRecord.TimeRecord, errors.Error) {
	return this.repository.List(_filters)
}

func (this *timeRecordServices) Get(id uuid.UUID, _filters filters.TimeRecordFilters) (timeRecord.TimeRecord, errors.Error) {
	return this.repository.Get(id, _filters)
}

func (this *timeRecordServices) Approve(timeRecordID uuid.UUID, approvedBy uuid.UUID) errors.Error {
	return this.repository.UpdateStatus(timeRecordID, approvedBy, timeRecordStatus.Approved.ID())
}

func (this *timeRecordServices) Disapprove(timeRecordID uuid.UUID, disapprovedBy uuid.UUID) errors.Error {
	return this.repository.UpdateStatus(timeRecordID, disapprovedBy, timeRecordStatus.Disapproved.ID())
}
