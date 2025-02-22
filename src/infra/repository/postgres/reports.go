package postgres

import (
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/student"
	"eletronic_point/src/core/interfaces/secondary"
	"eletronic_point/src/core/services/filters"
)

type reportsRepository struct{}

func NewReportsRepository() secondary.ReportsPort {
	return &reportsRepository{}
}

func (this reportsRepository) GetTimeRecordsByStudent(_filters filters.TimeRecordsByStudentFilters) ([]student.Student, errors.Error) {
	students, err := NewStudentRepository().List(_filters.StudentFilters)
	if err != nil {
		logger.Error().Msg(err.String())
		return nil, errors.NewUnexpected()
	}
	for i := range students {
		_filters.TimeRecordFilters.StudentID = students[i].ID()
		timeRecords, err := NewTimeRecordRepository().List(_filters.TimeRecordFilters)
		if err != nil {
			logger.Error().Msg(err.String())
			return nil, errors.NewUnexpected()
		}
		students[i].SetFrequencyHistory(timeRecords)
	}
	return students, nil
}
