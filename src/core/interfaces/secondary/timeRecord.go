package secondary

import (
	"eletronic_point/src/core/domain/errors"
	timeRecord "eletronic_point/src/core/domain/timeRecord"
	"eletronic_point/src/core/services/filters"

	"github.com/google/uuid"
)

type TimeRecordPort interface {
	Create(timeRecord timeRecord.TimeRecord) (*uuid.UUID, errors.Error)
	Update(timeRecord timeRecord.TimeRecord) errors.Error
	Delete(id uuid.UUID) errors.Error
	List(_filters filters.TimeRecordFilters) ([]timeRecord.TimeRecord, errors.Error)
	Get(id uuid.UUID, _filters filters.TimeRecordFilters) (timeRecord.TimeRecord, errors.Error)
}
