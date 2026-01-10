package primary

import (
	"eletronic_point/src/core/domain/campus"
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/services/filters"
	"github.com/google/uuid"
)

type CampusPort interface {
	List(_filters filters.CampusFilters) ([]campus.Campus, errors.Error)
	Get(id uuid.UUID) (campus.Campus, errors.Error)
	Create(data campus.Campus) (*uuid.UUID, errors.Error)
	Update(data campus.Campus) errors.Error
	Delete(id uuid.UUID) errors.Error
}
