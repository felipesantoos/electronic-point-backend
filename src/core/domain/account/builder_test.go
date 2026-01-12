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
				assert.Contains(t, err.Messages(), messages.AccountIDErrorMessage)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestBuilder_WithRole(t *testing.T) {
	r, _ := role.New(nil, "Admin", role.ADMIN_ROLE_CODE)
	b := NewBuilder().WithRole(r)
	acc, err := b.Build()

	assert.Nil(t, err)
	assert.Equal(t, r, acc.Role())
}

func TestBuilder_WithPerson(t *testing.T) {
	p, _ := person.New(nil, "John Doe", "john@example.com", "1990-01-01", "11144477735", "82999999999", "", "")
	b := NewBuilder().WithPerson(p)
	acc, err := b.Build()

	assert.Nil(t, err)
	assert.Equal(t, p, acc.Person())
}

func TestBuilder_WithProfessional(t *testing.T) {
	personID := uuid.New()
	prof, _ := professional.New(nil, &personID)
	b := NewBuilder().WithProfessional(prof)
	acc, err := b.Build()

	assert.Nil(t, err)
	assert.Equal(t, prof, acc.Professional())
}

func TestBuilder_WithStudent(t *testing.T) {
	p, _ := person.New(nil, "John Doe", "john@example.com", "1990-01-01", "11144477735", "82999999999", "", "")
	s, _ := student.NewBuilder().WithPerson(p).WithRegistration("123").Build()
	b := NewBuilder().WithStudent(s)
	acc, err := b.Build()

	assert.Nil(t, err)
	assert.Equal(t, s, acc.Student())
}

func TestBuilder_WithEmail(t *testing.T) {
	tests := []struct {
		name          string
		email         string
		expectedError bool
	}{
		{"Valid Email", "test@example.com", false},
		{"Invalid Email", "invalid-email", true},
		{"Empty Email", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBuilder().WithEmail(tt.email)
			_, err := b.Build()

			if tt.expectedError {
				assert.NotNil(t, err)
				assert.Contains(t, err.Messages(), messages.AccountEmailErrorMessage)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestBuilder_WithPassword(t *testing.T) {
	tests := []struct {
		name          string
		password      string
		expectedError bool
	}{
		{"Valid Password", "password123", false},
		{"Short Password", "12345", true},
		{"Empty Password", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBuilder().WithPassword(tt.password)
			_, err := b.Build()

			if tt.expectedError {
				assert.NotNil(t, err)
				assert.Contains(t, err.Messages(), messages.AccountPasswordErrorMessage)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestBuilder_Build(t *testing.T) {
	t.Run("Successful Build", func(t *testing.T) {
		id := uuid.New()
		email := "test@example.com"
		password := "password123"

		account, err := NewBuilder().
			WithID(id).
			WithEmail(email).
			WithPassword(password).
			Build()

		assert.Nil(t, err)
		assert.NotNil(t, account)
		assert.Equal(t, id, *account.ID())
		assert.Equal(t, email, account.Email())
		assert.Equal(t, password, account.Password())
	})

	t.Run("Build with Multiple Errors", func(t *testing.T) {
		_, err := NewBuilder().
			WithID(uuid.Nil).
			WithEmail("invalid").
			WithPassword("123").
			Build()

		assert.NotNil(t, err)
		assert.Len(t, err.Messages(), 3)
		assert.Contains(t, err.Messages(), messages.AccountIDErrorMessage)
		assert.Contains(t, err.Messages(), messages.AccountEmailErrorMessage)
		assert.Contains(t, err.Messages(), messages.AccountPasswordErrorMessage)
	})
}
