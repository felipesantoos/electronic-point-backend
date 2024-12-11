package primary

import (
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/timeRecord"

	"github.com/google/uuid"
)

type TimeRecordPort interface {
	Create(timeRecord timeRecord.TimeRecord) (*uuid.UUID, errors.Error)
	Update(timeRecord timeRecord.TimeRecord) errors.Error
	Delete(id uuid.UUID) errors.Error
	List() ([]timeRecord.TimeRecord, errors.Error)
	Get(id uuid.UUID) (timeRecord.TimeRecord, errors.Error)
}
