package internship

import (
	"eletronic_point/src/core/domain/internshipLocation"
	"eletronic_point/src/core/domain/simplifiedStudent"
	"eletronic_point/src/core/messages"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestInternship_Setters(t *testing.T) {
	id := uuid.New()
	start := time.Now()
	i := &internship{}

	err := i.SetID(id)
	assert.Nil(t, err)
	assert.Equal(t, id, i.ID())

	err = i.SetStartedIn(start)
	assert.Nil(t, err)
	assert.Equal(t, start, i.StartedIn())

	// Test SetEndedIn
	end := time.Now().Add(24 * 30 * time.Hour) // 30 days later
	err = i.SetEndedIn(&end)
	assert.Nil(t, err)
	assert.Equal(t, &end, i.EndedIn())

	// Test SetEndedIn nil
	err = i.SetEndedIn(nil)
	assert.Nil(t, err)
	assert.Nil(t, i.EndedIn())

	// Test SetScheduleEntryTime
	scheduleEntry := time.Date(2024, 1, 1, 8, 0, 0, 0, time.UTC)
	err = i.SetScheduleEntryTime(&scheduleEntry)
	assert.Nil(t, err)
	assert.Equal(t, &scheduleEntry, i.ScheduleEntryTime())

	// Test SetScheduleEntryTime nil
	err = i.SetScheduleEntryTime(nil)
	assert.Nil(t, err)
	assert.Nil(t, i.ScheduleEntryTime())

	// Test SetScheduleExitTime
	scheduleExit := time.Date(2024, 1, 1, 17, 0, 0, 0, time.UTC)
	err = i.SetScheduleExitTime(&scheduleExit)
	assert.Nil(t, err)
	assert.Equal(t, &scheduleExit, i.ScheduleExitTime())

	// Test SetScheduleExitTime nil
	err = i.SetScheduleExitTime(nil)
	assert.Nil(t, err)
	assert.Nil(t, i.ScheduleExitTime())
}

func TestInternship_Setters_Errors(t *testing.T) {
	i := &internship{}

	err := i.SetID(uuid.Nil)
	assert.NotNil(t, err)
	assert.Contains(t, err.Messages(), messages.InternshipIDErrorMessage)

	err = i.SetStartedIn(time.Time{})
	assert.NotNil(t, err)
	assert.Contains(t, err.Messages(), messages.InternshipStartedInErrorMessage)

	err = i.SetLocation(nil)
	assert.NotNil(t, err)
	assert.Contains(t, err.Messages(), messages.InternshipLocationErrorMessage)

	err = i.SetStudent(nil)
	assert.NotNil(t, err)
	assert.Contains(t, err.Messages(), messages.InternshipStudentErrorMessage)
}

func TestInternship_ComplexSetters(t *testing.T) {
	i := &internship{}

	// Test SetLocation with valid object
	locationID := uuid.New()
	location, _ := internshipLocation.NewBuilder().
		WithID(locationID).
		WithName("Location").
		WithNumber("123").
		WithStreet("Street").
		WithNeighborhood("Neighborhood").
		WithCity("City").
		WithZipCode("12345").
		Build()
	err := i.SetLocation(location)
	assert.Nil(t, err)
	assert.Equal(t, location, i.Location())

	// Test SetStudent with valid object
	studentID := uuid.New()
	student, _ := simplifiedStudent.NewBuilder().
		WithID(studentID).
		WithName("John Doe").
		Build()
	err = i.SetStudent(student)
	assert.Nil(t, err)
	assert.Equal(t, student, i.Student())
}

func TestInternship_Getters(t *testing.T) {
	id := uuid.New()
	start := time.Now()
	end := time.Now().Add(24 * 30 * time.Hour)
	scheduleEntry := time.Date(2024, 1, 1, 8, 0, 0, 0, time.UTC)
	scheduleExit := time.Date(2024, 1, 1, 17, 0, 0, 0, time.UTC)

	locationID := uuid.New()
	location, _ := internshipLocation.NewBuilder().
		WithID(locationID).
		WithName("Location").
		WithNumber("123").
		WithStreet("Street").
		WithNeighborhood("Neighborhood").
		WithCity("City").
		WithZipCode("12345").
		Build()

	studentID := uuid.New()
	student, _ := simplifiedStudent.NewBuilder().
		WithID(studentID).
		WithName("John Doe").
		Build()

	i := &internship{
		id:                id,
		startedIn:         start,
		endedIn:           &end,
		location:          location,
		_student:          student,
		scheduleEntryTime: &scheduleEntry,
		scheduleExitTime:  &scheduleExit,
	}

	assert.Equal(t, id, i.ID())
	assert.Equal(t, start, i.StartedIn())
	assert.Equal(t, &end, i.EndedIn())
	assert.Equal(t, location, i.Location())
	assert.Equal(t, student, i.Student())
	assert.Equal(t, &scheduleEntry, i.ScheduleEntryTime())
	assert.Equal(t, &scheduleExit, i.ScheduleExitTime())
}
