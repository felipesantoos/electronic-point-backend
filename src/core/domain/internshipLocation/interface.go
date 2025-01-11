package internshipLocation

import (
	"eletronic_point/src/core/domain/errors"

	"github.com/google/uuid"
)

type InternshipLocation interface {
	ID() uuid.UUID
	Name() string
	Number() string
	Street() string
	Neighborhood() string
	City() string
	ZipCode() string
	Lat() float64
	Long() float64

	SetID(uuid.UUID) errors.Error
	SetName(string) errors.Error
	SetNumber(string) errors.Error
	SetStreet(string) errors.Error
	SetNeighborhood(string) errors.Error
	SetCity(string) errors.Error
	SetZipCode(string) errors.Error
	SetLat(float64) errors.Error
	SetLong(float64) errors.Error
}
