package services

import (
	"eletronic_point/src/core/domain/campus"
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/interfaces/primary"
	"eletronic_point/src/core/interfaces/secondary"
	"eletronic_point/src/core/services/filters"
)

type campusServices struct {
	repository secondary.CampusPort
}

func NewCampusServices(repository secondary.CampusPort) primary.CampusPort {
	return &campusServices{repository}
}

func (this *campusServices) List(_filters filters.CampusFilters) ([]campus.Campus, errors.Error) {
	return this.repository.List(_filters)
}
