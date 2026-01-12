package student

import (
	"eletronic_point/src/core/messages"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestBuilder_WithRegistration(t *testing.T) {
	tests := []struct {
		name          string
		registration  string
		expectedError bool
	}{
		{"Valid Registration", "2023101010", false},
		{"Empty Registration", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBuilder().WithRegistration(tt.registration)
			_, err := b.Build()

			if tt.expectedError {
				assert.NotNil(t, err)
				assert.Contains(t, err.Messages(), messages.StudentRegistrationErrorMessage)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestBuilder_WithWorkloads(t *testing.T) {
	t.Run("Valid Workloads", func(t *testing.T) {
		b := NewBuilder().
			WithTotalWorkload(100).
			WithWorkloadCompleted(50).
			WithPendingWorkload(50)
		_, err := b.Build()
		assert.Nil(t, err)
	})

	t.Run("Invalid Total Workload", func(t *testing.T) {
		b := NewBuilder().WithTotalWorkload(-1)
		_, err := b.Build()
		assert.NotNil(t, err)
		assert.Contains(t, err.Messages(), messages.StudentTotalWorkloadErrorMessage)
	})

	t.Run("Invalid Completed Workload", func(t *testing.T) {
		b := NewBuilder().WithWorkloadCompleted(-1)
		_, err := b.Build()
		assert.NotNil(t, err)
		assert.Contains(t, err.Messages(), messages.StudentWorkloadCompletedErrorMessage)
	})

	t.Run("Invalid Pending Workload", func(t *testing.T) {
		b := NewBuilder().WithPendingWorkload(-1)
		_, err := b.Build()
		assert.NotNil(t, err)
		assert.Contains(t, err.Messages(), messages.StudentPendingWorkloadErrorMessage)
	})
}

func TestBuilder_WithResponsibleTeacherID(t *testing.T) {
	tests := []struct {
		name          string
		teacherID     uuid.UUID
		expectedError bool
	}{
		{"Valid UUID", uuid.New(), false},
		{"Nil UUID", uuid.Nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBuilder().WithResponsibleTeacherID(tt.teacherID)
			_, err := b.Build()

			if tt.expectedError {
				assert.NotNil(t, err)
				assert.Contains(t, err.Messages(), messages.StudentResponsibleTeacherIDErrorMessage)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestBuilder_Build(t *testing.T) {
	t.Run("Build with Multiple Errors", func(t *testing.T) {
		_, err := NewBuilder().
			WithRegistration("").
			WithTotalWorkload(-1).
			WithResponsibleTeacherID(uuid.Nil).
			Build()

		assert.NotNil(t, err)
		assert.Contains(t, err.Messages(), messages.StudentRegistrationErrorMessage)
		assert.Contains(t, err.Messages(), messages.StudentTotalWorkloadErrorMessage)
		assert.Contains(t, err.Messages(), messages.StudentResponsibleTeacherIDErrorMessage)
	})
}
