package usecases

import (
	"backend_template/src/core/domain/errors"
	"backend_template/src/core/domain/role"
)

type ResourcesUseCase interface {
	ListAccountRoles() ([]role.Role, errors.Error)
}
