package usecases

import (
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/role"
)

type ResourcesUseCase interface {
	ListAccountRoles() ([]role.Role, errors.Error)
}
