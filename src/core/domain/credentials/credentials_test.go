package credentials

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCredentials(t *testing.T) {
	email := "test@example.com"
	password := "password123"

	c := New(email, password)

	assert.NotNil(t, c)
	assert.Equal(t, email, c.Email())
	assert.Equal(t, password, c.Password())
}
