package queryObject

import (
	"eletronic_point/src/infra/repository/postgres/query"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestTimeRecordQueryObject_FromMap(t *testing.T) {
	timeRecordID := uuid.New()
	studentID := uuid.New()
	internshipID := uuid.New()
	statusID := uuid.New()
	instID := uuid.New()
	campusID := uuid.New()
	courseID := uuid.New()
	locationID := uuid.New()
	date := time.Now()
	entryTime := time.Now()
	dateStr := date.Format("2006-01-02 15:04:05 -0700 -0700")
	entryStr := entryTime.Format("2006-01-02 15:04:05 -0700 -0700")

	data := map[string]interface{}{
		query.TimeRecordID:           []uint8(timeRecordID.String()),
		query.TimeRecordDate:         dateStr,
		query.TimeRecordEntryTime:    entryStr,
		query.TimeRecordLocation:     "Office A",
		query.TimeRecordIsOffSite:    false,
		query.TimeRecordStudentID:    []uint8(studentID.String()),
		query.TimeRecordInternshipID: []uint8(internshipID.String()),
		query.TimeRecordStatusID:     []uint8(statusID.String()),
		query.TimeRecordStatusName:   "Pending",
		query.PersonName:             "John Doe",
		query.StudentTotalWorkload:   int64(100),
		query.InstitutionID:          []uint8(instID.String()),
		query.InstitutionName:        "Institution A",
		query.CampusID:               []uint8(campusID.String()),
		query.CampusName:             "Campus A",
		query.CourseID:               []uint8(courseID.String()),
		query.CourseName:             "Course A",
		query.InternshipStartedIn:    dateStr,
		query.InternshipLocationID:   []uint8(locationID.String()),
		query.InternshipLocationName: "Location A",
	}

	builder := TimeRecord()
	res, err := builder.FromMap(data)

	assert.Nil(t, err)
	assert.Equal(t, timeRecordID, res.ID())
	assert.Equal(t, date.Format("2006-01-02"), res.Date().Format("2006-01-02"))
	assert.Equal(t, entryTime.Format("15:04:05"), res.EntryTime().Format("15:04:05"))
	assert.Equal(t, "Office A", res.Location())
	assert.Equal(t, studentID, res.StudentID())
	assert.Equal(t, internshipID, res.InternshipID())
}

func TestTimeRecordQueryObject_FromMap_WithExitTime(t *testing.T) {
	timeRecordID := uuid.New()
	studentID := uuid.New()
	internshipID := uuid.New()
	statusID := uuid.New()
	instID := uuid.New()
	campusID := uuid.New()
	courseID := uuid.New()
	locationID := uuid.New()
	date := time.Now()
	entryTime := time.Now()
	exitTime := entryTime.Add(8 * time.Hour)
	dateStr := date.Format("2006-01-02 15:04:05 -0700 -0700")
	entryStr := entryTime.Format("2006-01-02 15:04:05 -0700 -0700")

	data := map[string]interface{}{
		query.TimeRecordID:           []uint8(timeRecordID.String()),
		query.TimeRecordDate:         dateStr,
		query.TimeRecordEntryTime:    entryStr,
		query.TimeRecordExitTime:     exitTime,
		query.TimeRecordLocation:     "Office A",
		query.TimeRecordIsOffSite:    false,
		query.TimeRecordStudentID:    []uint8(studentID.String()),
		query.TimeRecordInternshipID: []uint8(internshipID.String()),
		query.TimeRecordStatusID:     []uint8(statusID.String()),
		query.TimeRecordStatusName:   "Pending",
		query.PersonName:             "John Doe",
		query.StudentTotalWorkload:   int64(100),
		query.InstitutionID:          []uint8(instID.String()),
		query.InstitutionName:        "Institution A",
		query.CampusID:               []uint8(campusID.String()),
		query.CampusName:             "Campus A",
		query.CourseID:               []uint8(courseID.String()),
		query.CourseName:             "Course A",
		query.InternshipStartedIn:    dateStr,
		query.InternshipLocationID:   []uint8(locationID.String()),
		query.InternshipLocationName: "Location A",
	}

	builder := TimeRecord()
	res, err := builder.FromMap(data)

	assert.Nil(t, err)
	assert.NotNil(t, res.ExitTime())
	assert.Equal(t, exitTime.Format("15:04:05"), res.ExitTime().Format("15:04:05"))
}

func TestTimeRecordQueryObject_FromMap_InvalidUUID(t *testing.T) {
	studentID := uuid.New()
	internshipID := uuid.New()
	statusID := uuid.New()
	instID := uuid.New()
	campusID := uuid.New()
	courseID := uuid.New()
	locationID := uuid.New()
	date := time.Now()
	entryTime := time.Now()
	dateStr := date.Format("2006-01-02 15:04:05 -0700 -0700")
	entryStr := entryTime.Format("2006-01-02 15:04:05 -0700 -0700")

	data := map[string]interface{}{
		query.TimeRecordID:           []uint8("invalid-uuid"),
		query.TimeRecordDate:         dateStr,
		query.TimeRecordEntryTime:    entryStr,
		query.TimeRecordLocation:     "Office A",
		query.TimeRecordIsOffSite:    false,
		query.TimeRecordStudentID:    []uint8(studentID.String()),
		query.TimeRecordInternshipID: []uint8(internshipID.String()),
		query.TimeRecordStatusID:     []uint8(statusID.String()),
		query.TimeRecordStatusName:   "Pending",
		query.PersonName:             "John Doe",
		query.StudentTotalWorkload:   int64(100),
		query.InstitutionID:          []uint8(instID.String()),
		query.InstitutionName:        "Institution A",
		query.CampusID:               []uint8(campusID.String()),
		query.CampusName:             "Campus A",
		query.CourseID:               []uint8(courseID.String()),
		query.CourseName:             "Course A",
		query.InternshipStartedIn:    dateStr,
		query.InternshipLocationID:   []uint8(locationID.String()),
		query.InternshipLocationName: "Location A",
	}

	builder := TimeRecord()
	res, err := builder.FromMap(data)

	assert.NotNil(t, err)
	assert.Nil(t, res)
}

func TestTimeRecordQueryObject_FromMap_InvalidDate(t *testing.T) {
	timeRecordID := uuid.New()
	studentID := uuid.New()
	internshipID := uuid.New()
	statusID := uuid.New()
	instID := uuid.New()
	campusID := uuid.New()
	courseID := uuid.New()
	locationID := uuid.New()
	entryTime := time.Now()
	entryStr := entryTime.Format("2006-01-02 15:04:05 -0700 -0700")

	data := map[string]interface{}{
		query.TimeRecordID:           []uint8(timeRecordID.String()),
		query.TimeRecordDate:         "invalid-date",
		query.TimeRecordEntryTime:    entryStr,
		query.TimeRecordLocation:     "Office A",
		query.TimeRecordIsOffSite:    false,
		query.TimeRecordStudentID:    []uint8(studentID.String()),
		query.TimeRecordInternshipID: []uint8(internshipID.String()),
		query.TimeRecordStatusID:     []uint8(statusID.String()),
		query.TimeRecordStatusName:   "Pending",
		query.PersonName:             "John Doe",
		query.StudentTotalWorkload:   int64(100),
		query.InstitutionID:          []uint8(instID.String()),
		query.InstitutionName:        "Institution A",
		query.CampusID:               []uint8(campusID.String()),
		query.CampusName:             "Campus A",
		query.CourseID:               []uint8(courseID.String()),
		query.CourseName:             "Course A",
		query.InternshipStartedIn:    entryStr,
		query.InternshipLocationID:   []uint8(locationID.String()),
		query.InternshipLocationName: "Location A",
	}

	builder := TimeRecord()
	res, err := builder.FromMap(data)

	assert.NotNil(t, err)
	assert.Nil(t, res)
}

func TestTimeRecordQueryObject_FromRows_NilRows(t *testing.T) {
	builder := TimeRecord()
	res, err := builder.FromRows(nil)

	assert.NotNil(t, err)
	assert.Nil(t, res)
}
