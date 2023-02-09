package adapters

import "dit_backend/src/infra"

type PasswordResetAdapter interface {
	AskPasswordResetMail(email string) infra.Error
	FindPasswordResetByToken(token string) infra.Error
	UpdatePasswordByPasswordReset(token, newPassword string) infra.Error
}
