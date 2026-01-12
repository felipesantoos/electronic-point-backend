package internship

import (
	"eletronic_point/src/core/messages"
	"testing"
	"time"

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
				assert.Contains(t, err.Messages(), messages.InternshipIDErrorMessage)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestBuilder_WithStartedIn(t *testing.T) {
	tests := []struct {
		name          string
		startedIn     time.Time
		expectedError bool
	}{
		{"Valid Time", time.Now(), false},
		{"Zero Time", time.Time{}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBuilder().WithStartedIn(tt.startedIn)
			_, err := b.Build()

			if tt.expectedError {
				assert.NotNil(t, err)
				assert.Contains(t, err.Messages(), messages.InternshipStartedInErrorMessage)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestBuilder_Build(t *testing.T) {
	t.Run("Build with Multiple Errors", func(t *testing.T) {
		_, err := NewBuilder().
			WithID(uuid.Nil).
			WithStartedIn(time.Time{}).
			WithLocation(nil).
			WithStudent(nil).
			Build()

		assert.NotNil(t, err)
		assert.Contains(t, err.Messages(), messages.InternshipIDErrorMessage)
		assert.Contains(t, err.Messages(), messages.InternshipStartedInErrorMessage)
		assert.Contains(t, err.Messages(), messages.InternshipLocationErrorMessage)
		assert.Contains(t, err.Messages(), messages.InternshipStudentErrorMessage)
	})
}
