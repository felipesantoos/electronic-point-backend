package redis

import (
	"backend_template/src/core/domain/authorization"
	"backend_template/src/core/domain/errors"
	"backend_template/src/core/interfaces/adapters"
	"fmt"

	"github.com/google/uuid"
)

type redisSessionRepository struct{}

func NewSessionRepository() adapters.SessionAdapter {
	return &redisSessionRepository{}
}

func (instance *redisSessionRepository) Store(uID uuid.UUID, accessToken string) errors.Error {
	conn, err := getConnection()
	if err != nil {
		return err
	}
	uSessionKey := instance.getUserSessionKey(uID)
	if err := conn.Set(uSessionKey, accessToken, authorization.TOKEN_TIMEOUT).Err(); err != nil {
		errLogger.Log().Msg(fmt.Sprintf("an error occurred when trying to save user session: %s", err.Error()))
		return errors.NewUnexpected()
	}
	return nil
}

func (instance *redisSessionRepository) Exists(uID uuid.UUID, token string) (bool, errors.Error) {
	conn, err := getConnection()
	if err != nil {
		return false, err
	}
	uSessionKey := instance.getUserSessionKey(uID)
	return valueExists(conn, uSessionKey, token)
}

func (instance *redisSessionRepository) GetSessionByAccountID(uID uuid.UUID) (string, errors.Error) {
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

func (instance *redisSessionRepository) RemoveSession(uID uuid.UUID) errors.Error {
	conn, err := getConnection()
	if err != nil {
		return err
	}
	uSessionKey := instance.getUserSessionKey(uID)
	if value, err := getValueFromKey(conn, uSessionKey); err != nil {
		return err
	} else if value == "" {
		errLogger.Log().Msg(fmt.Sprintf("this user doesn't has a stored session"))
		return errors.NewFromString("this user doesn't has a stored session")
	}
	if result := conn.Del(uSessionKey); result.Err() != nil {
		errLogger.Log().Msg(fmt.Sprintf("an error occurred when removing user session: %s", result.Err().Error()))
		return errors.NewUnexpected()
	}
	return nil
}

func (*redisSessionRepository) getUserSessionKey(uID uuid.UUID) string {
	return fmt.Sprintf("user_session:%s", uID.String())
}
