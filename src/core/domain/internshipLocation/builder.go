package internshipLocation

import (
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/messages"
	"strings"

	"github.com/google/uuid"
)

type builder struct {
	fields        []string
	errorMessages []string
	location      *internshipLocation
}

func NewBuilder() *builder {
	return &builder{
		fields:        []string{},
		errorMessages: []string{},
		location:      &internshipLocation{},
	}
}

func (b *builder) WithID(id uuid.UUID) *builder {
	if id == uuid.Nil {
		b.fields = append(b.fields, messages.InternshipLocationID)
		b.errorMessages = append(b.errorMessages, messages.InternshipLocationIDErrorMessage)
		return b
	}
	b.location.id = id
	return b
}

func (b *builder) WithName(name string) *builder {
	name = strings.TrimSpace(name)
	if len(name) == 0 {
		b.fields = append(b.fields, messages.InternshipLocationName)
		b.errorMessages = append(b.errorMessages, messages.InternshipLocationNameErrorMessage)
		return b
	}
	b.location.name = name
	return b
}

func (b *builder) WithNumber(number string) *builder {
	number = strings.TrimSpace(number)
	if len(number) == 0 {
		b.fields = append(b.fields, messages.InternshipLocationNumber)
		b.errorMessages = append(b.errorMessages, messages.InternshipLocationNumberErrorMessage)
		return b
	}
	b.location.number = number
	return b
}

func (b *builder) WithStreet(street string) *builder {
	street = strings.TrimSpace(street)
	if len(street) == 0 {
		b.fields = append(b.fields, messages.InternshipLocationStreet)
		b.errorMessages = append(b.errorMessages, messages.InternshipLocationStreetErrorMessage)
		return b
	}
	b.location.street = street
	return b
}

func (b *builder) WithNeighborhood(neighborhood string) *builder {
	neighborhood = strings.TrimSpace(neighborhood)
	if len(neighborhood) == 0 {
		b.fields = append(b.fields, messages.InternshipLocationNeighborhood)
		b.errorMessages = append(b.errorMessages, messages.InternshipLocationNeighborhoodErrorMessage)
		return b
	}
	b.location.neighborhood = neighborhood
	return b
}

func (b *builder) WithCity(city string) *builder {
	city = strings.TrimSpace(city)
	if len(city) == 0 {
		b.fields = append(b.fields, messages.InternshipLocationCity)
		b.errorMessages = append(b.errorMessages, messages.InternshipLocationCityErrorMessage)
		return b
	}
	b.location.city = city
	return b
}

func (b *builder) WithZipCode(zipCode string) *builder {
	zipCode = strings.TrimSpace(zipCode)
	if len(zipCode) == 0 {
		b.fields = append(b.fields, messages.InternshipLocationZipCode)
		b.errorMessages = append(b.errorMessages, messages.InternshipLocationZipCodeErrorMessage)
		return b
	}
	b.location.zipCode = zipCode
	return b
}

func (b *builder) WithLat(lat float64) *builder {
	b.location.lat = lat
	return b
}

func (b *builder) WithLong(long float64) *builder {
	b.location.long = long
	return b
}

func (b *builder) Build() (*internshipLocation, errors.Error) {
	if len(b.errorMessages) > 0 {
		return nil, errors.NewValidationWithMetadata(b.errorMessages, map[string]interface{}{
			messages.Fields: b.fields,
		})
	}
	return b.location, nil
}
