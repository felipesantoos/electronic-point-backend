package queryObject

import (
	"eletronic_point/src/infra/repository/postgres/query"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestTimeRecordStatusQueryObject_FromMap(t *testing.T) {
	id := uuid.New()
	data := map[string]interface{}{
		query.TimeRecordStatusID:   []uint8(id.String()),
		query.TimeRecordStatusName: "Pending",
	}

	builder := TimeRecordStatus()
	res, err := builder.FromMap(data)

	assert.Nil(t, err)
	assert.Equal(t, id, res.ID())
	assert.Equal(t, "Pending", res.Name())
}

func TestTimeRecordStatusQueryObject_FromMap_InvalidUUID(t *testing.T) {
	data := map[string]interface{}{
		query.TimeRecordStatusID:   []uint8("invalid-uuid"),
		query.TimeRecordStatusName: "Pending",
	}

	builder := TimeRecordStatus()
	res, err := builder.FromMap(data)

	assert.NotNil(t, err)
	assert.Nil(t, res)
}

func TestTimeRecordStatusQueryObject_FromRows_NilRows(t *testing.T) {
	builder := TimeRecordStatus()
	res, err := builder.FromRows(nil)

	assert.NotNil(t, err)
	assert.Nil(t, res)
}
