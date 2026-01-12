package role

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRole_Helpers(t *testing.T) {
	r := &role{code: ADMIN_ROLE_CODE}
	assert.True(t, r.IsAdmin())
	assert.False(t, r.IsStudent())

	r.code = STUDENT_ROLE_CODE
	assert.True(t, r.IsStudent())
}

func TestRole_IsValid(t *testing.T) {
	r := &role{code: "INVALID"}
	assert.NotNil(t, r.IsValid())

	r.code = ADMIN_ROLE_CODE
	assert.Nil(t, r.IsValid())
}
