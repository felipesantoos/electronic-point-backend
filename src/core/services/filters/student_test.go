package filters

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestStudentFilters(t *testing.T) {
	teacherID := uuid.New()
	instID := uuid.New()
	campusID := uuid.New()
	search := "test"
	f := StudentFilters{
		TeacherID:     &teacherID,
		InstitutionID: &instID,
		CampusID:      &campusID,
		Search:        &search,
	}

	assert.Equal(t, &teacherID, f.TeacherID)
	assert.Equal(t, &instID, f.InstitutionID)
	assert.Equal(t, &campusID, f.CampusID)
	assert.Equal(t, &search, f.Search)
}
