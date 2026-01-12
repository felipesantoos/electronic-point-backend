package person

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestPerson_Setters(t *testing.T) {
	p := &person{}
	
	err := p.SetName("John Doe")
	assert.Nil(t, err)
	assert.Equal(t, "John Doe", p.Name())

	err = p.SetEmail("john@example.com")
	assert.Nil(t, err)
	assert.Equal(t, "john@example.com", p.Email())

	err = p.SetBirthDate("1990-01-01")
	assert.Nil(t, err)
	assert.Equal(t, "1990-01-01", p.BirthDate())

	err = p.SetCPF("11144477735")
	assert.Nil(t, err)
	assert.Equal(t, "11144477735", p.CPF())
}

func TestPerson_Setters_Errors(t *testing.T) {
	p := &person{}
	
	err := p.SetName("John") // Single word
	assert.NotNil(t, err)

	err = p.SetEmail("invalid")
	assert.NotNil(t, err)

	err = p.SetBirthDate("01/01/1990") // Wrong format
	assert.NotNil(t, err)

	err = p.SetCPF("123") // Invalid
	assert.NotNil(t, err)
}

func TestPerson_IsValid(t *testing.T) {
	t.Run("Valid Person", func(t *testing.T) {
		id := uuid.New()
		p := &person{
			id:        &id,
			name:      "John Doe",
			email:     "john@example.com",
			birthDate: "1990-01-01",
			cpf:       "11144477735",
			phone:     "82999999999",
		}
		assert.Nil(t, p.IsValid())
	})

	t.Run("Invalid Name - Single Word", func(t *testing.T) {
		p := &person{name: "John"}
		err := p.IsValid()
		assert.NotNil(t, err)
		assert.Contains(t, err.Messages(), "you need to provide a name with two words or more.")
	})

	t.Run("Invalid Name - Empty", func(t *testing.T) {
		p := &person{name: ""}
		err := p.IsValid()
		assert.NotNil(t, err)
	})

	t.Run("Invalid CPF - Short Length", func(t *testing.T) {
		p := &person{
			name:      "John Doe",
			email:     "john@example.com",
			birthDate: "1990-01-01",
			cpf:       "1234567890", // 10 characters, should be 11
			phone:     "82999999999",
		}
		err := p.IsValid()
		assert.NotNil(t, err)
		assert.Contains(t, err.Messages(), "the CPF number must have 11 characters")
	})

	t.Run("Invalid CPF - Invalid Format", func(t *testing.T) {
		p := &person{
			name:      "John Doe",
			email:     "john@example.com",
			birthDate: "1990-01-01",
			cpf:       "12345678900", // 11 characters but invalid CPF
			phone:     "82999999999",
		}
		err := p.IsValid()
		assert.NotNil(t, err)
		assert.Contains(t, err.Messages(), "you need to provide a valid CPF")
	})

	t.Run("Invalid Email", func(t *testing.T) {
		p := &person{
			name:      "John Doe",
			email:     "invalid-email",
			birthDate: "1990-01-01",
			cpf:       "11144477735",
			phone:     "82999999999",
		}
		err := p.IsValid()
		assert.NotNil(t, err)
		assert.Contains(t, err.Messages(), "you must provide a valid email!")
	})

	t.Run("Invalid BirthDate - Wrong Format", func(t *testing.T) {
		p := &person{
			name:      "John Doe",
			email:     "john@example.com",
			birthDate: "01/01/1990", // Wrong format, should be yyyy-MM-dd
			cpf:       "11144477735",
			phone:     "82999999999",
		}
		err := p.IsValid()
		assert.NotNil(t, err)
		assert.Contains(t, err.Messages(), "you must provide a date according to the following syntax: yyyy-MM-dd")
	})

	t.Run("Invalid BirthDate - Empty", func(t *testing.T) {
		p := &person{
			name:      "John Doe",
			email:     "john@example.com",
			birthDate: "",
			cpf:       "11144477735",
			phone:     "82999999999",
		}
		err := p.IsValid()
		assert.NotNil(t, err)
	})

	t.Run("Multiple Validation Errors", func(t *testing.T) {
		p := &person{
			name:      "John", // Invalid
			email:     "invalid", // Invalid
			birthDate: "wrong-format", // Invalid
			cpf:       "123", // Invalid
		}
		err := p.IsValid()
		assert.NotNil(t, err)
		// Should have multiple error messages
		assert.Greater(t, len(err.Messages()), 1)
	})
}
