package queryObject

import (
	"eletronic_point/src/infra/repository/postgres/query"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestInternshipQueryObject_FromMap(t *testing.T) {
	internshipID := uuid.New()
	locationID := uuid.New()
	studentID := uuid.New()
	instID := uuid.New()
	campusID := uuid.New()
	courseID := uuid.New()
	startedIn := time.Now()
	startedInStr := startedIn.Format("2006-01-02 15:04:05 -0700 -0700")

	data := map[string]interface{}{
		query.InternshipID:                   []uint8(internshipID.String()),
		query.InternshipStartedIn:            startedInStr,
		query.InternshipLocationID:           []uint8(locationID.String()),
		query.InternshipLocationName:         "Location A",
		query.InternshipLocationNumber:       "123",
		query.InternshipLocationStreet:       "Street A",
		query.InternshipLocationNeighborhood: "Neighborhood A",
		query.InternshipLocationCity:         "City A",
		query.InternshipLocationZipCode:      "12345",
		query.PersonID:                       []uint8(studentID.String()),
		query.PersonName:                     "John Doe",
		query.StudentTotalWorkload:           int64(100),
		query.InstitutionID:                  []uint8(instID.String()),
		query.InstitutionName:                "Institution A",
		query.CampusID:                       []uint8(campusID.String()),
		query.CampusName:                     "Campus A",
		query.CourseID:                       []uint8(courseID.String()),
		query.CourseName:                     "Course A",
	}

	builder := Internship()
	res, err := builder.FromMap(data, false)

	assert.Nil(t, err)
	assert.Equal(t, internshipID, res.ID())
	assert.Equal(t, startedIn.Format("2006-01-02"), res.StartedIn().Format("2006-01-02"))
	assert.NotNil(t, res.Location())
}

func TestInternshipQueryObject_FromMap_WithStudentInfo(t *testing.T) {
	internshipID := uuid.New()
	locationID := uuid.New()
	studentID := uuid.New()
	instID := uuid.New()
	campusID := uuid.New()
	courseID := uuid.New()
	startedIn := time.Now()
	startedInStr := startedIn.Format("2006-01-02 15:04:05 -0700 -0700")

	data := map[string]interface{}{
		query.InternshipID:                   []uint8(internshipID.String()),
		query.InternshipStartedIn:            startedInStr,
		query.InternshipLocationID:           []uint8(locationID.String()),
		query.InternshipLocationName:         "Location A",
		query.InternshipLocationNumber:       "123",
		query.InternshipLocationStreet:       "Street A",
		query.InternshipLocationNeighborhood: "Neighborhood A",
		query.InternshipLocationCity:         "City A",
		query.InternshipLocationZipCode:      "12345",
		query.PersonID:                       []uint8(studentID.String()),
		query.PersonName:                     "John Doe",
		query.StudentTotalWorkload:           int64(100),
		query.InstitutionID:                  []uint8(instID.String()),
		query.InstitutionName:                "Institution A",
		query.CampusID:                       []uint8(campusID.String()),
		query.CampusName:                     "Campus A",
		query.CourseID:                       []uint8(courseID.String()),
		query.CourseName:                     "Course A",
	}

	builder := Internship()
	res, err := builder.FromMap(data, true)

	assert.Nil(t, err)
	assert.Equal(t, internshipID, res.ID())
	assert.NotNil(t, res.Student())
	assert.Equal(t, studentID, *res.Student().ID())
}

func TestInternshipQueryObject_FromMap_InvalidUUID(t *testing.T) {
	locationID := uuid.New()
	studentID := uuid.New()
	instID := uuid.New()
	campusID := uuid.New()
	courseID := uuid.New()
	startedIn := time.Now()
	startedInStr := startedIn.Format("2006-01-02 15:04:05 -0700 -0700")

	data := map[string]interface{}{
		query.InternshipID:                   []uint8("invalid-uuid"),
		query.InternshipStartedIn:            startedInStr,
		query.InternshipLocationID:           []uint8(locationID.String()),
		query.InternshipLocationName:         "Location A",
		query.InternshipLocationNumber:       "123",
		query.InternshipLocationStreet:       "Street A",
		query.InternshipLocationNeighborhood: "Neighborhood A",
		query.InternshipLocationCity:         "City A",
		query.InternshipLocationZipCode:      "12345",
		query.PersonID:                       []uint8(studentID.String()),
		query.PersonName:                     "John Doe",
		query.StudentTotalWorkload:           int64(100),
		query.InstitutionID:                  []uint8(instID.String()),
		query.InstitutionName:                "Institution A",
		query.CampusID:                       []uint8(campusID.String()),
		query.CampusName:                     "Campus A",
		query.CourseID:                       []uint8(courseID.String()),
		query.CourseName:                     "Course A",
	}

	builder := Internship()
	res, err := builder.FromMap(data, false)

	assert.NotNil(t, err)
	assert.Nil(t, res)
}

func TestInternshipQueryObject_FromMap_InvalidDate(t *testing.T) {
	internshipID := uuid.New()
	locationID := uuid.New()
	studentID := uuid.New()
	instID := uuid.New()
	campusID := uuid.New()
	courseID := uuid.New()

	data := map[string]interface{}{
		query.InternshipID:                   []uint8(internshipID.String()),
		query.InternshipStartedIn:            "invalid-date",
		query.InternshipLocationID:           []uint8(locationID.String()),
		query.InternshipLocationName:         "Location A",
		query.InternshipLocationNumber:       "123",
		query.InternshipLocationStreet:       "Street A",
		query.InternshipLocationNeighborhood: "Neighborhood A",
		query.InternshipLocationCity:         "City A",
		query.InternshipLocationZipCode:      "12345",
		query.PersonID:                       []uint8(studentID.String()),
		query.PersonName:                     "John Doe",
		query.StudentTotalWorkload:           int64(100),
		query.InstitutionID:                  []uint8(instID.String()),
		query.InstitutionName:                "Institution A",
		query.CampusID:                       []uint8(campusID.String()),
		query.CampusName:                     "Campus A",
		query.CourseID:                       []uint8(courseID.String()),
		query.CourseName:                     "Course A",
	}

	builder := Internship()
	res, err := builder.FromMap(data, false)

	assert.NotNil(t, err)
	assert.Nil(t, res)
}

func TestInternshipQueryObject_FromRows_NilRows(t *testing.T) {
	builder := Internship()
	res, err := builder.FromRows(nil, false)

	assert.NotNil(t, err)
	assert.Nil(t, res)
}
