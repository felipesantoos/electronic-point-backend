package postgres

import (
	"dit_backend/src/core/domain/account"
	"dit_backend/src/core/domain/credentials"
	"dit_backend/src/core/interfaces/adapters"
	"dit_backend/src/core/utils"
	"dit_backend/src/infra"
	"dit_backend/src/infra/repository"
	"dit_backend/src/infra/repository/postgres/query"
)

type authPostgres struct{}

func NewAuthPostgres() adapters.AuthAdapter {
	return &authPostgres{}
}

func (instance *authPostgres) Login(credentials credentials.Credentials) (account.Account, infra.Error) {
	account, err := instance.getAccountByCustomQuery(query.Account().Select().ByCredentials(), credentials.Email())
	if err != nil {
		return nil, err
	}
	if err := comparePasswords(account.Password(), credentials.Password()); err != nil {
		return nil, err
	}
	return account, nil
}

func (instance *authPostgres) getAccountByCustomQuery(query string, args ...interface{}) (account.Account, infra.Error) {
	rows, queryErr := repository.Queryx(query, args...)
	if queryErr != nil {
		return nil, queryErr
	}
	if !rows.Next() {
		return nil, infra.NewSourceErrFromStr("account not found")
	}
	var serializedAccount = map[string]interface{}{}
	scanErr := rows.MapScan(serializedAccount)
	if scanErr != nil {
		return nil, infra.NewUnexpectedSourceErr()
	}
	account, convErr := account.NewFromMap(serializedAccount)
	if convErr != nil {
		return nil, infra.NewInternalSourceErr(convErr)
	}
	return account, nil
}

func comparePasswords(current, confirmation string) infra.Error {
	if !utils.PasswordIsValid(current, confirmation) {
		return infra.NewSourceErrFromStr("invalid password")
	}
	return nil
}
