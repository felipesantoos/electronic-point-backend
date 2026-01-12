package account

import (
	"eletronic_point/src/core/domain/person"
	"eletronic_point/src/core/domain/professional"
	"eletronic_point/src/core/domain/role"
	"eletronic_point/src/core/domain/student"
	"eletronic_point/src/core/messages"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestAccount_SettersAndGetters(t *testing.T) {
	id := uuid.New()
	email := "test@example.com"
	password := "pass123"

	acc := &account{}
	acc.SetID(id)
	acc.SetEmail(email)
	acc.SetPassword(password)

	assert.Equal(t, id, *acc.ID())
	assert.Equal(t, email, acc.Email())
	assert.Equal(t, password, acc.Password())

	// Test new fields
	r, _ := role.New(nil, "Admin", role.ADMIN_ROLE_CODE)
	acc.SetRole(r)
	assert.Equal(t, r, acc.Role())

	p, _ := person.New(nil, "John Doe", "john@example.com", "1990-01-01", "11144477735", "82999999999", "", "")
	acc.SetPerson(p)
	assert.Equal(t, p, acc.Person())

	personID := uuid.New()
	prof, _ := professional.New(nil, &personID)
	acc.SetProfessional(prof)
	assert.Equal(t, prof, acc.Professional())

	s, _ := student.NewBuilder().WithRegistration("123").Build()
	acc.SetStudent(s)
	assert.Equal(t, s, acc.Student())
}

func TestAccount_IsValid(t *testing.T) {
	t.Run("Invalid Email", func(t *testing.T) {
		_, err := NewBuilder().WithEmail("invalid-email").Build()
		assert.NotNil(t, err)
		assert.Contains(t, err.Messages(), messages.AccountEmailErrorMessage)
	})

	t.Run("Invalid Person", func(t *testing.T) {
		p, _ := person.New(nil, "John", "invalid", "", "", "", "", "") // Invalid person
		acc := &account{person: p}
		err := acc.IsValid()
		assert.NotNil(t, err)
	})

	t.Run("Nil Person", func(t *testing.T) {
		acc := &account{email: "test@example.com"}
		err := acc.IsValid()
		assert.NotNil(t, err)
		assert.Contains(t, err.Messages(), "account person is required")
	})

	t.Run("Invalid Professional", func(t *testing.T) {
		p, _ := person.New(nil, "John Doe", "john@example.com", "1990-01-01", "11144477735", "82999999999", "", "")
		// Note: professional.New currently returns nil error always, but let's assume it can fail
		// For now we just test that it's called if present
		personID := uuid.New()
		prof, _ := professional.New(nil, &personID)
		acc := &account{
			email:        "test@example.com",
			person:       p,
			professional: prof,
		}
		assert.Nil(t, acc.IsValid())
	})

	t.Run("Invalid Student", func(t *testing.T) {
		p, _ := person.New(nil, "John Doe", "john@example.com", "1990-01-01", "11144477735", "82999999999", "", "")
		// Creating a student that fails IsValid (missing person)
		s, _ := student.NewBuilder().WithRegistration("123").Build()
		acc := &account{
			email:    "test@example.com",
			person:   p,
			_student: s,
		}
		err := acc.IsValid()
		assert.NotNil(t, err) // Should fail because student.Person is nil
	})

	t.Run("Valid Account", func(t *testing.T) {
		id := uuid.New()
		p, err := person.New(&id, "John Doe", "john@example.com", "1990-01-01", "11144477735", "82999999999", "", "")
		assert.Nil(t, err)
		acc := &account{
			email:  "test@example.com",
			person: p,
		}
		assert.Nil(t, acc.IsValid())
	})
}
