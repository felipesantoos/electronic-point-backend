package postgres

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResourcesRepository_Integration(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	SetupTestDB(t)
	CleanDB(t)

	repo := NewResourcesRepository()

	t.Run("ListAccountRoles", func(t *testing.T) {
		roles, err := repo.ListAccountRoles()
		assert.Nil(t, err)
		assert.NotEmpty(t, roles)
		// Verify at least common roles exist
		roleCodes := make(map[string]bool)
		for _, role := range roles {
			roleCodes[role.Code()] = true
		}
		// Just verify we got some roles back
		assert.True(t, len(roleCodes) > 0)
	})
}
