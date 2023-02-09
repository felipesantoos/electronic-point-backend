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

func (instance *passwordReset) AccountID() uuid.UUID {
	return instance.accountID
}

func (instance *passwordReset) Token() string {
	return instance.token
}

func (instance *passwordReset) CreatedAt() string {
	return instance.createdAt
}
