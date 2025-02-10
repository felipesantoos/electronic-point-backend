package filters

import (
	"time"

	"github.com/google/uuid"
)

type TimeRecordFilters struct {
	StudentID *uuid.UUID
	StartDate *time.Time
	EndDate   *time.Time
	TeacherID *uuid.UUID
	StatusID  *uuid.UUID
}
