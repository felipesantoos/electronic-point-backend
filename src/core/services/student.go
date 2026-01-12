package services

import (
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/student"
	"eletronic_point/src/core/interfaces/primary"
	"eletronic_point/src/core/interfaces/secondary"
	"eletronic_point/src/core/services/filters"

	"github.com/google/uuid"
)

type studentServices struct {
	repository secondary.StudentPort
}

func NewStudentService(repository secondary.StudentPort) primary.StudentPort {
	return &studentServices{repository}
}

func (this *studentServices) Create(_student student.Student) (*uuid.UUID, errors.Error) {
	return this.repository.Create(_student)
}

func (this *studentServices) Update(student student.Student) errors.Error {
	return this.repository.Update(student)
}

func (this *studentServices) Delete(id uuid.UUID) errors.Error {
	return this.repository.Delete(id)
}

func (this *studentServices) List(_filters filters.StudentFilters) ([]student.Student, errors.Error) {
	return this.repository.List(_filters)
}

func (this *studentServices) Get(id uuid.UUID, _filters filters.StudentFilters) (student.Student, errors.Error) {
	return this.repository.Get(id, _filters)
}
