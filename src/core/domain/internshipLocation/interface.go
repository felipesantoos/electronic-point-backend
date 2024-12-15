package internshipLocation

import (
	"eletronic_point/src/core/domain/errors"

	"github.com/google/uuid"
)

type InternshipLocation interface {
	ID() uuid.UUID
	Name() string
	Address() string
	City() string
	Lat() *float64
	Long() *float64

	SetID(uuid.UUID) errors.Error
	SetName(string) errors.Error
	SetAddress(string) errors.Error
	SetCity(string) errors.Error
	SetLat(*float64) errors.Error
	SetLong(*float64) errors.Error
}
