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
}

type authorization struct {
	token string
}

func New() Authorization {
	return &authorization{}
}

func NewFromAccount(acc account.Account) (Authorization, errors.Error) {
	instance := &authorization{}
	if err := instance.GenerateToken(acc); err != nil {
		return nil, err
	}
	return instance, nil
}

func NewFromToken(accessToken string) Authorization {
	return &authorization{accessToken}
}

func (instance *authorization) Token() string {
	return instance.token
}

func (instance *authorization) GenerateToken(account account.Account) errors.Error {
	secret := os.Getenv("SERVER_SECRET")
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, newClaims(
		account,
		BEARER_TOKEN_TYPE,
		time.Now().Add(TOKEN_TIMEOUT).Unix(),
	)).SignedString([]byte(secret))
	if err != nil {
		logger.Error().Msg(err.Error())
		return errors.NewUnexpected()
	}
	instance.token = token
	return nil
}
