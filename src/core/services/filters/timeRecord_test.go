package filters

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestTimeRecordFilters(t *testing.T) {
	studentID := uuid.New()
	teacherID := uuid.New()
	statusID := uuid.New()
	startDate := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2024, 1, 31, 23, 59, 59, 0, time.UTC)
	search := "test"
	f := TimeRecordFilters{
		StudentID: &studentID,
		StartDate: &startDate,
		EndDate:   &endDate,
		TeacherID: &teacherID,
		StatusID:  &statusID,
		Search:    &search,
	}

	assert.Equal(t, &studentID, f.StudentID)
	assert.Equal(t, &startDate, f.StartDate)
	assert.Equal(t, &endDate, f.EndDate)
	assert.Equal(t, &teacherID, f.TeacherID)
	assert.Equal(t, &statusID, f.StatusID)
	assert.Equal(t, &search, f.Search)
}
