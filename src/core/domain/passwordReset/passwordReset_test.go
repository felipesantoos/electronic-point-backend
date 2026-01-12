package passwordReset

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewPasswordReset(t *testing.T) {
	pwr := New()
	assert.NotNil(t, pwr)
}

func TestPasswordReset_Fields(t *testing.T) {
	accountID := uuid.New()
	token := "reset-token"
	createdAt := "2023-01-01T00:00:00Z"

	pwr := &passwordReset{
		accountID: accountID,
		token:     token,
		createdAt: createdAt,
	}

	assert.Equal(t, accountID, pwr.AccountID())
	assert.Equal(t, token, pwr.Token())
	assert.Equal(t, createdAt, pwr.CreatedAt())
}
