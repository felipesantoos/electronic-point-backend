package campus

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
				assert.Contains(t, err.Messages(), messages.CampusIDErrorMessage)
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
		{"Valid Name", "Campus Name", false},
		{"Empty Name", "  ", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBuilder().WithName(tt.input)
			_, err := b.Build()

			if tt.expectedError {
				assert.NotNil(t, err)
				assert.Contains(t, err.Messages(), messages.CampusNameErrorMessage)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestBuilder_WithInstitutionID(t *testing.T) {
	tests := []struct {
		name          string
		institutionID uuid.UUID
		expectedError bool
	}{
		{"Valid Institution ID", uuid.New(), false},
		{"Nil Institution ID", uuid.Nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBuilder().WithInstitutionID(tt.institutionID)
			_, err := b.Build()

			if tt.expectedError {
				assert.NotNil(t, err)
				assert.Contains(t, err.Messages(), messages.CampusInstitutionIDErrorMessage)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestBuilder_Build(t *testing.T) {
	t.Run("Successful Build", func(t *testing.T) {
		id := uuid.New()
		institutionID := uuid.New()
		campus, err := NewBuilder().
			WithID(id).
			WithName("Campus Name").
			WithInstitutionID(institutionID).
			Build()

		assert.Nil(t, err)
		assert.NotNil(t, campus)
		assert.Equal(t, id, campus.id)
		assert.Equal(t, "Campus Name", campus.name)
		assert.Equal(t, institutionID, campus.institutionID)
	})

	t.Run("Build with Errors", func(t *testing.T) {
		_, err := NewBuilder().
			WithID(uuid.Nil).
			WithName("  ").
			WithInstitutionID(uuid.Nil).
			Build()

		assert.NotNil(t, err)
		assert.Contains(t, err.Messages(), messages.CampusIDErrorMessage)
		assert.Contains(t, err.Messages(), messages.CampusNameErrorMessage)
		assert.Contains(t, err.Messages(), messages.CampusInstitutionIDErrorMessage)
	})
}
