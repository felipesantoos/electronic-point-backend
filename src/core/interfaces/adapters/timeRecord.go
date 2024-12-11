package adapters

import (
	"backend_template/src/core/domain/errors"
	timeRecord "backend_template/src/core/domain/timeRecord"

	"github.com/google/uuid"
)

type TimeRecordAdapter interface {
	Create(timeRecord timeRecord.TimeRecord) (*uuid.UUID, errors.Error)
	Update(timeRecord timeRecord.TimeRecord) errors.Error
	Delete(id uuid.UUID) errors.Error
	List() ([]timeRecord.TimeRecord, errors.Error)
	Get(id uuid.UUID) (timeRecord.TimeRecord, errors.Error)
}
