package request

import (
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/timeRecord"
	"time"

	"github.com/google/uuid"
)

type TimeRecord struct {
	InternshipID  uuid.UUID  `json:"internship_id" form:"internship_id" example:"9ec95529-06d5-47e2-b617-1606088ac9e6"`
	StudentID     *uuid.UUID `json:"student_id,omitempty" form:"student_id" example:"02e62826-bf41-4944-adb2-051b6a30a131"`
	TeacherID     *uuid.UUID `json:"teacher_id,omitempty" form:"teacher_id" example:"ea11bb4b-9aed-4444-9c00-f80bde564063"`
	Date          time.Time  `json:"date" form:"date" example:"2024-12-01T00:00:00Z"`
	EntryTime     time.Time  `json:"entry_time" form:"entry_time" example:"2024-12-01T08:00:00Z"`
	ExitTime      *time.Time `json:"exit_time" form:"exit_time" example:"2024-12-01T16:00:00Z"`
	Location      string     `json:"location" form:"location" example:"Localização 1"`
	IsOffSite     bool       `json:"is_off_site" form:"is_off_site" example:"false"`
	Justification *string    `json:"justification" form:"justification" example:""`
}

func (this *TimeRecord) ToDomain() (timeRecord.TimeRecord, errors.Error) {
	builder := timeRecord.NewBuilder().
		WithInternshipID(this.InternshipID).
		WithDate(this.Date).
		WithEntryTime(this.EntryTime).
		WithExitTime(this.ExitTime).
		WithLocation(this.Location).
		WithIsOffSite(this.IsOffSite).
		WithJustification(this.Justification)
	return builder.Build()
}
