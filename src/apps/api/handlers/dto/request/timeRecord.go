package request

import (
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/timeRecord"
	"time"

	"github.com/google/uuid"
)

type TimeRecord struct {
	Date          time.Time  `json:"date" example:"2024-12-01T00:00:00Z"`
	EntryTime     time.Time  `json:"entry_time" example:"2024-12-01T08:00:00Z"`
	ExitTime      *time.Time `json:"exit_time" example:"2024-12-01T16:00:00Z"`
	Location      string     `json:"location" example:"Localização 1"`
	IsOffSite     bool       `json:"is_off_site" example:"false"`
	Justification *string    `json:"justification" example:""`
	StudentID     uuid.UUID  `json:"student_id" example:"5fa6d07d-4e5a-4d27-8f8b-3de0dbec5c65"`
}

func (this *TimeRecord) ToDomain() (timeRecord.TimeRecord, errors.Error) {
	builder := timeRecord.NewBuilder().
		WithDate(this.Date).
		WithEntryTime(this.EntryTime).
		WithExitTime(this.ExitTime).
		WithLocation(this.Location).
		WithIsOffSite(this.IsOffSite).
		WithJustification(this.Justification).
		WithStudentID(this.StudentID)
	return builder.Build()
}
