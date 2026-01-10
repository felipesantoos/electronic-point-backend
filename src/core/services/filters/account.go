package filters

import (
	"github.com/google/uuid"
)

type AccountFilters struct {
	RoleID *uuid.UUID
	Search *string
}
