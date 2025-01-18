package course

import (
	"eletronic_point/src/core/domain/errors"

	"github.com/google/uuid"
)

type Course interface {
	ID() uuid.UUID
	Name() string

	SetID(uuid.UUID) errors.Error
	SetName(string) errors.Error
}
