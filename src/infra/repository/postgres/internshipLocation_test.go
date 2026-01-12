package postgres

import (
	"eletronic_point/src/core/domain/internshipLocation"
	"eletronic_point/src/core/services/filters"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestInternshipLocationRepository_Integration(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	SetupTestDB(t)
	CleanDB(t)

	repo := NewInternshipLocationRepository()

	t.Run("Create and Get", func(t *testing.T) {
		locationID := uuid.New()
		location, _ := internshipLocation.NewBuilder().
			WithID(locationID).
			WithName("Test Location").
			WithNumber("123").
			WithStreet("Test Street").
			WithNeighborhood("Test Neighborhood").
			WithCity("Test City").
			WithZipCode("12345").
			Build()
		
		createdID, err := repo.Create(location)
		assert.Nil(t, err)
		assert.NotNil(t, createdID)

		// Get location
		found, err := repo.Get(*createdID)
		assert.Nil(t, err)
		assert.Equal(t, "Test Location", found.Name())
		assert.Equal(t, "123", found.Number())
		assert.Equal(t, "Test Street", found.Street())
	})

	t.Run("List", func(t *testing.T) {
		// Create location
		locationID := uuid.New()
		location, _ := internshipLocation.NewBuilder().
			WithID(locationID).
			WithName("Test Location 2").
			WithNumber("456").
			WithStreet("Test Street 2").
			WithNeighborhood("Test Neighborhood 2").
			WithCity("Test City 2").
			WithZipCode("54321").
			Build()
		
		_, err := repo.Create(location)
		assert.Nil(t, err)

		// List locations
		locations, err := repo.List(filters.InternshipLocationFilters{})
		assert.Nil(t, err)
		assert.NotEmpty(t, locations)
	})

	t.Run("Update", func(t *testing.T) {
		// Create location first
		locationID := uuid.New()
		location, _ := internshipLocation.NewBuilder().
			WithID(locationID).
			WithName("Test Location 3").
			WithNumber("789").
			WithStreet("Test Street 3").
			WithNeighborhood("Test Neighborhood 3").
			WithCity("Test City 3").
			WithZipCode("98765").
			Build()
		
		createdID, err := repo.Create(location)
		assert.Nil(t, err)

		// Update location
		updated, err := repo.Get(*createdID)
		assert.Nil(t, err)
		
		updatedLocation, _ := internshipLocation.NewBuilder().
			WithID(updated.ID()).
			WithName("Updated Location").
			WithNumber(updated.Number()).
			WithStreet(updated.Street()).
			WithNeighborhood(updated.Neighborhood()).
			WithCity(updated.City()).
			WithZipCode(updated.ZipCode()).
			Build()

		err = repo.Update(updatedLocation)
		assert.Nil(t, err)

		// Verify update
		found, err := repo.Get(*createdID)
		assert.Nil(t, err)
		assert.Equal(t, "Updated Location", found.Name())
	})

	t.Run("Delete", func(t *testing.T) {
		// Create location first
		locationID := uuid.New()
		location, _ := internshipLocation.NewBuilder().
			WithID(locationID).
			WithName("Test Location 4").
			WithNumber("000").
			WithStreet("Test Street 4").
			WithNeighborhood("Test Neighborhood 4").
			WithCity("Test City 4").
			WithZipCode("00000").
			Build()
		
		createdID, err := repo.Create(location)
		assert.Nil(t, err)

		// Delete location
		err = repo.Delete(*createdID)
		assert.Nil(t, err)

		// Verify deletion
		_, err = repo.Get(*createdID)
		assert.NotNil(t, err)
	})
}
