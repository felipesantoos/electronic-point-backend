package redis

import (
	"backend_template/src/core/domain/errors"
	"backend_template/src/core/domain/simplifiedAccount"
	"backend_template/src/core/interfaces/adapters"
	"backend_template/src/core/utils"
	"backend_template/src/infra/mail"
	"backend_template/src/infra/repository"
	"backend_template/src/infra/repository/postgres/query"
	"fmt"
	"os"
	"time"

	"github.com/go-redis/redis"
	"github.com/thanhpk/randstr"
)

type redisPasswordResetRepository struct{}

func NewPasswordResetRepository() adapters.PasswordResetAdapter {
	return &redisPasswordResetRepository{}
}

func (instance *redisPasswordResetRepository) AskPasswordResetMail(email string) errors.Error {
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
	account, buildErr := simplifiedAccount.NewFromMap(serializedAccount)
	if buildErr != nil {
		return errors.NewInternal(buildErr)
	}
	if accountID, _ := instance.getPasswordResetEntry("*", account.ID().String()); accountID != "" {
		return errors.NewFromString("A token was already generated for reseting the password of this account")
	}
	token := randstr.Hex(16)
	redisConn, err := getConnection()
	if err != nil {
		return err
	}
	if _, err := redisConn.Set(getPasswordResetKey(token, account.ID().String()), account.ID().String(), time.Hour).Result(); err != nil {
		errLogger.Log().Msg(fmt.Sprintf("an error occurred when trying to add a new token entry for reseting password for email: %s", email))
		return errors.NewUnexpected()
	}
	passwordResetLink := buildPasswordResetLink(token)
	go func() {
		err := mail.SendPasswordResetEmail(account.Name(), account.Email(), passwordResetLink)
		if err != nil {
			errLogger.Log().Msg(fmt.Sprintf("Error when sending reset password email to %s: %v", account.Email(), err))
		}
	}()
	return nil
}

func (instance *redisPasswordResetRepository) FindPasswordResetByToken(token string) errors.Error {
	if _, err := instance.getPasswordResetEntry(token, "*"); err != nil {
		return err
	}
	return nil
}

func (instance *redisPasswordResetRepository) UpdatePasswordByPasswordReset(token, newPassword string) errors.Error {
	accountId, err := instance.getPasswordResetEntry(token, "*")
	if err != nil {
		return err
	}
	encryptedPassword, encryptErr := utils.EncryptPassword(newPassword)
	if encryptErr != nil {
		return errors.New(encryptErr)
	}
	result, queryErr := repository.ExecQuery(query.Account().Update().Password(), encryptedPassword, accountId)
	if queryErr != nil {
		return queryErr
	} else if rowsAff, err := result.RowsAffected(); err != nil {
		return errors.NewInternal(err)
	} else if rowsAff == 0 {
		return errors.NewUnexpected()
	} else if err := instance.deletePasswordResetEntryByAccountID(token); err != nil {
		return err
	}
	return nil
}

func (instance *redisPasswordResetRepository) getPasswordResetEntry(token, accountID string) (string, errors.Error) {
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

func (instance *redisPasswordResetRepository) deletePasswordResetEntryByAccountID(token string) errors.Error {
	conn, connErr := getConnection()
	if connErr != nil {
		return connErr
	}
	if _, err := conn.Del(getPasswordResetKey(token, "*")).Result(); err != nil {
		errLogger.Log().Msg(fmt.Sprintf("An unexpected error occurred when trying to delete the %s token entry: %v", token, err))
		return errors.NewUnexpected()
	}
	return nil
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
