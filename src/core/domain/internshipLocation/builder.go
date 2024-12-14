package internshiplocation

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

func (b *builder) WithAddress(address string) *builder {
	address = strings.TrimSpace(address)
	if len(address) == 0 {
		b.fields = append(b.fields, messages.InternshipLocationAddress)
		b.errorMessages = append(b.errorMessages, messages.InternshipLocationAddressErrorMessage)
		return b
	}
	b.location.address = address
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

func (b *builder) WithLat(lat *float64) *builder {
	b.location.lat = lat
	return b
}

func (b *builder) WithLong(long *float64) *builder {
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
