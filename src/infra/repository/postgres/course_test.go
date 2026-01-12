package postgres

import (
	"eletronic_point/src/core/domain/course"
	"eletronic_point/src/core/services/filters"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCourseRepository_Integration(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	SetupTestDB(t)
	CleanDB(t)

	repo := NewCourseRepository()

	t.Run("Create and Get", func(t *testing.T) {
		courseID := uuid.New()
		c, _ := course.NewBuilder().WithID(courseID).WithName("Test Course").Build()
		createdID, err := repo.Create(c)
		assert.Nil(t, err)
		assert.NotNil(t, createdID)

		// Get course
		found, err := repo.Get(*createdID)
		assert.Nil(t, err)
		assert.Equal(t, "Test Course", found.Name())
	})

	t.Run("List", func(t *testing.T) {
		// Create course
		courseID := uuid.New()
		c, _ := course.NewBuilder().WithID(courseID).WithName("Test Course 2").Build()
		_, err := repo.Create(c)
		assert.Nil(t, err)

		// List courses
		courses, err := repo.List(filters.CourseFilters{})
		assert.Nil(t, err)
		assert.NotEmpty(t, courses)
	})

	t.Run("Update", func(t *testing.T) {
		// Create course
		courseID := uuid.New()
		c, _ := course.NewBuilder().WithID(courseID).WithName("Test Course 3").Build()
		createdID, err := repo.Create(c)
		assert.Nil(t, err)

		// Update course
		updated, err := repo.Get(*createdID)
		assert.Nil(t, err)
		
		updatedCourse, _ := course.NewBuilder().
			WithID(updated.ID()).
			WithName("Updated Course").
			Build()

		err = repo.Update(updatedCourse)
		assert.Nil(t, err)

		// Verify update
		found, err := repo.Get(*createdID)
		assert.Nil(t, err)
		assert.Equal(t, "Updated Course", found.Name())
	})

	t.Run("Delete", func(t *testing.T) {
		// Create course
		courseID := uuid.New()
		c, _ := course.NewBuilder().WithID(courseID).WithName("Test Course 4").Build()
		createdID, err := repo.Create(c)
		assert.Nil(t, err)

		// Delete course
		err = repo.Delete(*createdID)
		assert.Nil(t, err)

		// Verify deletion
		_, err = repo.Get(*createdID)
		assert.NotNil(t, err)
	})
}
