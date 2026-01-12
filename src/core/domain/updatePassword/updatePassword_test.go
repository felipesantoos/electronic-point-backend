package updatepassword

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	current := "old-pass"
	newPass := "new-pass"

	upwd := New(current, newPass)

	assert.NotNil(t, upwd)
	assert.Equal(t, current, upwd.CurrentPassword())
	assert.Equal(t, newPass, upwd.NewPassword())
}
