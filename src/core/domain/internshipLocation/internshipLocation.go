package internshipLocation

import (
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/messages"

	"github.com/google/uuid"
)

type internshipLocation struct {
	id           uuid.UUID
	name         string
	number       string
	street       string
	neighborhood string
	city         string
	zipCode      string
	lat          float64
	long         float64
}

func (i *internshipLocation) ID() uuid.UUID {
	return i.id
}

func (i *internshipLocation) Name() string {
	return i.name
}

func (i *internshipLocation) Number() string {
	return i.number
}

func (i *internshipLocation) Street() string {
	return i.street
}

func (i *internshipLocation) Neighborhood() string {
	return i.neighborhood
}

func (i *internshipLocation) City() string {
	return i.city
}

func (i *internshipLocation) ZipCode() string {
	return i.zipCode
}

func (i *internshipLocation) Lat() float64 {
	return i.lat
}

func (i *internshipLocation) Long() float64 {
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

func (i *internshipLocation) SetNumber(number string) errors.Error {
	if number == "" {
		return errors.NewFromString(messages.InternshipLocationNumberErrorMessage)
	}
	i.number = number
	return nil
}

func (i *internshipLocation) SetStreet(street string) errors.Error {
	if street == "" {
		return errors.NewFromString(messages.InternshipLocationStreetErrorMessage)
	}
	i.street = street
	return nil
}

func (i *internshipLocation) SetNeighborhood(neighborhood string) errors.Error {
	if neighborhood == "" {
		return errors.NewFromString(messages.InternshipLocationNeighborhoodErrorMessage)
	}
	i.neighborhood = neighborhood
	return nil
}

func (i *internshipLocation) SetCity(city string) errors.Error {
	if city == "" {
		return errors.NewFromString(messages.InternshipLocationCityErrorMessage)
	}
	i.city = city
	return nil
}

func (i *internshipLocation) SetZipCode(zipCode string) errors.Error {
	if zipCode == "" {
		return errors.NewFromString(messages.InternshipLocationZipCodeErrorMessage)
	}
	i.zipCode = zipCode
	return nil
}

func (i *internshipLocation) SetLat(lat float64) errors.Error {
	i.lat = lat
	return nil
}

func (i *internshipLocation) SetLong(long float64) errors.Error {
	return nil
}
