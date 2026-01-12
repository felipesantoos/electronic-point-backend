package services

import (
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/internshipLocation"
	"eletronic_point/src/core/interfaces/primary"
	"eletronic_point/src/core/interfaces/secondary"
	"eletronic_point/src/core/services/filters"

	"github.com/google/uuid"
)

type internshipLocationServices struct {
	repository secondary.InternshipLocationPort
}

func NewInternshipLocationService(repository secondary.InternshipLocationPort) primary.InternshipLocationPort {
	return &internshipLocationServices{repository}
}

func (this *internshipLocationServices) Create(_internshipLocation internshipLocation.InternshipLocation) (*uuid.UUID, errors.Error) {
	return this.repository.Create(_internshipLocation)
}

func (this *internshipLocationServices) Update(_internshipLocation internshipLocation.InternshipLocation) errors.Error {
	return this.repository.Update(_internshipLocation)
}

func (this *internshipLocationServices) Delete(id uuid.UUID) errors.Error {
	return this.repository.Delete(id)
}

func (this *internshipLocationServices) List(_filters filters.InternshipLocationFilters) ([]internshipLocation.InternshipLocation, errors.Error) {
	return this.repository.List(_filters)
}

func (this *internshipLocationServices) Get(id uuid.UUID) (internshipLocation.InternshipLocation, errors.Error) {
	return this.repository.Get(id)
}
