package authorization

import (
	"backend_template/src/core"
	"backend_template/src/core/domain/account"
	"backend_template/src/core/domain/errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/rs/zerolog"
)

var logger zerolog.Logger = core.Logger()

const (
	TOKEN_TIMEOUT     = time.Hour
	BEARER_TOKEN_TYPE = "bearer"
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
	if err := auth.GenerateToken(acc); err != nil {
		return nil, err
	}
	return auth, nil
}

func NewFromToken(accessToken string, expirationTime *time.Time) Authorization {
	return &authorization{accessToken, expirationTime}
}

func (auth *authorization) GenerateToken(account account.Account) errors.Error {
	secret := os.Getenv("SERVER_SECRET")
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, newClaims(
		account,
		BEARER_TOKEN_TYPE,
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
	time := time.Now().Add(TOKEN_TIMEOUT)
	return &time
}
