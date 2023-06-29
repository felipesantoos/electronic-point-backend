package adapters

import (
	"backend_template/src/core/domain/errors"

	"github.com/google/uuid"
)

type SessionAdapter interface {
	Store(uID *uuid.UUID, accessToken string) errors.Error
	Exists(uID *uuid.UUID, token string) (bool, errors.Error)
	RemoveSession(uID *uuid.UUID) errors.Error
	GetSessionByAccountID(uID *uuid.UUID) (string, errors.Error)
}
