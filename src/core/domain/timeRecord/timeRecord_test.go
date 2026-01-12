package timeRecord

import (
	"eletronic_point/src/core/domain/internship"
	"eletronic_point/src/core/domain/internshipLocation"
	"eletronic_point/src/core/domain/simplifiedStudent"
	"eletronic_point/src/core/domain/timeRecordStatus"
	"eletronic_point/src/core/messages"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestTimeRecord_Setters(t *testing.T) {
	id := uuid.New()
	date := time.Now()
	entry := time.Now()
	t_rec := &timeRecord{}

	err := t_rec.SetID(id)
	assert.Nil(t, err)
	assert.Equal(t, id, t_rec.ID())

	err = t_rec.SetDate(date)
	assert.Nil(t, err)
	assert.Equal(t, date, t_rec.Date())

	err = t_rec.SetEntryTime(entry)
	assert.Nil(t, err)
	assert.Equal(t, entry, t_rec.EntryTime())

	// Test ExitTime
	exit := time.Now().Add(8 * time.Hour)
	err = t_rec.SetExitTime(&exit)
	assert.Nil(t, err)
	assert.Equal(t, &exit, t_rec.ExitTime())

	// Test Location
	err = t_rec.SetLocation("Office A")
	assert.Nil(t, err)
	assert.Equal(t, "Office A", t_rec.Location())

	// Test IsOffSite
	err = t_rec.SetIsOffSite(true)
	assert.Nil(t, err)
	assert.True(t, t_rec.IsOffSite())

	err = t_rec.SetIsOffSite(false)
	assert.Nil(t, err)
	assert.False(t, t_rec.IsOffSite())

	// Test Justification
	justification := "Doctor appointment"
	err = t_rec.SetJustification(&justification)
	assert.Nil(t, err)
	assert.Equal(t, &justification, t_rec.Justification())

	// Test Justification nil
	err = t_rec.SetJustification(nil)
	assert.Nil(t, err)
	assert.Nil(t, t_rec.Justification())

	// Test StudentID
	studentID := uuid.New()
	err = t_rec.SetStudentID(studentID)
	assert.Nil(t, err)
	assert.Equal(t, studentID, t_rec.StudentID())

	// Test InternshipID
	internshipID := uuid.New()
	err = t_rec.SetInternshipID(internshipID)
	assert.Nil(t, err)
	assert.Equal(t, internshipID, t_rec.InternshipID())
}

func TestTimeRecord_Setters_Errors(t *testing.T) {
	t_rec := &timeRecord{}

	err := t_rec.SetID(uuid.Nil)
	assert.NotNil(t, err)
	assert.Contains(t, err.Messages(), messages.TimeRecordIDErrorMessage)

	err = t_rec.SetDate(time.Time{})
	assert.NotNil(t, err)
	assert.Contains(t, err.Messages(), messages.TimeRecordDateErrorMessage)

	err = t_rec.SetEntryTime(time.Time{})
	assert.NotNil(t, err)
	assert.Contains(t, err.Messages(), messages.TimeRecordEntryTimeErrorMessage)

	err = t_rec.SetLocation("")
	assert.NotNil(t, err)
	assert.Contains(t, err.Messages(), messages.TimeRecordLocationErrorMessage)

	err = t_rec.SetStudentID(uuid.Nil)
	assert.NotNil(t, err)
	assert.Contains(t, err.Messages(), messages.TimeRecordStudentIDErrorMessage)

	err = t_rec.SetInternshipID(uuid.Nil)
	assert.NotNil(t, err)
	assert.Contains(t, err.Messages(), messages.InternshipIDErrorMessage)
}

func TestTimeRecord_ComplexSetters(t *testing.T) {
	t_rec := &timeRecord{}

	// Test SetStudent
	studentID := uuid.New()
	student, _ := simplifiedStudent.NewBuilder().WithID(studentID).WithName("John Doe").Build()
	err := t_rec.SetStudent(student)
	assert.Nil(t, err)
	assert.Equal(t, student, t_rec.Student())

	// Test SetInternship
	internshipID := uuid.New()
	startTime := time.Now()
	location, _ := internshipLocation.NewBuilder().WithID(uuid.New()).WithName("Location").WithNumber("123").WithStreet("Street").WithNeighborhood("Neighborhood").WithCity("City").WithZipCode("12345").Build()
	intern, _ := internship.NewBuilder().WithID(internshipID).WithStartedIn(startTime).WithLocation(location).Build()
	err = t_rec.SetInternship(intern)
	assert.Nil(t, err)
	assert.Equal(t, intern, t_rec.Internship())

	// Test SetTimeRecordStatus
	statusID := uuid.New()
	status, _ := timeRecordStatus.NewBuilder().WithID(statusID).WithName("Pending").Build()
	err = t_rec.SetTimeRecordStatus(status)
	assert.Nil(t, err)
	assert.Equal(t, status, t_rec.TimeRecordStatus())
}

func TestTimeRecord_TimeRecordStatus_Errors(t *testing.T) {
	t_rec := &timeRecord{}

	err := t_rec.SetTimeRecordStatus(nil)
	assert.NotNil(t, err)
	assert.Contains(t, err.Messages(), messages.TimeRecordStatusErrorMessage)
}

func TestTimeRecord_Getters(t *testing.T) {
	id := uuid.New()
	date := time.Now()
	entry := time.Now()
	exit := time.Now().Add(8 * time.Hour)
	location := "Office A"
	isOffSite := true
	justification := "Appointment"
	studentID := uuid.New()
	internshipID := uuid.New()

	t_rec := &timeRecord{
		id:            id,
		date:          date,
		entryTime:     entry,
		exitTime:      &exit,
		location:      location,
		isOffSite:     isOffSite,
		justification: &justification,
		studentID:     studentID,
		internshipID:  internshipID,
	}

	assert.Equal(t, id, t_rec.ID())
	assert.Equal(t, date, t_rec.Date())
	assert.Equal(t, entry, t_rec.EntryTime())
	assert.Equal(t, &exit, t_rec.ExitTime())
	assert.Equal(t, location, t_rec.Location())
	assert.Equal(t, isOffSite, t_rec.IsOffSite())
	assert.Equal(t, &justification, t_rec.Justification())
	assert.Equal(t, studentID, t_rec.StudentID())
	assert.Equal(t, internshipID, t_rec.InternshipID())
}
