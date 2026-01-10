package primary

import (
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/timeRecordStatus"

	"github.com/google/uuid"
)

type TimeRecordStatusPort interface {
	List() ([]timeRecordStatus.TimeRecordStatus, errors.Error)
	Get(id uuid.UUID) (timeRecordStatus.TimeRecordStatus, errors.Error)
	Create(data timeRecordStatus.TimeRecordStatus) (*uuid.UUID, errors.Error)
	Update(data timeRecordStatus.TimeRecordStatus) errors.Error
	Delete(id uuid.UUID) errors.Error
}
