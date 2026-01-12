package postgres

import (
	"eletronic_point/src/core/domain/campus"
	"eletronic_point/src/core/domain/institution"
	"eletronic_point/src/core/services/filters"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCampusRepository_Integration(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	SetupTestDB(t)
	CleanDB(t)

	repo := NewCampusRepository()
	instRepo := NewInstitutionRepository()

	t.Run("Create and Get", func(t *testing.T) {
		// Create institution first
		instID := uuid.New()
		inst, _ := institution.NewBuilder().WithID(instID).WithName("Test Institution").Build()
		createdInstID, err := instRepo.Create(inst)
		assert.Nil(t, err)
		assert.NotNil(t, createdInstID)

		// Create campus
		campusID := uuid.New()
		c, _ := campus.NewBuilder().WithID(campusID).WithName("Test Campus").WithInstitutionID(*createdInstID).Build()
		createdID, err := repo.Create(c)
		assert.Nil(t, err)
		assert.NotNil(t, createdID)

		// Get campus
		found, err := repo.Get(*createdID)
		assert.Nil(t, err)
		assert.Equal(t, "Test Campus", found.Name())
		assert.Equal(t, *createdInstID, found.InstitutionID())
	})

	t.Run("List", func(t *testing.T) {
		// Create institution first
		instID := uuid.New()
		inst, _ := institution.NewBuilder().WithID(instID).WithName("Test Institution 2").Build()
		createdInstID, err := instRepo.Create(inst)
		assert.Nil(t, err)

		// Create campus
		campusID := uuid.New()
		c, _ := campus.NewBuilder().WithID(campusID).WithName("Test Campus 2").WithInstitutionID(*createdInstID).Build()
		_, err = repo.Create(c)
		assert.Nil(t, err)

		// List campuses
		campuses, err := repo.List(filters.CampusFilters{})
		assert.Nil(t, err)
		assert.NotEmpty(t, campuses)
	})

	t.Run("Update", func(t *testing.T) {
		// Create institution first
		instID := uuid.New()
		inst, _ := institution.NewBuilder().WithID(instID).WithName("Test Institution 3").Build()
		createdInstID, err := instRepo.Create(inst)
		assert.Nil(t, err)

		// Create campus
		campusID := uuid.New()
		c, _ := campus.NewBuilder().WithID(campusID).WithName("Test Campus 3").WithInstitutionID(*createdInstID).Build()
		createdID, err := repo.Create(c)
		assert.Nil(t, err)

		// Update campus
		updated, err := repo.Get(*createdID)
		assert.Nil(t, err)
		
		updatedCampus, _ := campus.NewBuilder().
			WithID(updated.ID()).
			WithName("Updated Campus").
			WithInstitutionID(updated.InstitutionID()).
			Build()

		err = repo.Update(updatedCampus)
		assert.Nil(t, err)

		// Verify update
		found, err := repo.Get(*createdID)
		assert.Nil(t, err)
		assert.Equal(t, "Updated Campus", found.Name())
	})

	t.Run("Delete", func(t *testing.T) {
		// Create institution first
		instID := uuid.New()
		inst, _ := institution.NewBuilder().WithID(instID).WithName("Test Institution 4").Build()
		createdInstID, err := instRepo.Create(inst)
		assert.Nil(t, err)

		// Create campus
		campusID := uuid.New()
		c, _ := campus.NewBuilder().WithID(campusID).WithName("Test Campus 4").WithInstitutionID(*createdInstID).Build()
		createdID, err := repo.Create(c)
		assert.Nil(t, err)

		// Delete campus
		err = repo.Delete(*createdID)
		assert.Nil(t, err)

		// Verify deletion
		_, err = repo.Get(*createdID)
		assert.NotNil(t, err)
	})
}
