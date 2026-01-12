package services

import (
	"eletronic_point/src/core/domain/campus"
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/interfaces/primary"
	"eletronic_point/src/core/interfaces/secondary"
	"eletronic_point/src/core/services/filters"
	"github.com/google/uuid"
)

type campusServices struct {
	repository secondary.CampusPort
}

func NewCampusService(repository secondary.CampusPort) primary.CampusPort {
	return &campusServices{repository}
}

func (this *campusServices) List(_filters filters.CampusFilters) ([]campus.Campus, errors.Error) {
	return this.repository.List(_filters)
}

func (this *campusServices) Get(id uuid.UUID) (campus.Campus, errors.Error) {
	return this.repository.Get(id)
}

func (this *campusServices) Create(data campus.Campus) (*uuid.UUID, errors.Error) {
	return this.repository.Create(data)
}

func (this *campusServices) Update(data campus.Campus) errors.Error {
	return this.repository.Update(data)
}

func (this *campusServices) Delete(id uuid.UUID) errors.Error {
	return this.repository.Delete(id)
}
