package primary

import (
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/internshipLocation"
	"eletronic_point/src/core/services/filters"

	"github.com/google/uuid"
)

type InternshipLocationPort interface {
	Create(internshipLocation internshipLocation.InternshipLocation) (*uuid.UUID, errors.Error)
	Update(internshipLocation internshipLocation.InternshipLocation) errors.Error
	Delete(id uuid.UUID) errors.Error
	List(_filters filters.InternshipLocationFilters) ([]internshipLocation.InternshipLocation, errors.Error)
	Get(id uuid.UUID) (internshipLocation.InternshipLocation, errors.Error)
}
