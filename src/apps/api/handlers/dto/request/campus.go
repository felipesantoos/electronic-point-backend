package request

import "github.com/google/uuid"

type Campus struct {
	Name          string    `json:"name" form:"name"`
	InstitutionID uuid.UUID `json:"institution_id" form:"institution_id"`
}
