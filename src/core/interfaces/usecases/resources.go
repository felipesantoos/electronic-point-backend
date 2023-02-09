package usecases

import (
	"dit_backend/src/core/domain/errors"
	"dit_backend/src/core/domain/role"
)

type ResourcesUseCase interface {
	ListAccountRoles() ([]role.Role, errors.Error)
}
