package redis

import (
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/simplifiedAccount"
	"eletronic_point/src/core/interfaces/adapters"
	"eletronic_point/src/infra/mail"
	"eletronic_point/src/infra/repository"
	"eletronic_point/src/infra/repository/postgres/query"
	"fmt"
	"os"
	"time"

	"github.com/go-redis/redis"
	"github.com/google/uuid"
	"github.com/thanhpk/randstr"
)

type redisPasswordResetRepository struct{}

func NewPasswordResetRepository() adapters.PasswordResetAdapter {
	return &redisPasswordResetRepository{}
}

func (r *redisPasswordResetRepository) AskPasswordResetMail(email string) errors.Error {
	rows, queryErr := repository.Queryx(query.Account().Select().SimplifiedByEmail(), email)
	if queryErr != nil {
		return queryErr
	}
	if !rows.Next() {
		return errors.NewFromString("account not found")
	}
	var serializedAccount = map[string]interface{}{}
	scanErr := rows.MapScan(serializedAccount)
	if scanErr != nil {
		return errors.NewUnexpected()
	}
	account, buildErr := newSimplifiedAccountFromMapRows(serializedAccount)
	if buildErr != nil {
		return buildErr
	}
	if accountID, _ := r.getPasswordResetEntry("*", account.ID().String()); accountID != "" {
		return errors.NewFromString("A token was already generated for reseting the password of this account")
	}
	token := randstr.Hex(16)
	redisConn, err := getConnection()
	if err != nil {
		return err
	}
	if _, err := redisConn.Set(getPasswordResetKey(token, account.ID().String()), account.ID().String(), time.Hour).Result(); err != nil {
		logger.Log().Msg(fmt.Sprintf("an error occurred when trying to add a new token entry for reseting password for email: %s", email))
		return errors.NewUnexpected()
	}
	passwordResetLink := buildPasswordResetLink(token)
	go func() {
		err := mail.SendPasswordResetEmail(account.Name(), account.Email(), passwordResetLink)
		if err != nil {
			logger.Log().Msg(fmt.Sprintf("Error when sending reset password email to %s: %v", account.Email(), err))
		}
	}()
	return nil
}

func (r *redisPasswordResetRepository) FindPasswordResetByToken(token string) errors.Error {
	if _, err := r.getPasswordResetEntry(token, "*"); err != nil {
		return err
	}
	return nil
}

func (r *redisPasswordResetRepository) GetAccountIDByResetPasswordToken(token string) (*uuid.UUID, errors.Error) {
	accountID, err := r.getPasswordResetEntry(token, "*")
	if err != nil {
		return nil, err
	}
	if parsedUUID, err := uuid.Parse(accountID); err != nil {
		return nil, errors.NewInternal(err)
	} else {
		return &parsedUUID, nil
	}
}

func (r *redisPasswordResetRepository) DeleteResetPasswordEntry(token string) errors.Error {
	if err := r.deletePasswordResetEntryByAccountID(token); err != nil {
		return err
	}
	return nil
}

func (r *redisPasswordResetRepository) getPasswordResetEntry(token, accountID string) (string, errors.Error) {
	conn, connErr := getConnection()
	if connErr != nil {
		return "", connErr
	}
	keys, err := conn.Keys(getPasswordResetKey(token, accountID)).Result()
	if err == redis.Nil || len(keys) == 0 {
		return "", errors.NewFromString("the provided token doesn't exists or expired.")
	} else if err != nil {
		return "", errors.NewUnexpected()
	}
	accountID, err = conn.Get(keys[0]).Result()
	if err != nil {
		return "", errors.NewUnexpected()
	}
	return accountID, nil
}

func (r *redisPasswordResetRepository) deletePasswordResetEntryByAccountID(token string) errors.Error {
	conn, connErr := getConnection()
	if connErr != nil {
		return connErr
	}
	if _, err := conn.Del(getPasswordResetKey(token, "*")).Result(); err != nil {
		logger.Log().Msg(fmt.Sprintf("An unexpected error occurred when trying to delete the %s token entry: %v", token, err))
		return errors.NewUnexpected()
	}
	return nil
}

func newSimplifiedAccountFromMapRows(data map[string]interface{}) (simplifiedAccount.SimplifiedAccount, errors.Error) {
	var id uuid.UUID
	if parsedID, err := uuid.Parse(string(data["account_id"].([]uint8))); err != nil {
		return nil, errors.NewUnexpected()
	} else {
		id = parsedID
	}
	name := fmt.Sprint(data["person_name"])
	birthDate := fmt.Sprint(data["person_birth_date"])
	email := fmt.Sprint(data["account_email"])
	cpf := fmt.Sprint(data["person_cpf"])
	return simplifiedAccount.New(&id, name, birthDate, email, cpf), nil
}

func getPasswordResetKey(token, accountID string) string {
	return fmt.Sprintf("reset_password_token:%s:%s", token, accountID)
}

func buildPasswordResetLink(token string) string {
	host := os.Getenv("SERVER_UI_HOST")
	if host == "" {
		return ""
	}
	return fmt.Sprintf("%s/%s", host, token)
}
