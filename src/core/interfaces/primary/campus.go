package primary

import (
	"eletronic_point/src/core/domain/campus"
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/services/filters"
)

type CampusPort interface {
	List(_filters filters.CampusFilters) ([]campus.Campus, errors.Error)
}
