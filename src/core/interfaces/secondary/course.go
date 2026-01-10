package secondary

import (
	"eletronic_point/src/core/domain/course"
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/services/filters"
	"github.com/google/uuid"
)

type CoursePort interface {
	List(_filters filters.CourseFilters) ([]course.Course, errors.Error)
	Get(id uuid.UUID) (course.Course, errors.Error)
	Create(data course.Course) (*uuid.UUID, errors.Error)
	Update(data course.Course) errors.Error
	Delete(id uuid.UUID) errors.Error
}
