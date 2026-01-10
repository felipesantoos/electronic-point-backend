package primary

import (
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/institution"
	"eletronic_point/src/core/services/filters"
	"github.com/google/uuid"
)

type InstitutionPort interface {
	List(_filters filters.InstitutionFilters) ([]institution.Institution, errors.Error)
	Get(id uuid.UUID) (institution.Institution, errors.Error)
	Create(data institution.Institution) (*uuid.UUID, errors.Error)
	Update(data institution.Institution) errors.Error
	Delete(id uuid.UUID) errors.Error
}
