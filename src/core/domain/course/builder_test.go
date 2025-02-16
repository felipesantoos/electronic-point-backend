package course

import (
	"eletronic_point/src/core/messages"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestBuilder_WithID(t *testing.T) {
	tests := []struct {
		name          string
		id            uuid.UUID
		expectedError bool
	}{
		{"Valid ID", uuid.New(), false},
		{"Nil ID", uuid.Nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBuilder().WithID(tt.id)
			_, err := b.Build()

			if tt.expectedError {
				assert.NotNil(t, err)
				assert.Contains(t, err.Messages(), messages.CourseIDErrorMessage)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestBuilder_WithName(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		expectedError bool
	}{
		{"Valid Name", "Course Name", false},
		{"Empty Name", "  ", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBuilder().WithName(tt.input)
			_, err := b.Build()

			if tt.expectedError {
				assert.NotNil(t, err)
				assert.Contains(t, err.Messages(), messages.CourseNameErrorMessage)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestBuilder_Build(t *testing.T) {
	t.Run("Successful Build", func(t *testing.T) {
		id := uuid.New()
		courseName := "Course Name"
		course, err := NewBuilder().
			WithID(id).
			WithName(courseName).
			Build()

		assert.Nil(t, err)
		assert.NotNil(t, course)
		assert.Equal(t, id, course.id)
		assert.Equal(t, courseName, course.name)
	})

	t.Run("Build with Errors", func(t *testing.T) {
		_, err := NewBuilder().
			WithID(uuid.Nil).
			WithName("  ").
			Build()

		assert.NotNil(t, err)
		assert.Contains(t, err.Messages(), messages.CourseIDErrorMessage)
		assert.Contains(t, err.Messages(), messages.CourseNameErrorMessage)
	})
}
