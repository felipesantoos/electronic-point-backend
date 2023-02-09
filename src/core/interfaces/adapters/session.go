package adapters

import (
	"dit_backend/src/infra"

	"github.com/google/uuid"
)

type SessionAdapter interface {
	Store(uID uuid.UUID, accessToken string) infra.Error
	Exists(uID uuid.UUID, token string) (bool, infra.Error)
	RemoveSession(uID uuid.UUID) infra.Error
	GetSessionByAccountID(uID uuid.UUID) (string, infra.Error)
}
