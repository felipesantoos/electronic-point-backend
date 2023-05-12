package adapters

import (
	"backend_template/src/core/domain/errors"
	"backend_template/src/core/domain/role"
)

type ResourcesAdapter interface {
	ListAccountRoles() ([]role.Role, errors.Error)
}
