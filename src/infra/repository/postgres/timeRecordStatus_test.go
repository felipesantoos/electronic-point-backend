package postgres

import (
	"eletronic_point/src/core/domain/timeRecordStatus"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestTimeRecordStatusRepository_Integration(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	SetupTestDB(t)
	CleanDB(t)

	repo := NewTimeRecordStatusRepository()

	t.Run("Create and Get", func(t *testing.T) {
		statusID := uuid.New()
		status, _ := timeRecordStatus.NewBuilder().WithID(statusID).WithName("Test Status").Build()
		createdID, err := repo.Create(status)
		assert.Nil(t, err)
		assert.NotNil(t, createdID)

		// Get status
		found, err := repo.Get(*createdID)
		assert.Nil(t, err)
		assert.Equal(t, "Test Status", found.Name())
	})

	t.Run("List", func(t *testing.T) {
		// List statuses
		statuses, err := repo.List()
		assert.Nil(t, err)
		assert.NotEmpty(t, statuses)
	})

	t.Run("Update", func(t *testing.T) {
		// Create status first
		statusID := uuid.New()
		status, _ := timeRecordStatus.NewBuilder().WithID(statusID).WithName("Test Status 2").Build()
		createdID, err := repo.Create(status)
		assert.Nil(t, err)

		// Update status
		updated, err := repo.Get(*createdID)
		assert.Nil(t, err)
		
		updatedStatus, _ := timeRecordStatus.NewBuilder().
			WithID(updated.ID()).
			WithName("Updated Status").
			Build()

		err = repo.Update(updatedStatus)
		assert.Nil(t, err)

		// Verify update
		found, err := repo.Get(*createdID)
		assert.Nil(t, err)
		assert.Equal(t, "Updated Status", found.Name())
	})

	t.Run("Delete", func(t *testing.T) {
		// Create status first
		statusID := uuid.New()
		status, _ := timeRecordStatus.NewBuilder().WithID(statusID).WithName("Test Status 3").Build()
		createdID, err := repo.Create(status)
		assert.Nil(t, err)

		// Delete status
		err = repo.Delete(*createdID)
		assert.Nil(t, err)

		// Verify deletion
		_, err = repo.Get(*createdID)
		assert.NotNil(t, err)
	})
}
