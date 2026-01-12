package redis

import (
	"eletronic_point/src/apps/api/utils"
	"eletronic_point/src/core/domain/authorization"
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/interfaces/secondary"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type redisSessionRepository struct{}

func NewSessionRepository() secondary.SessionPort {
	return &redisSessionRepository{}
}

func (r *redisSessionRepository) Store(uID *uuid.UUID, accessToken string) errors.Error {
	conn, err := getConnection()
	if err != nil {
		return err
	}
	uSessionKey := r.getUserSessionKey(uID)
	tokenTimeout := authorization.TOKEN_TIMEOUT
	if !utils.IsAPIInProdMode() {
		tokenTimeout = time.Hour * 24
	}
	if err := conn.Set(uSessionKey, accessToken, tokenTimeout).Err(); err != nil {
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

func (r *redisSessionRepository) StoreRefreshToken(uID *uuid.UUID, refreshToken string) errors.Error {
	conn, err := getConnection()
	if err != nil {
		return err
	}
	uRefreshTokenKey := r.getUserRefreshTokenKey(uID)
	tokenTimeout := authorization.REFRESH_TOKEN_TIMEOUT
	if !utils.IsAPIInProdMode() {
		tokenTimeout = time.Hour * 24 * 7
	}
	if err := conn.Set(uRefreshTokenKey, refreshToken, tokenTimeout).Err(); err != nil {
		logger.Log().Msg(fmt.Sprintf("an error occurred when trying to save user refresh token: %s", err.Error()))
		return errors.NewUnexpected()
	}
	return nil
}

func (r *redisSessionRepository) ValidateRefreshToken(uID *uuid.UUID, refreshToken string) (bool, errors.Error) {
	conn, err := getConnection()
	if err != nil {
		return false, err
	}
	uRefreshTokenKey := r.getUserRefreshTokenKey(uID)
	return valueExists(conn, uRefreshTokenKey, refreshToken)
}

func (r *redisSessionRepository) RemoveRefreshToken(uID *uuid.UUID, refreshToken string) errors.Error {
	conn, err := getConnection()
	if err != nil {
		return err
	}
	uRefreshTokenKey := r.getUserRefreshTokenKey(uID)
	if exists, err := valueExists(conn, uRefreshTokenKey, refreshToken); err != nil {
		return err
	} else if !exists {
		return nil
	}
	if result := conn.Del(uRefreshTokenKey); result.Err() != nil {
		logger.Log().Msg(fmt.Sprintf("an error occurred when removing user refresh token: %s", result.Err().Error()))
		return errors.NewUnexpected()
	}
	return nil
}

func (r *redisSessionRepository) RemoveAllRefreshTokens(uID *uuid.UUID) errors.Error {
	conn, err := getConnection()
	if err != nil {
		return err
	}
	uRefreshTokenKey := r.getUserRefreshTokenKey(uID)
	if result := conn.Del(uRefreshTokenKey); result.Err() != nil {
		logger.Log().Msg(fmt.Sprintf("an error occurred when removing all user refresh tokens: %s", result.Err().Error()))
		return errors.NewUnexpected()
	}
	return nil
}

func (*redisSessionRepository) getUserSessionKey(uID *uuid.UUID) string {
	return fmt.Sprintf("user_session:%s", uID.String())
}

func (*redisSessionRepository) getUserRefreshTokenKey(uID *uuid.UUID) string {
	return fmt.Sprintf("refresh_token:%s", uID.String())
}
