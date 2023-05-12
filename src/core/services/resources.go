package services

import (
	"backend_template/src/core/domain/errors"
	"backend_template/src/core/domain/role"
	"backend_template/src/core/interfaces/adapters"
	"backend_template/src/core/interfaces/usecases"
)

type resourcesService struct {
	adapter adapters.ResourcesAdapter
}

func NewResourcesService(adapter adapters.ResourcesAdapter) usecases.ResourcesUseCase {
	return &resourcesService{adapter}
}

func (instance *resourcesService) ListAccountRoles() ([]role.Role, errors.Error) {
	return instance.adapter.ListAccountRoles()
}
