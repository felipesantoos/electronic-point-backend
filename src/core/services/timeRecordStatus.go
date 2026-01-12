package services

import (
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/timeRecordStatus"
	"eletronic_point/src/core/interfaces/primary"
	"eletronic_point/src/core/interfaces/secondary"

	"github.com/google/uuid"
)

type timeRecordStatusServices struct {
	repository secondary.TimeRecordStatusPort
}

func NewTimeRecordStatusService(repository secondary.TimeRecordStatusPort) primary.TimeRecordStatusPort {
	return &timeRecordStatusServices{repository}
}

func (this *timeRecordStatusServices) List() ([]timeRecordStatus.TimeRecordStatus, errors.Error) {
	return this.repository.List()
}

func (this *timeRecordStatusServices) Get(id uuid.UUID) (timeRecordStatus.TimeRecordStatus, errors.Error) {
	return this.repository.Get(id)
}

func (this *timeRecordStatusServices) Create(data timeRecordStatus.TimeRecordStatus) (*uuid.UUID, errors.Error) {
	return this.repository.Create(data)
}

func (this *timeRecordStatusServices) Update(data timeRecordStatus.TimeRecordStatus) errors.Error {
	return this.repository.Update(data)
}

func (this *timeRecordStatusServices) Delete(id uuid.UUID) errors.Error {
	return this.repository.Delete(id)
}
