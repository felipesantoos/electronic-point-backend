package filters

import "github.com/google/uuid"

type CampusFilters struct {
	Name          *string
	InstitutionID *uuid.UUID
}
