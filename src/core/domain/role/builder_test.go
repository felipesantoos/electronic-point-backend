package role

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
				assert.Contains(t, err.Messages(), messages.RoleIDErrorMessage)
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
		{"Valid Name", "Administrator", false},
		{"Empty Name", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBuilder().WithName(tt.input)
			_, err := b.Build()

			if tt.expectedError {
				assert.NotNil(t, err)
				assert.Contains(t, err.Messages(), messages.RoleNameErrorMessage)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestBuilder_WithCode(t *testing.T) {
	tests := []struct {
		name          string
		code          string
		expectedError bool
	}{
		{"Valid Admin Code", ADMIN_ROLE_CODE, false},
		{"Valid Student Code", STUDENT_ROLE_CODE, false},
		{"Invalid Code", "INVALID", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBuilder().WithCode(tt.code)
			_, err := b.Build()

			if tt.expectedError {
				assert.NotNil(t, err)
				assert.Contains(t, err.Messages(), messages.RoleCodeErrorMessage)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestBuilder_Build(t *testing.T) {
	t.Run("Successful Build", func(t *testing.T) {
		id := uuid.New()
		name := "Administrator"
		code := ADMIN_ROLE_CODE

		r, err := NewBuilder().
			WithID(id).
			WithName(name).
			WithCode(code).
			Build()

		assert.Nil(t, err)
		assert.NotNil(t, r)
		assert.Equal(t, id, *r.ID())
		assert.Equal(t, name, r.Name())
		assert.Equal(t, code, r.Code())
	})

	t.Run("Build with Multiple Errors", func(t *testing.T) {
		_, err := NewBuilder().
			WithID(uuid.Nil).
			WithName("").
			WithCode("INVALID").
			Build()

		assert.NotNil(t, err)
		assert.Contains(t, err.Messages(), messages.RoleIDErrorMessage)
		assert.Contains(t, err.Messages(), messages.RoleNameErrorMessage)
		assert.Contains(t, err.Messages(), messages.RoleCodeErrorMessage)
	})
}
