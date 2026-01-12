package person

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
				assert.Contains(t, err.Messages(), messages.PersonIDErrorMessage)
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
				assert.Contains(t, err.Messages(), messages.PersonNameErrorMessage)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestBuilder_WithEmail(t *testing.T) {
	tests := []struct {
		name          string
		email         string
		expectedError bool
	}{
		{"Valid Email", "john.doe@example.com", false},
		{"Invalid Email", "john.doe", true},
		{"Empty Email", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBuilder().WithEmail(tt.email)
			_, err := b.Build()

			if tt.expectedError {
				assert.NotNil(t, err)
				assert.Contains(t, err.Messages(), messages.PersonEmailErrorMessage)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestBuilder_WithCPF(t *testing.T) {
	tests := []struct {
		name          string
		cpf           string
		expectedError bool
	}{
		{"Valid CPF", "11144477735", false}, // Validated CPF
		{"Invalid CPF", "12345678900", true},
		{"Short CPF", "123", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBuilder().WithCPF(tt.cpf)
			_, err := b.Build()

			if tt.expectedError {
				assert.NotNil(t, err)
				assert.Contains(t, err.Messages(), messages.PersonCPFErrorMessage)
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
		email := "john.doe@example.com"
		birthDate := "1990-01-01"
		cpf := "11144477735"
		phone := "82999999999"

		p, err := NewBuilder().
			WithID(id).
			WithName(name).
			WithEmail(email).
			WithBirthDate(birthDate).
			WithCPF(cpf).
			WithPhone(phone).
			WithCreatedAt("2023-01-01T00:00:00Z").
			WithUpdatedAt("2023-01-01T00:00:00Z").
			Build()

		assert.Nil(t, err)
		assert.NotNil(t, p)
		assert.Equal(t, id, *p.ID())
		assert.Equal(t, name, p.Name())
		assert.Equal(t, email, p.Email())
		assert.Equal(t, cpf, p.CPF())
	})

	t.Run("Build with Multiple Errors", func(t *testing.T) {
		_, err := NewBuilder().
			WithID(uuid.Nil).
			WithName("").
			WithEmail("invalid").
			Build()

		assert.NotNil(t, err)
		assert.Contains(t, err.Messages(), messages.PersonIDErrorMessage)
		assert.Contains(t, err.Messages(), messages.PersonNameErrorMessage)
		assert.Contains(t, err.Messages(), messages.PersonEmailErrorMessage)
	})
}
