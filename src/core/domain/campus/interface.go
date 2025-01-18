package campus

import (
	"eletronic_point/src/core/domain/errors"

	"github.com/google/uuid"
)

type Campus interface {
	ID() uuid.UUID
	Name() string
	InstitutionID() uuid.UUID

	SetID(uuid.UUID) errors.Error
	SetName(string) errors.Error
	SetInstitutionID(uuid.UUID) errors.Error
}
