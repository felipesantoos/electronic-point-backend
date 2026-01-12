package queryObject

import (
	"eletronic_point/src/infra/repository/postgres/query"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCampusQueryObject_FromMap(t *testing.T) {
	id := uuid.New()
	instID := uuid.New()
	data := map[string]interface{}{
		query.CampusID:            []uint8(id.String()),
		query.CampusName:          "Campus A",
		query.CampusInstitutionID: []uint8(instID.String()),
	}

	builder := Campus()
	res, err := builder.FromMap(data)

	assert.Nil(t, err)
	assert.Equal(t, id, res.ID())
	assert.Equal(t, "Campus A", res.Name())
	assert.Equal(t, instID, res.InstitutionID())
}

func TestCampusQueryObject_FromMap_InvalidUUID(t *testing.T) {
	instID := uuid.New()
	data := map[string]interface{}{
		query.CampusID:            []uint8("invalid-uuid"),
		query.CampusName:          "Campus A",
		query.CampusInstitutionID: []uint8(instID.String()),
	}

	builder := Campus()
	res, err := builder.FromMap(data)

	assert.NotNil(t, err)
	assert.Nil(t, res)
}

func TestCampusQueryObject_FromMap_InvalidInstitutionID(t *testing.T) {
	id := uuid.New()
	data := map[string]interface{}{
		query.CampusID:            []uint8(id.String()),
		query.CampusName:          "Campus A",
		query.CampusInstitutionID: []uint8("invalid-uuid"),
	}

	builder := Campus()
	res, err := builder.FromMap(data)

	assert.NotNil(t, err)
	assert.Nil(t, res)
}

func TestCampusQueryObject_FromRows_NilRows(t *testing.T) {
	builder := Campus()
	res, err := builder.FromRows(nil)

	assert.NotNil(t, err)
	assert.Nil(t, res)
}
