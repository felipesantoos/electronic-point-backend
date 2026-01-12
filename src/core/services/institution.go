package services

import (
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/institution"
	"eletronic_point/src/core/interfaces/primary"
	"eletronic_point/src/core/interfaces/secondary"
	"eletronic_point/src/core/services/filters"
	"github.com/google/uuid"
)

type institutionServices struct {
	repository secondary.InstitutionPort
}

func NewInstitutionService(repository secondary.InstitutionPort) primary.InstitutionPort {
	return &institutionServices{repository}
}

func (this *institutionServices) List(_filters filters.InstitutionFilters) ([]institution.Institution, errors.Error) {
	return this.repository.List(_filters)
}

func (this *institutionServices) Get(id uuid.UUID) (institution.Institution, errors.Error) {
	return this.repository.Get(id)
}

func (this *institutionServices) Create(data institution.Institution) (*uuid.UUID, errors.Error) {
	return this.repository.Create(data)
}

func (this *institutionServices) Update(data institution.Institution) errors.Error {
	return this.repository.Update(data)
}

func (this *institutionServices) Delete(id uuid.UUID) errors.Error {
	return this.repository.Delete(id)
}
