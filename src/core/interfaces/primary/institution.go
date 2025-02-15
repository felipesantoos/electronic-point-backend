package primary

import (
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/institution"
	"eletronic_point/src/core/services/filters"
)

type InstitutionPort interface {
	List(_filters filters.InstitutionFilters) ([]institution.Institution, errors.Error)
}
