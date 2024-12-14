package internshiplocation

import (
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/messages"

	"github.com/google/uuid"
)

type internshipLocation struct {
	id      uuid.UUID
	name    string
	address string
	city    string
	lat     *float64
	long    *float64
}

func (i *internshipLocation) ID() uuid.UUID {
	return i.id
}

func (i *internshipLocation) Name() string {
	return i.name
}

func (i *internshipLocation) Address() string {
	return i.address
}

func (i *internshipLocation) City() string {
	return i.city
}

func (i *internshipLocation) Lat() *float64 {
	return i.lat
}

func (i *internshipLocation) Long() *float64 {
	return i.long
}

func (i *internshipLocation) SetID(id uuid.UUID) errors.Error {
	if id == uuid.Nil {
		return errors.NewFromString(messages.InternshipLocationIDErrorMessage)
	}
	i.id = id
	return nil
}

func (i *internshipLocation) SetName(name string) errors.Error {
	if name == "" {
		return errors.NewFromString(messages.InternshipLocationNameErrorMessage)
	}
	i.name = name
	return nil
}

func (i *internshipLocation) SetAddress(address string) errors.Error {
	if address == "" {
		return errors.NewFromString(messages.InternshipLocationAddressErrorMessage)
	}
	i.address = address
	return nil
}

func (i *internshipLocation) SetCity(city string) errors.Error {
	if city == "" {
		return errors.NewFromString(messages.InternshipLocationCityErrorMessage)
	}
	i.city = city
	return nil
}

func (i *internshipLocation) SetLat(lat *float64) errors.Error {
	i.lat = lat
	return nil
}

func (i *internshipLocation) SetLong(long *float64) errors.Error {
	return nil
}
