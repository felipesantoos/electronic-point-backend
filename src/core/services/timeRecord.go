package services

import (
	"backend_template/src/core/domain/errors"
	"backend_template/src/core/domain/timeRecord"
	"backend_template/src/core/interfaces/adapters"
	"backend_template/src/core/interfaces/usecases"

	"github.com/google/uuid"
)

type timeRecordService struct {
	adapter adapters.TimeRecordAdapter
}

func NewTimeRecordService(adapter adapters.TimeRecordAdapter) usecases.TimeRecordUseCase {
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
