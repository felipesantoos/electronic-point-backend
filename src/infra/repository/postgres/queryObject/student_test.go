package queryObject

import (
	"eletronic_point/src/infra/repository/postgres/query"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestStudentQueryObject_FromMap(t *testing.T) {
	personID := uuid.New()
	instID := uuid.New()
	campusID := uuid.New()
	courseID := uuid.New()

	data := map[string]interface{}{
		query.PersonID:             []uint8(personID.String()),
		query.PersonName:           "John Doe",
		query.PersonBirthDate:      "1990-01-01",
		query.PersonEmail:          "john@example.com",
		query.PersonCPF:            "11144477735",
		query.PersonPhone:          "82999999999",
		query.StudentRegistration:  "12345",
		query.InstitutionID:        []uint8(instID.String()),
		query.InstitutionName:      "Institution A",
		query.CampusID:             []uint8(campusID.String()),
		query.CampusName:           "Campus A",
		query.CourseID:             []uint8(courseID.String()),
		query.CourseName:           "Course A",
		query.StudentTotalWorkload: int64(100),
	}

	builder := Student()
	res, err := builder.FromMap(data)

	assert.Nil(t, err)
	assert.Equal(t, personID, *res.ID())
	assert.Equal(t, "John Doe", res.Name())
	assert.Equal(t, "12345", res.Registration())
	assert.Equal(t, 100, res.TotalWorkload())
}

func TestStudentQueryObject_FromMap_InvalidUUID(t *testing.T) {
	instID := uuid.New()
	campusID := uuid.New()
	courseID := uuid.New()

	data := map[string]interface{}{
		query.PersonID:             []uint8("invalid-uuid"),
		query.PersonName:           "John Doe",
		query.PersonBirthDate:      "1990-01-01",
		query.PersonEmail:          "john@example.com",
		query.PersonCPF:            "11144477735",
		query.PersonPhone:          "82999999999",
		query.StudentRegistration:  "12345",
		query.InstitutionID:        []uint8(instID.String()),
		query.InstitutionName:      "Institution A",
		query.CampusID:             []uint8(campusID.String()),
		query.CampusName:           "Campus A",
		query.CourseID:             []uint8(courseID.String()),
		query.CourseName:           "Course A",
		query.StudentTotalWorkload: int64(100),
	}

	builder := Student()
	res, err := builder.FromMap(data)

	assert.NotNil(t, err)
	assert.Nil(t, res)
}

func TestStudentQueryObject_FromRows_NilRows(t *testing.T) {
	builder := Student()
	res, err := builder.FromRows(nil)

	assert.NotNil(t, err)
	assert.Nil(t, res)
}
