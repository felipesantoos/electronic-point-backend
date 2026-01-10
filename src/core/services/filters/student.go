package filters

import (
	"github.com/google/uuid"
)

type StudentFilters struct {
	TeacherID     *uuid.UUID
	InstitutionID *uuid.UUID
	CampusID      *uuid.UUID
	Search        *string
}
