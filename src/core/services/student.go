package services

import (
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/student"
	"eletronic_point/src/core/interfaces/primary"
	"eletronic_point/src/core/interfaces/secondary"

	"github.com/google/uuid"
)

type studentServices struct {
	repository secondary.StudentPort
}

func NewStudentServices(repository secondary.StudentPort) primary.StudentPort {
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

func (this *studentServices) List() ([]student.Student, errors.Error) {
	return this.repository.List()
}

func (this *studentServices) Get(id uuid.UUID) (student.Student, errors.Error) {
	return this.repository.Get(id)
}
