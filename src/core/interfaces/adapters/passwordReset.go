package adapters

import "backend_template/src/core/domain/errors"

type PasswordResetAdapter interface {
	AskPasswordResetMail(email string) errors.Error
	FindPasswordResetByToken(token string) errors.Error
	UpdatePasswordByPasswordReset(token, newPassword string) errors.Error
}
