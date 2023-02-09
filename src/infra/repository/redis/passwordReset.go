package redis

import (
	"dit_backend/src/core/domain/simplifiedAccount"
	"dit_backend/src/core/interfaces/adapters"
	"dit_backend/src/core/utils"
	"dit_backend/src/infra"
	"dit_backend/src/infra/mail"
	"dit_backend/src/infra/repository"
	"dit_backend/src/infra/repository/postgres/query"
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

func (instance *redisPasswordResetRepository) AskPasswordResetMail(email string) infra.Error {
	rows, queryErr := repository.Queryx(query.Account().Select().SimplifiedByEmail(), email)
	if queryErr != nil {
		return queryErr
	}
	if !rows.Next() {
		return infra.NewSourceErrFromStr("account not found")
	}
	var serializedAccount = map[string]interface{}{}
	scanErr := rows.MapScan(serializedAccount)
	if scanErr != nil {
		return infra.NewUnexpectedSourceErr()
	}
	account, buildErr := simplifiedAccount.NewFromMap(serializedAccount)
	if buildErr != nil {
		return infra.NewInternalSourceErr(buildErr)
	}
	if accountID, _ := instance.getPasswordResetEntry("*", account.ID().String()); accountID != "" {
		return infra.NewSourceErrFromStr("A token was already generated for reseting the password of this account")
	}
	token := randstr.Hex(16)
	redisConn, err := getConnection()
	if err != nil {
		return err
	}
	if _, err := redisConn.Set(getPasswordResetKey(token, account.ID().String()), account.ID().String(), time.Hour).Result(); err != nil {
		errLogger.Log().Msg(fmt.Sprintf("an error occurred when trying to add a new token entry for reseting password for email: %s", email))
		return infra.NewUnexpectedSourceErr()
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

func (instance *redisPasswordResetRepository) FindPasswordResetByToken(token string) infra.Error {
	if _, err := instance.getPasswordResetEntry(token, "*"); err != nil {
		return err
	}
	return nil
}

func (instance *redisPasswordResetRepository) UpdatePasswordByPasswordReset(token, newPassword string) infra.Error {
	accountId, err := instance.getPasswordResetEntry(token, "*")
	if err != nil {
		return err
	}
	encryptedPassword, encryptErr := utils.EncryptPassword(newPassword)
	if encryptErr != nil {
		return infra.NewSourceErr(encryptErr)
	}
	result, queryErr := repository.ExecQuery(query.Account().Update().Password(), encryptedPassword, accountId)
	if queryErr != nil {
		return queryErr
	} else if rowsAff, err := result.RowsAffected(); err != nil {
		return infra.NewInternalSourceErr(err)
	} else if rowsAff == 0 {
		return infra.NewUnexpectedSourceErr()
	} else if err := instance.deletePasswordResetEntryByAccountID(token); err != nil {
		return err
	}
	return nil
}

func (instance *redisPasswordResetRepository) getPasswordResetEntry(token, accountID string) (string, infra.Error) {
	conn, connErr := getConnection()
	if connErr != nil {
		return "", connErr
	}
	keys, err := conn.Keys(getPasswordResetKey(token, accountID)).Result()
	if err == redis.Nil || len(keys) == 0 {
		return "", infra.NewSourceErrFromStr("the provided token doesn't exists or expired.")
	} else if err != nil {
		return "", infra.NewUnexpectedSourceErr()
	}
	accountID, err = conn.Get(keys[0]).Result()
	if err != nil {
		return "", infra.NewUnexpectedSourceErr()
	}
	return accountID, nil
}

func (instance *redisPasswordResetRepository) deletePasswordResetEntryByAccountID(token string) infra.Error {
	conn, connErr := getConnection()
	if connErr != nil {
		return connErr
	}
	if _, err := conn.Del(getPasswordResetKey(token, "*")).Result(); err != nil {
		errLogger.Log().Msg(fmt.Sprintf("An unexpected error occurred when trying to delete the %s token entry: %v", token, err))
		return infra.NewUnexpectedSourceErr()
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
