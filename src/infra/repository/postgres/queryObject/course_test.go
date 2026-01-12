package queryObject

import (
	"eletronic_point/src/infra/repository/postgres/query"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCourseQueryObject_FromMap(t *testing.T) {
	id := uuid.New()
	data := map[string]interface{}{
		query.CourseID:   []uint8(id.String()),
		query.CourseName: "Course A",
	}

	builder := Course()
	res, err := builder.FromMap(data)

	assert.Nil(t, err)
	assert.Equal(t, id, res.ID())
	assert.Equal(t, "Course A", res.Name())
}

func TestCourseQueryObject_FromMap_InvalidUUID(t *testing.T) {
	data := map[string]interface{}{
		query.CourseID:   []uint8("invalid-uuid"),
		query.CourseName: "Course A",
	}

	builder := Course()
	res, err := builder.FromMap(data)

	assert.NotNil(t, err)
	assert.Nil(t, res)
}

func TestCourseQueryObject_FromRows_NilRows(t *testing.T) {
	builder := Course()
	res, err := builder.FromRows(nil)

	assert.NotNil(t, err)
	assert.Nil(t, res)
}
