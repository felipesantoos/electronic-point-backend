package redis

import (
	"backend_template/src/core/domain/authorization"
	"backend_template/src/core/domain/errors"
	"backend_template/src/core/interfaces/adapters"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type redisSessionRepository struct{}

func NewSessionRepository() adapters.SessionAdapter {
	return &redisSessionRepository{}
}

func (r *redisSessionRepository) Store(uID *uuid.UUID, accessToken string) errors.Error {
	conn, err := getConnection()
	if err != nil {
		return err
	}
	uSessionKey := r.getUserSessionKey(uID)
	if err := conn.Set(uSessionKey, accessToken, authorization.TOKEN_TIMEOUT).Err(); err != nil {
		logger.Log().Msg(fmt.Sprintf("an error occurred when trying to save user session: %s", err.Error()))
		return errors.NewUnexpected()
	}
	return nil
}

func (r *redisSessionRepository) Exists(uID *uuid.UUID, token string) (bool, errors.Error) {
	conn, err := getConnection()
	if err != nil {
		return false, err
	}
	uSessionKey := r.getUserSessionKey(uID)
	return valueExists(conn, uSessionKey, token)
}

func (r *redisSessionRepository) GetSessionByAccountID(uID *uuid.UUID) (authorization.Authorization, errors.Error) {
	conn, err := getConnection()
	if err != nil {
		return nil, err
	}
	uSessionKey := r.getUserSessionKey(uID)
	accessToken, err := getValueFromKey(conn, uSessionKey)
	if err != nil {
		return nil, err
	} else if accessToken == "" {
		return nil, nil
	}
	expDuration, err := getKeyDuration(conn, uSessionKey)
	expTime := time.Now().Add(*expDuration)
	return authorization.NewFromToken(accessToken, &expTime), nil
}

func (r *redisSessionRepository) RemoveSession(uID *uuid.UUID) errors.Error {
	conn, err := getConnection()
	if err != nil {
		return err
	}
	uSessionKey := r.getUserSessionKey(uID)
	if value, err := getValueFromKey(conn, uSessionKey); err != nil {
		return err
	} else if value == "" {
		logger.Log().Msg(fmt.Sprintf("this user doesn't has a stored session"))
		return errors.NewFromString("this user doesn't has a stored session")
	}
	if result := conn.Del(uSessionKey); result.Err() != nil {
		logger.Log().Msg(fmt.Sprintf("an error occurred when removing user session: %s", result.Err().Error()))
		return errors.NewUnexpected()
	}
	return nil
}

func (*redisSessionRepository) getUserSessionKey(uID *uuid.UUID) string {
	return fmt.Sprintf("user_session:%s", uID.String())
}
