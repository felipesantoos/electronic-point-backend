package primary

import (
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/timeRecordStatus"

	"github.com/google/uuid"
)

type TimeRecordStatusPort interface {
	List() ([]timeRecordStatus.TimeRecordStatus, errors.Error)
	Get(id uuid.UUID) (timeRecordStatus.TimeRecordStatus, errors.Error)
}
