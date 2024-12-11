package services

import (
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/timeRecord"
	"eletronic_point/src/core/interfaces/primary"
	"eletronic_point/src/core/interfaces/secondary"

	"github.com/google/uuid"
)

type timeRecordService struct {
	adapter secondary.TimeRecordPort
}

func NewTimeRecordService(adapter secondary.TimeRecordPort) primary.TimeRecordPort {
	return &timeRecordService{adapter}
}

func (s *timeRecordService) Create(timeRecord timeRecord.TimeRecord) (*uuid.UUID, errors.Error) {
	return s.adapter.Create(timeRecord)
}

func (s *timeRecordService) Update(timeRecord timeRecord.TimeRecord) errors.Error {
	return s.adapter.Update(timeRecord)
}

func (s *timeRecordService) Delete(id uuid.UUID) errors.Error {
	return s.adapter.Delete(id)
}

func (s *timeRecordService) List() ([]timeRecord.TimeRecord, errors.Error) {
	return s.adapter.List()
}

func (s *timeRecordService) Get(id uuid.UUID) (timeRecord.TimeRecord, errors.Error) {
	return s.adapter.Get(id)
}
