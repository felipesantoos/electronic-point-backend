package services

import (
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/timeRecord"
	"eletronic_point/src/core/interfaces/primary"
	"eletronic_point/src/core/interfaces/secondary"
	"eletronic_point/src/core/services/filters"

	"github.com/google/uuid"
)

type timeRecordServices struct {
	repository secondary.TimeRecordPort
}

func NewTimeRecordServices(repository secondary.TimeRecordPort) primary.TimeRecordPort {
	return &timeRecordServices{repository}
}

func (this *timeRecordServices) Create(_timeRecord timeRecord.TimeRecord) (*uuid.UUID, errors.Error) {
	return this.repository.Create(_timeRecord)
}

func (this *timeRecordServices) Update(_timeRecord timeRecord.TimeRecord) errors.Error {
	return this.repository.Update(_timeRecord)
}

func (this *timeRecordServices) Delete(id uuid.UUID) errors.Error {
	return this.repository.Delete(id)
}

func (this *timeRecordServices) List(_filters filters.TimeRecordFilters) ([]timeRecord.TimeRecord, errors.Error) {
	return this.repository.List(_filters)
}

func (this *timeRecordServices) Get(id uuid.UUID) (timeRecord.TimeRecord, errors.Error) {
	return this.repository.Get(id)
}
