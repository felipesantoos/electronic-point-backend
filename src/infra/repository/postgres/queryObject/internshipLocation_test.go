package queryObject

import (
	"eletronic_point/src/infra/repository/postgres/query"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestInternshipLocationQueryObject_FromMap(t *testing.T) {
	id := uuid.New()
	data := map[string]interface{}{
		query.InternshipLocationID:           []uint8(id.String()),
		query.InternshipLocationName:         "Location A",
		query.InternshipLocationNumber:       "123",
		query.InternshipLocationStreet:       "Street A",
		query.InternshipLocationNeighborhood: "Neighborhood A",
		query.InternshipLocationCity:         "City A",
		query.InternshipLocationZipCode:      "12345",
	}

	builder := InternshipLocation()
	res, err := builder.FromMap(data)

	assert.Nil(t, err)
	assert.Equal(t, id, res.ID())
	assert.Equal(t, "Location A", res.Name())
	assert.Equal(t, "123", res.Number())
	assert.Equal(t, "Street A", res.Street())
	assert.Equal(t, "Neighborhood A", res.Neighborhood())
	assert.Equal(t, "City A", res.City())
	assert.Equal(t, "12345", res.ZipCode())
}

func TestInternshipLocationQueryObject_FromMap_WithLatLong(t *testing.T) {
	id := uuid.New()
	latBytes := []uint8("123.456")
	longBytes := []uint8("789.012")
	data := map[string]interface{}{
		query.InternshipLocationID:           []uint8(id.String()),
		query.InternshipLocationName:         "Location A",
		query.InternshipLocationNumber:       "123",
		query.InternshipLocationStreet:       "Street A",
		query.InternshipLocationNeighborhood: "Neighborhood A",
		query.InternshipLocationCity:         "City A",
		query.InternshipLocationZipCode:      "12345",
		query.InternshipLocationLat:          latBytes,
		query.InternshipLocationLong:         longBytes,
	}

	builder := InternshipLocation()
	res, err := builder.FromMap(data)

	assert.Nil(t, err)
	assert.Equal(t, id, res.ID())
	assert.Equal(t, 123.456, res.Lat())
	assert.Equal(t, 789.012, res.Long())
}

func TestInternshipLocationQueryObject_FromMap_InvalidUUID(t *testing.T) {
	data := map[string]interface{}{
		query.InternshipLocationID:           []uint8("invalid-uuid"),
		query.InternshipLocationName:         "Location A",
		query.InternshipLocationNumber:       "123",
		query.InternshipLocationStreet:       "Street A",
		query.InternshipLocationNeighborhood: "Neighborhood A",
		query.InternshipLocationCity:         "City A",
		query.InternshipLocationZipCode:      "12345",
	}

	builder := InternshipLocation()
	res, err := builder.FromMap(data)

	assert.NotNil(t, err)
	assert.Nil(t, res)
}

func TestInternshipLocationQueryObject_FromRows_NilRows(t *testing.T) {
	builder := InternshipLocation()
	res, err := builder.FromRows(nil)

	assert.NotNil(t, err)
	assert.Nil(t, res)
}
