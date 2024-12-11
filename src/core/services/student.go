package services

import (
	"backend_template/src/core/domain/errors"
	"backend_template/src/core/domain/student"
	"backend_template/src/core/interfaces/adapters"
	"backend_template/src/core/interfaces/usecases"

	"github.com/google/uuid"
)

type studentService struct {
	adapter adapters.StudentAdapter
}

func NewStudentService(adapter adapters.StudentAdapter) usecases.StudentUseCase {
	return &studentService{adapter}
}

func (s *studentService) Create(student student.Student) (*uuid.UUID, errors.Error) {
	return s.adapter.Create(student)
}

func (s *studentService) Update(student student.Student) errors.Error {
	return s.adapter.Update(student)
}

func (s *studentService) Delete(id uuid.UUID) errors.Error {
	return s.adapter.Delete(id)
}

func (s *studentService) List() ([]student.Student, errors.Error) {
	return s.adapter.List()
}

func (s *studentService) Get(id uuid.UUID) (student.Student, errors.Error) {
	return s.adapter.Get(id)
}
