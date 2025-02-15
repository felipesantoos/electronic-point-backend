package services

import (
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/institution"
	"eletronic_point/src/core/interfaces/primary"
	"eletronic_point/src/core/interfaces/secondary"
	"eletronic_point/src/core/services/filters"
)

type institutionServices struct {
	repository secondary.InstitutionPort
}

func NewInstitutionServices(repository secondary.InstitutionPort) primary.InstitutionPort {
	return &institutionServices{repository}
}

func (this *institutionServices) List(_filters filters.InstitutionFilters) ([]institution.Institution, errors.Error) {
	return this.repository.List(_filters)
}
