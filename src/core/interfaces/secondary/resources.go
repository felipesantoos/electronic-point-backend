package secondary

import (
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/role"
)

type ResourcesPort interface {
	ListAccountRoles() ([]role.Role, errors.Error)
}
