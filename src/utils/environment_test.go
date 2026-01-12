package utils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetenvWithDefault(t *testing.T) {
	key := "TEST_ENV_VAR"
	defaultValue := "default"
	
	// Test default
	os.Unsetenv(key)
	assert.Equal(t, defaultValue, GetenvWithDefault(key, defaultValue))

	// Test actual value
	os.Setenv(key, "actual")
	defer os.Unsetenv(key)
	assert.Equal(t, "actual", GetenvWithDefault(key, defaultValue))
}
