package filters

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCampusFilters(t *testing.T) {
	name := "Campus A"
	instID := uuid.New()
	f := CampusFilters{
		Name:          &name,
		InstitutionID: &instID,
	}

	assert.Equal(t, &name, f.Name)
	assert.Equal(t, &instID, f.InstitutionID)
}
