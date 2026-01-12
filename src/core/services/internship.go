package services

import (
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/internship"
	"eletronic_point/src/core/interfaces/primary"
	"eletronic_point/src/core/interfaces/secondary"
	"eletronic_point/src/core/services/filters"

	"github.com/google/uuid"
)

type internshipServices struct {
	repository secondary.InternshipPort
}

func NewInternshipService(repository secondary.InternshipPort) primary.InternshipPort {
	return &internshipServices{repository}
}

func (this *internshipServices) Create(_internship internship.Internship) (*uuid.UUID, errors.Error) {
	return this.repository.Create(_internship)
}

func (this *internshipServices) Update(_internship internship.Internship) errors.Error {
	return this.repository.Update(_internship)
}

func (this *internshipServices) Delete(id uuid.UUID) errors.Error {
	return this.repository.Delete(id)
}

func (this *internshipServices) List(_filters filters.InternshipFilters) ([]internship.Internship, errors.Error) {
	return this.repository.List(_filters)
}

func (this *internshipServices) Get(id uuid.UUID) (internship.Internship, errors.Error) {
	return this.repository.Get(id)
}
