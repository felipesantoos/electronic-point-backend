package institution

import (
	"eletronic_point/src/core/domain/errors"

	"github.com/google/uuid"
)

type Institution interface {
	ID() uuid.UUID
	Name() string

	SetID(uuid.UUID) errors.Error
	SetName(string) errors.Error
}
