package postgres

import (
	"database/sql"
	"eletronic_point/src/core/domain/account"
	"eletronic_point/src/core/domain/credentials"
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/role"
	secondary "eletronic_point/src/core/interfaces/adapters"
	"eletronic_point/src/infra/repository"
	"eletronic_point/src/infra/repository/postgres/query"
	"encoding/hex"
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type authPostgresRepository struct{}

func NewAuthPostgresRepository() secondary.AuthPort {
	return &authPostgresRepository{}
}

func (r *authPostgresRepository) Login(credentials credentials.Credentials) (account.Account, errors.Error) {
	rows, err := repository.Queryx(query.Account().Select().ByCredentials(), credentials.Email())
	if err != nil {
		return nil, err
	}
	if !rows.Next() {
		return nil, errors.NewFromString("email and/or password are incorrect")
	}
	var id, password string
	scanErr := rows.Scan(&id, &password)
	if scanErr != nil {
		if scanErr == sql.ErrNoRows {
			return nil, errors.NewFromString("email and/or password are incorrect")
		}
		logger.Error().Msg(scanErr.Error())
		return nil, errors.NewUnexpected()
	}
	if err := comparePasswords(password, credentials.Password()); err != nil {
		return nil, err
	}
	account, err := r.getAccountByCustomQuery(query.Account().Select().ByID(), id)
	if err != nil {
		logger.Error().Msg(err.String())
		return nil, err
	}
	return account, nil
}

func (r *authPostgresRepository) getAccountByCustomQuery(query string, args ...interface{}) (account.Account, errors.Error) {
	rows, queryErr := repository.Queryx(query, args...)
	if queryErr != nil {
		return nil, queryErr
	}
	if !rows.Next() {
		return nil, errors.NewFromString("account not found")
	}
	var serializedAccount = map[string]interface{}{}
	scanErr := rows.MapScan(serializedAccount)
	if scanErr != nil {
		return nil, errors.NewUnexpected()
	}
	account, convErr := newAccountFromMapRows(serializedAccount)
	if convErr != nil {
		return nil, convErr
	}
	return account, nil
}

func (r *authPostgresRepository) ResetAccountPassword(accountID *uuid.UUID, newPassword string) errors.Error {
	encryptedPassword, encryptErr := encryptPassword(newPassword)
	if encryptErr != nil {
		return errors.New(encryptErr)
	}
	result, queryErr := repository.ExecQuery(query.Account().Update().Password(), encryptedPassword, accountID.String())
	if queryErr != nil {
		return queryErr
	} else if rowsAff, err := result.RowsAffected(); err != nil {
		return errors.NewInternal(err)
	} else if rowsAff == 0 {
		return errors.NewUnexpected()
	}
	return nil
}

func passwordIsValid(originalPassword, givenPassword string) bool {
	decodedPassword, _ := hex.DecodeString(originalPassword)
	err := bcrypt.CompareHashAndPassword(decodedPassword, []byte(givenPassword))
	return err == nil
}

func comparePasswords(current, confirmation string) errors.Error {
	if !passwordIsValid(current, confirmation) {
		return errors.NewFromString("invalid password")
	}
	return nil
}

func newRoleFromMapRows(data map[string]interface{}) (role.Role, errors.Error) {
	var err error
	var id uuid.UUID
	var name = fmt.Sprint(data["name"])
	var code = fmt.Sprint(data["code"])
	id, err = uuid.Parse(string(data["id"].([]uint8)))
	if err != nil {
		return nil, errors.NewUnexpected()
	}
	return role.New(&id, name, code)
}
