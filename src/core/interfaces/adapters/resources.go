package adapters

import (
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/role"
)

type ResourcesAdapter interface {
	ListAccountRoles() ([]role.Role, errors.Error)
}
