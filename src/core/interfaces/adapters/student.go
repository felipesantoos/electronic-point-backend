package adapters

import (
	"backend_template/src/core/domain/errors"
	"backend_template/src/core/domain/student"

	"github.com/google/uuid"
)

type StudentAdapter interface {
	Create(student student.Student) (*uuid.UUID, errors.Error)
	Update(student student.Student) errors.Error
	Delete(id uuid.UUID) errors.Error
	List() ([]student.Student, errors.Error)
	Get(id uuid.UUID) (student.Student, errors.Error)
}
