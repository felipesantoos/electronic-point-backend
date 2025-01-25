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

func NewTimeRecordStatusServices(repository secondary.TimeRecordStatusPort) primary.TimeRecordStatusPort {
	return &timeRecordStatusServices{repository}
}

func (this *timeRecordStatusServices) List() ([]timeRecordStatus.TimeRecordStatus, errors.Error) {
	return this.repository.List()
}

func (this *timeRecordStatusServices) Get(id uuid.UUID) (timeRecordStatus.TimeRecordStatus, errors.Error) {
	return this.repository.Get(id)
}
