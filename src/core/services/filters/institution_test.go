package filters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInstitutionFilters(t *testing.T) {
	name := "Institution A"
	f := InstitutionFilters{
		Name: &name,
	}

	assert.Equal(t, &name, f.Name)
}
