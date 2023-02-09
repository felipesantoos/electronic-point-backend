package redis

import (
	"dit_backend/src/core/domain/authorization"
	"dit_backend/src/core/interfaces/adapters"
	"dit_backend/src/infra"
	"fmt"

	"github.com/google/uuid"
)

type redisSessionRepository struct{}

func NewSessionRepository() adapters.SessionAdapter {
	return &redisSessionRepository{}
}

func (instance *redisSessionRepository) Store(uID uuid.UUID, accessToken string) infra.Error {
	conn, err := getConnection()
	if err != nil {
		return err
	}
	uSessionKey := instance.getUserSessionKey(uID)
	if err := conn.Set(uSessionKey, accessToken, authorization.TOKEN_TIMEOUT).Err(); err != nil {
		errLogger.Log().Msg(fmt.Sprintf("an error occurred when trying to save user session: %s", err.Error()))
		return infra.NewUnexpectedSourceErr()
	}
	return nil
}

func (instance *redisSessionRepository) Exists(uID uuid.UUID, token string) (bool, infra.Error) {
	conn, err := getConnection()
	if err != nil {
		return false, err
	}
	uSessionKey := instance.getUserSessionKey(uID)
	return valueExists(conn, uSessionKey, token)
}

func (instance *redisSessionRepository) GetSessionByAccountID(uID uuid.UUID) (string, infra.Error) {
	conn, err := getConnection()
	if err != nil {
		return "", err
	}
	uSessionKey := instance.getUserSessionKey(uID)
	accessToken, err := getValueFromKey(conn, uSessionKey)
	if err != nil {
		return "", err
	} else if accessToken == "" {
		return "", nil
	}
	return accessToken, nil
}

func (instance *redisSessionRepository) RemoveSession(uID uuid.UUID) infra.Error {
	conn, err := getConnection()
	if err != nil {
		return err
	}
	uSessionKey := instance.getUserSessionKey(uID)
	if value, err := getValueFromKey(conn, uSessionKey); err != nil {
		return err
	} else if value == "" {
		errLogger.Log().Msg(fmt.Sprintf("this user doesn't has a stored session"))
		return infra.NewSourceErrFromStr("this user doesn't has a stored session")
	}
	if result := conn.Del(uSessionKey); result.Err() != nil {
		errLogger.Log().Msg(fmt.Sprintf("an error occurred when removing user session: %s", result.Err().Error()))
		return infra.NewUnexpectedSourceErr()
	}
	return nil
}

func (*redisSessionRepository) getUserSessionKey(uID uuid.UUID) string {
	return fmt.Sprintf("user_session:%s", uID.String())
}
