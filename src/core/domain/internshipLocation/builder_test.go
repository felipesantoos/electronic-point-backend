package internshipLocation

import (
	"eletronic_point/src/core/messages"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestBuilder_WithID(t *testing.T) {
	tests := []struct {
		name          string
		id            uuid.UUID
		expectedError bool
	}{
		{"Valid ID", uuid.New(), false},
		{"Nil ID", uuid.Nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBuilder().WithID(tt.id)
			_, err := b.Build()

			if tt.expectedError {
				assert.NotNil(t, err)
				assert.Contains(t, err.Messages(), messages.InternshipLocationIDErrorMessage)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestBuilder_WithName(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		expectedError bool
	}{
		{"Valid Name", "Location Name", false},
		{"Empty Name", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBuilder().WithName(tt.input)
			_, err := b.Build()

			if tt.expectedError {
				assert.NotNil(t, err)
				assert.Contains(t, err.Messages(), messages.InternshipLocationNameErrorMessage)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestBuilder_WithAddressFields(t *testing.T) {
	t.Run("Valid Fields", func(t *testing.T) {
		b := NewBuilder().
			WithNumber("123").
			WithStreet("Street").
			WithNeighborhood("Neighborhood").
			WithCity("City").
			WithZipCode("12345-678")
		_, err := b.Build()
		assert.Nil(t, err)
	})

	t.Run("Empty Number", func(t *testing.T) {
		b := NewBuilder().WithNumber("")
		_, err := b.Build()
		assert.NotNil(t, err)
		assert.Contains(t, err.Messages(), messages.InternshipLocationNumberErrorMessage)
	})

	t.Run("Empty Street", func(t *testing.T) {
		b := NewBuilder().WithStreet("")
		_, err := b.Build()
		assert.NotNil(t, err)
		assert.Contains(t, err.Messages(), messages.InternshipLocationStreetErrorMessage)
	})

	t.Run("Empty Neighborhood", func(t *testing.T) {
		b := NewBuilder().WithNeighborhood("")
		_, err := b.Build()
		assert.NotNil(t, err)
		assert.Contains(t, err.Messages(), messages.InternshipLocationNeighborhoodErrorMessage)
	})

	t.Run("Empty City", func(t *testing.T) {
		b := NewBuilder().WithCity("")
		_, err := b.Build()
		assert.NotNil(t, err)
		assert.Contains(t, err.Messages(), messages.InternshipLocationCityErrorMessage)
	})

	t.Run("Empty ZipCode", func(t *testing.T) {
		b := NewBuilder().WithZipCode("")
		_, err := b.Build()
		assert.NotNil(t, err)
		assert.Contains(t, err.Messages(), messages.InternshipLocationZipCodeErrorMessage)
	})
}

func TestBuilder_Build(t *testing.T) {
	t.Run("Successful Build", func(t *testing.T) {
		id := uuid.New()
		name := "Location Name"
		number := "123"
		street := "Street"
		neighborhood := "Neighborhood"
		city := "City"
		zipCode := "12345-678"
		lat := -10.0
		long := -20.0

		loc, err := NewBuilder().
			WithID(id).
			WithName(name).
			WithNumber(number).
			WithStreet(street).
			WithNeighborhood(neighborhood).
			WithCity(city).
			WithZipCode(zipCode).
			WithLat(lat).
			WithLong(long).
			Build()

		assert.Nil(t, err)
		assert.NotNil(t, loc)
		assert.Equal(t, id, loc.ID())
		assert.Equal(t, name, loc.Name())
		assert.Equal(t, number, loc.Number())
		assert.Equal(t, street, loc.Street())
		assert.Equal(t, neighborhood, loc.Neighborhood())
		assert.Equal(t, city, loc.City())
		assert.Equal(t, zipCode, loc.ZipCode())
		assert.Equal(t, lat, loc.Lat())
		assert.Equal(t, long, loc.Long())
	})
}
