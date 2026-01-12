package session

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	accountID := uuid.New()
	accessToken := "access"
	refreshToken := "refresh"

	s := New(accountID, accessToken, refreshToken)

	assert.NotNil(t, s)
	assert.Equal(t, accountID, s.AccountID())
	assert.Equal(t, accessToken, s.AccessToken())
	assert.Equal(t, refreshToken, s.RefreshToken())
}

func TestNewReference(t *testing.T) {
	accountID := uuid.New()
	s := NewReference(accountID)
	assert.Equal(t, accountID, s.AccountID())
	assert.Empty(t, s.AccessToken())
}

func TestNewTokenReference(t *testing.T) {
	accountID := uuid.New()
	token := "token"
	s := NewTokenReference(accountID, token)
	assert.Equal(t, accountID, s.AccountID())
	assert.Equal(t, token, s.AccessToken())
}

func TestNewRefreshTokenReference(t *testing.T) {
	accountID := uuid.New()
	token := "refresh"
	s := NewRefreshTokenReference(accountID, token)
	assert.Equal(t, accountID, s.AccountID())
	assert.Equal(t, token, s.RefreshToken())
}
