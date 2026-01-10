package request

import (
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/internship"
	"eletronic_point/src/core/domain/internshipLocation"
	"eletronic_point/src/core/domain/simplifiedStudent"
	"time"

	"github.com/google/uuid"
)

type Internship struct {
	StudentID         uuid.UUID  `json:"student_id" example:"02e62826-bf41-4944-adb2-051b6a30a131"`
	LocationID        uuid.UUID  `json:"location_id" example:"0d0518be-bd9d-4ac4-9f95-6c536e00b247"`
	StartedIn         time.Time  `json:"started_in" example:"2024-06-01T00:00:00Z"`
	EndedIn           *time.Time `json:"ended_in" example:"null"`
	ScheduleEntryTime *time.Time `json:"schedule_entry_time" example:"2024-06-01T08:00:00Z"`
	ScheduleExitTime  *time.Time `json:"schedule_exit_time" example:"2024-06-01T17:00:00Z"`
}

func (this *Internship) ToDomain() (internship.Internship, errors.Error) {
	location, validationError := internshipLocation.NewBuilder().WithID(this.LocationID).Build()
	if validationError != nil {
		return nil, validationError
	}
	_student, validationError := simplifiedStudent.NewBuilder().WithID(this.StudentID).Build()
	if validationError != nil {
		return nil, validationError
	}
	_internship := internship.NewBuilder().
		WithStudent(_student).
		WithLocation(location).
		WithStartedIn(this.StartedIn).
		WithEndedIn(this.EndedIn).
		WithScheduleEntryTime(this.ScheduleEntryTime).
		WithScheduleExitTime(this.ScheduleExitTime)
	return _internship.Build()
}
