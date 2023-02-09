package adapters

import (
	"dit_backend/src/core/domain/role"
	"dit_backend/src/infra"
)

type ResourcesAdapter interface {
	ListAccountRoles() ([]role.Role, infra.Error)
}
