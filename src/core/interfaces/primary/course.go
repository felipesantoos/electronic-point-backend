package primary

import (
	"eletronic_point/src/core/domain/course"
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/services/filters"
)

type CoursePort interface {
	List(_filters filters.CourseFilters) ([]course.Course, errors.Error)
}
