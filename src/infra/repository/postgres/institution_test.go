package postgres

import (
	"eletronic_point/src/core/domain/institution"
	"eletronic_point/src/core/services/filters"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestInstitutionRepository_Integration(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	SetupTestDB(t)
	CleanDB(t)

	repo := NewInstitutionRepository()

	t.Run("Create and Get", func(t *testing.T) {
		instID := uuid.New()
		inst, _ := institution.NewBuilder().WithID(instID).WithName("Test Institution").Build()
		createdID, err := repo.Create(inst)
		assert.Nil(t, err)
		assert.NotNil(t, createdID)

		// Get institution
		found, err := repo.Get(*createdID)
		assert.Nil(t, err)
		assert.Equal(t, "Test Institution", found.Name())
	})

	t.Run("List", func(t *testing.T) {
		// Create institution
		instID := uuid.New()
		inst, _ := institution.NewBuilder().WithID(instID).WithName("Test Institution 2").Build()
		_, err := repo.Create(inst)
		assert.Nil(t, err)

		// List institutions
		institutions, err := repo.List(filters.InstitutionFilters{})
		assert.Nil(t, err)
		assert.NotEmpty(t, institutions)
	})

	t.Run("Update", func(t *testing.T) {
		// Create institution
		instID := uuid.New()
		inst, _ := institution.NewBuilder().WithID(instID).WithName("Test Institution 3").Build()
		createdID, err := repo.Create(inst)
		assert.Nil(t, err)

		// Update institution
		updated, err := repo.Get(*createdID)
		assert.Nil(t, err)
		
		updatedInstitution, _ := institution.NewBuilder().
			WithID(updated.ID()).
			WithName("Updated Institution").
			Build()

		err = repo.Update(updatedInstitution)
		assert.Nil(t, err)

		// Verify update
		found, err := repo.Get(*createdID)
		assert.Nil(t, err)
		assert.Equal(t, "Updated Institution", found.Name())
	})

	t.Run("Delete", func(t *testing.T) {
		// Create institution
		instID := uuid.New()
		inst, _ := institution.NewBuilder().WithID(instID).WithName("Test Institution 4").Build()
		createdID, err := repo.Create(inst)
		assert.Nil(t, err)

		// Delete institution
		err = repo.Delete(*createdID)
		assert.Nil(t, err)

		// Verify deletion
		_, err = repo.Get(*createdID)
		assert.NotNil(t, err)
	})
}
