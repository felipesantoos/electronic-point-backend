package simplifiedStudent

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
				assert.Contains(t, err.Messages(), messages.StudentIDErrorMessage)
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
		{"Valid Name", "John Doe", false},
		{"Single Name", "John", true},
		{"Empty Name", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBuilder().WithName(tt.input)
			_, err := b.Build()

			if tt.expectedError {
				assert.NotNil(t, err)
				assert.Contains(t, err.Messages(), messages.StudentNameErrorMessage)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestBuilder_Build(t *testing.T) {
	t.Run("Successful Build", func(t *testing.T) {
		id := uuid.New()
		name := "John Doe"
		b := NewBuilder().WithID(id).WithName(name).WithTotalWorkload(100)
		s, err := b.Build()

		assert.Nil(t, err)
		assert.NotNil(t, s)
		assert.Equal(t, id, *s.ID())
		assert.Equal(t, name, s.Name())
		assert.Equal(t, 100, s.TotalWorkload())
	})

	t.Run("Build with Multiple Errors", func(t *testing.T) {
		_, err := NewBuilder().
			WithID(uuid.Nil).
			WithName("").
			WithInstitution(nil).
			Build()

		assert.NotNil(t, err)
		assert.Contains(t, err.Messages(), messages.StudentIDErrorMessage)
		assert.Contains(t, err.Messages(), messages.StudentNameErrorMessage)
		assert.Contains(t, err.Messages(), messages.StudentInstitutionErrorMessage)
	})
}
