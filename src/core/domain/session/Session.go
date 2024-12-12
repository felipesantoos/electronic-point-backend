package session

import (
	"github.com/google/uuid"
)

type Session struct {
	accountID    uuid.UUID
	accessToken  string
	refreshToken string
}

func (s Session) AccountID() uuid.UUID {
	return s.accountID
}

func (s Session) AccessToken() string {
	return s.accessToken
}

func (s Session) RefreshToken() string {
	return s.refreshToken
}

func New(accountID uuid.UUID, accessToken, refreshToken string) *Session {
	return &Session{
		accountID:    accountID,
		accessToken:  accessToken,
		refreshToken: refreshToken,
	}
}

func NewReference(accountID uuid.UUID) *Session {
	return &Session{accountID: accountID}
}

func NewTokenReference(accountID uuid.UUID, token string) *Session {
	return &Session{
		accountID:   accountID,
		accessToken: token,
	}
}

func NewRefreshTokenReference(accountID uuid.UUID, refreshToken string) *Session {
	return &Session{
		accountID:    accountID,
		refreshToken: refreshToken,
	}
}
