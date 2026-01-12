package secondary

import (
	"eletronic_point/src/core/domain/authorization"
	"eletronic_point/src/core/domain/errors"

	"github.com/google/uuid"
)

type SessionPort interface {
	Store(uID *uuid.UUID, accessToken string) errors.Error
	Exists(uID *uuid.UUID, token string) (bool, errors.Error)
	RemoveSession(uID *uuid.UUID) errors.Error
	GetSessionByAccountID(uID *uuid.UUID) (authorization.Authorization, errors.Error)

	StoreRefreshToken(uID *uuid.UUID, refreshToken string) errors.Error
	ValidateRefreshToken(uID *uuid.UUID, refreshToken string) (bool, errors.Error)
	RemoveRefreshToken(uID *uuid.UUID, refreshToken string) errors.Error
	RemoveAllRefreshTokens(uID *uuid.UUID) errors.Error
}
