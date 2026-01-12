package authorization

import (
	"eletronic_point/src/core"
	"eletronic_point/src/core/domain/account"
	"eletronic_point/src/core/domain/errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/rs/zerolog"
)

var logger zerolog.Logger = core.Logger()

const (
	TOKEN_TIMEOUT         = time.Hour
	REFRESH_TOKEN_TIMEOUT = 7 * 24 * time.Hour
	BEARER_TOKEN_TYPE     = "bearer"
	REFRESH_TOKEN_TYPE    = "refresh"
)

type Authorization interface {
	Token() string
	ExpirationTime() *time.Time
}

type authorization struct {
	token   string
	expTime *time.Time
}

func New() Authorization {
	return &authorization{}
}

func NewFromAccount(acc account.Account) (Authorization, errors.Error) {
	auth := &authorization{}
	auth.expTime = generateTokenExpirationTime()
	if err := auth.GenerateToken(acc, BEARER_TOKEN_TYPE); err != nil {
		return nil, err
	}
	return auth, nil
}

func NewRefreshToken(acc account.Account) (Authorization, errors.Error) {
	auth := &authorization{}
	auth.expTime = generateRefreshTokenExpirationTime()
	if err := auth.GenerateToken(acc, REFRESH_TOKEN_TYPE); err != nil {
		return nil, err
	}
	return auth, nil
}

func NewFromToken(accessToken string, expirationTime *time.Time) Authorization {
	return &authorization{accessToken, expirationTime}
}

func (auth *authorization) GenerateToken(account account.Account, tokenType string) errors.Error {
	secret := os.Getenv("SERVER_SECRET")
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, newClaims(
		account,
		tokenType,
		auth.expTime.Unix(),
	)).SignedString([]byte(secret))
	if err != nil {
		logger.Error().Msg(err.Error())
		return errors.NewUnexpected()
	}
	auth.token = token
	return nil
}

func (auth *authorization) Token() string {
	return auth.token
}

func (auth *authorization) ExpirationTime() *time.Time {
	return auth.expTime
}

func generateTokenExpirationTime() *time.Time {
	t := time.Now().Add(TOKEN_TIMEOUT)
	return &t
}

func generateRefreshTokenExpirationTime() *time.Time {
	t := time.Now().Add(REFRESH_TOKEN_TIMEOUT)
	return &t
}
