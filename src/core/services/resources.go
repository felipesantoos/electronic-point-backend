package services

import (
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/role"
	secondary "eletronic_point/src/core/interfaces/adapters"
	"eletronic_point/src/core/interfaces/usecases"
)

type resourcesService struct {
	adapter secondary.ResourcesPort
}

func NewResourcesService(adapter secondary.ResourcesPort) usecases.ResourcesUseCase {
	return &resourcesService{adapter}
}

func (s *resourcesService) ListAccountRoles() ([]role.Role, errors.Error) {
	return s.adapter.ListAccountRoles()
}
