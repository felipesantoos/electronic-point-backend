package secondary

import (
	"eletronic_point/src/core/domain/errors"

	"github.com/google/uuid"
)

type PasswordResetPort interface {
	AskPasswordResetMail(email string) errors.Error
	FindPasswordResetByToken(token string) errors.Error
	GetAccountIDByResetPasswordToken(token string) (*uuid.UUID, errors.Error)
	DeleteResetPasswordEntry(token string) errors.Error
}
