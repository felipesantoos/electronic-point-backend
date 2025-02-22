package primary

import (
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/student"
	"eletronic_point/src/core/services/filters"
)

type ReportsPort interface {
	GetTimeRecordsByStudent(_filters filters.TimeRecordsByStudentFilters) ([]student.Student, errors.Error)
}
