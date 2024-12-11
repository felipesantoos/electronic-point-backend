package adapters

import (
	"eletronic_point/src/core/domain/errors"

	"github.com/google/uuid"
)

type PasswordResetAdapter interface {
	AskPasswordResetMail(email string) errors.Error
	FindPasswordResetByToken(token string) errors.Error
	GetAccountIDByResetPasswordToken(token string) (*uuid.UUID, errors.Error)
	DeleteResetPasswordEntry(token string) errors.Error
}
