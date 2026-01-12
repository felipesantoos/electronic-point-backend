package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsEmailValid(t *testing.T) {
	assert.True(t, IsEmailValid("test@example.com"))
	assert.False(t, IsEmailValid("invalid-email"))
}
