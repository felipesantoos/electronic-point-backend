package primary

import (
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/internship"
	"eletronic_point/src/core/services/filters"

	"github.com/google/uuid"
)

type InternshipPort interface {
	Create(internship internship.Internship) (*uuid.UUID, errors.Error)
	Update(internship internship.Internship) errors.Error
	Delete(id uuid.UUID) errors.Error
	List(_filters filters.InternshipFilters) ([]internship.Internship, errors.Error)
	Get(id uuid.UUID) (internship.Internship, errors.Error)
}
