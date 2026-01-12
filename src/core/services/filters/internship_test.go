package filters

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestInternshipFilters(t *testing.T) {
	id := uuid.New()
	f := InternshipFilters{
		StudentID: &id,
	}

	assert.Equal(t, &id, f.StudentID)
}
