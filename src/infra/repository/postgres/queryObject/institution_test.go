package queryObject

import (
	"eletronic_point/src/infra/repository/postgres/query"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestInstitutionQueryObject_FromMap(t *testing.T) {
	id := uuid.New()
	data := map[string]interface{}{
		query.InstitutionID:   []uint8(id.String()),
		query.InstitutionName: "Institution A",
	}

	builder := Institution()
	res, err := builder.FromMap(data)

	assert.Nil(t, err)
	assert.Equal(t, id, res.ID())
	assert.Equal(t, "Institution A", res.Name())
}

func TestInstitutionQueryObject_FromMap_InvalidUUID(t *testing.T) {
	data := map[string]interface{}{
		query.InstitutionID:   []uint8("invalid-uuid"),
		query.InstitutionName: "Institution A",
	}

	builder := Institution()
	res, err := builder.FromMap(data)

	assert.NotNil(t, err)
	assert.Nil(t, res)
}

func TestInstitutionQueryObject_FromRows_NilRows(t *testing.T) {
	builder := Institution()
	res, err := builder.FromRows(nil)

	assert.NotNil(t, err)
	assert.Nil(t, res)
}
