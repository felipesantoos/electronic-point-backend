package services

import (
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/role"
	"eletronic_point/src/core/interfaces/adapters"
	"eletronic_point/src/core/interfaces/usecases"
)

type resourcesService struct {
	adapter adapters.ResourcesAdapter
}

func NewResourcesService(adapter adapters.ResourcesAdapter) usecases.ResourcesUseCase {
	return &resourcesService{adapter}
}

func (s *resourcesService) ListAccountRoles() ([]role.Role, errors.Error) {
	return s.adapter.ListAccountRoles()
}
