package tokenextractor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRolesFromAuthHeader_Anonymous(t *testing.T) {
	roles, err := GetRolesFromAuthHeader("")
	assert.Nil(t, err)
	assert.Equal(t, []string{AnonymousRole}, roles)
}
