package services

import (
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/student"
	"eletronic_point/src/core/interfaces/primary"
	"eletronic_point/src/core/interfaces/secondary"
	"eletronic_point/src/core/services/filters"
)

type reportsServices struct {
	repository secondary.ReportsPort
}

func NewReportsServices(repository secondary.ReportsPort) primary.ReportsPort {
	return &reportsServices{repository}
}

func (this *reportsServices) GetTimeRecordsByStudent(_filters filters.TimeRecordsByStudentFilters) ([]student.Student, errors.Error) {
	return this.repository.GetTimeRecordsByStudent(_filters)
}
