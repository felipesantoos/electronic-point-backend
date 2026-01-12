package simplifiedAccount

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestSimplifiedAccount(t *testing.T) {
	id := uuid.New()
	name := "John Doe"
	birthDate := "1990-01-01"
	email := "john@example.com"
	cpf := "03611322055"

	sacc := New(&id, name, birthDate, email, cpf)

	assert.Equal(t, &id, sacc.ID())
	assert.Equal(t, name, sacc.Name())
	assert.Equal(t, birthDate, sacc.BirthDate())
	assert.Equal(t, email, sacc.Email())
	assert.Equal(t, cpf, sacc.CPF())

	newID := uuid.New()
	sacc.SetID(&newID)
	sacc.SetName("Jane Doe")
	sacc.SetBirthDate("1995-05-05")
	sacc.SetEmail("jane@example.com")
	sacc.SetCPF("12345678901")

	assert.Equal(t, &newID, sacc.ID())
	assert.Equal(t, "Jane Doe", sacc.Name())
	assert.Equal(t, "1995-05-05", sacc.BirthDate())
	assert.Equal(t, "jane@example.com", sacc.Email())
	assert.Equal(t, "12345678901", sacc.CPF())
}
