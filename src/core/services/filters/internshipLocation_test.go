package filters

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestInternshipLocationFilters(t *testing.T) {
	studentID := uuid.New()
	search := "test"
	f := InternshipLocationFilters{
		StudentID: &studentID,
		Search:    &search,
	}

	assert.Equal(t, &studentID, f.StudentID)
	assert.Equal(t, &search, f.Search)
}
