package filters

import (
	"github.com/google/uuid"
)

type InternshipFilters struct {
	StudentID *uuid.UUID
	TeacherID *uuid.UUID
}
