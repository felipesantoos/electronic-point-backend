package filters

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestAccountFilters(t *testing.T) {
	id := uuid.New()
	search := "test"
	f := AccountFilters{
		RoleID: &id,
		Search: &search,
	}

	assert.Equal(t, &id, f.RoleID)
	assert.Equal(t, &search, f.Search)
}
