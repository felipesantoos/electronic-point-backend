package services

import (
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/role"
	"eletronic_point/src/core/interfaces/primary"
	"eletronic_point/src/core/interfaces/secondary"
)

type resourcesService struct {
	adapter secondary.ResourcesPort
}

func NewResourcesService(adapter secondary.ResourcesPort) primary.ResourcesPort {
	return &resourcesService{adapter}
}

func (s *resourcesService) ListAccountRoles() ([]role.Role, errors.Error) {
	return s.adapter.ListAccountRoles()
}
