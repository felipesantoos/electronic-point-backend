package services

import (
	"dit_backend/src/core/domain/errors"
	"dit_backend/src/core/domain/role"
	"dit_backend/src/core/interfaces/adapters"
	"dit_backend/src/core/interfaces/usecases"
)

type resourcesService struct {
	adapter adapters.ResourcesAdapter
}

func NewResourcesService(adapter adapters.ResourcesAdapter) usecases.ResourcesUseCase {
	return &resourcesService{adapter}
}

func (instance *resourcesService) ListAccountRoles() ([]role.Role, errors.Error) {
	roles, err := instance.adapter.ListAccountRoles()
	if err != nil {
		return nil, errors.NewFromInfra(err)
	}
	return roles, nil
}
