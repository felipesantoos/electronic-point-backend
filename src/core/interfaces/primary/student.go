package primary

import (
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/student"
	"eletronic_point/src/core/services/filters"

	"github.com/google/uuid"
)

type StudentPort interface {
	Create(student student.Student) (*uuid.UUID, errors.Error)
	Update(student student.Student) errors.Error
	Delete(id uuid.UUID) errors.Error
	List(_filters filters.StudentFilters) ([]student.Student, errors.Error)
	Get(id uuid.UUID, _filters filters.StudentFilters) (student.Student, errors.Error)
}
