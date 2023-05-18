package passwordReset

import "github.com/google/uuid"

type PasswordReset interface {
	AccountID() uuid.UUID
	Token() string
	CreatedAt() string
}

type passwordReset struct {
	accountID uuid.UUID
	token     string
	createdAt string
}

func New() PasswordReset {
	return &passwordReset{}
}

func (pwr *passwordReset) AccountID() uuid.UUID {
	return pwr.accountID
}

func (pwr *passwordReset) Token() string {
	return pwr.token
}

func (pwr *passwordReset) CreatedAt() string {
	return pwr.createdAt
}
