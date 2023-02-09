package postgres

import (
	"database/sql"
	"dit_backend/src/core/domain/account"
	"dit_backend/src/core/domain/role"
	updatepassword "dit_backend/src/core/domain/updatePassword"
	"dit_backend/src/core/interfaces/adapters"
	"dit_backend/src/core/utils"
	"dit_backend/src/infra"
	mail "dit_backend/src/infra/mail"
	"dit_backend/src/infra/repository"
	"dit_backend/src/infra/repository/postgres/query"

	"github.com/google/uuid"
	"github.com/thanhpk/randstr"
)

type accountRepository struct {
}

func NewAccountRepository() adapters.AccountAdapter {
	return &accountRepository{}
}

func (instance *accountRepository) List() ([]account.Account, infra.Error) {
	rows, err := repository.Queryx(query.Account().Select().All())
	if err != nil {
		return nil, err
	}
	accounts := []account.Account{}
	for rows.Next() {
		var serializedAccount = map[string]interface{}{}
		rows.MapScan(serializedAccount)
		account, err := account.NewFromMap(serializedAccount)
		if err != nil {
			return nil, infra.NewInternalSourceErr(err)
		}
		accounts = append(accounts, account)
	}
	return accounts, nil
}

func (instance *accountRepository) FindByID(uID uuid.UUID) (account.Account, infra.Error) {
	rows, err := repository.Queryx(query.Account().Select().ByID(), uID.String())
	if err != nil {
		return nil, err
	}
	rows.Next()
	var serializedAccount = map[string]interface{}{}
	rows.MapScan(serializedAccount)
	account, buildErr := account.NewFromMap(serializedAccount)
	if buildErr != nil {
		return nil, infra.NewInternalSourceErr(buildErr)
	}
	return account, nil
}

func (instance *accountRepository) Create(account account.Account) (*uuid.UUID, infra.Error) {
	account.SetPassword(randstr.Hex(8))
	encryptedPassword, err := utils.EncryptPassword(account.Password())
	if err != nil {
		return nil, infra.NewUnexpectedSourceErr()
	}
	tx, txErr := repository.BeginTransaction()
	if txErr != nil {
		return nil, txErr
	}
	defer tx.CloseConn()
	personID, createPersonErr := txQueryRowReturningID(
		tx,
		query.Person().Insert(),
		account.Person().Name(),
		account.Person().BirthDate(),
		account.Person().Email(),
		account.Person().CPF(),
		account.Person().Phone(),
	)
	if createPersonErr != nil {
		return nil, createPersonErr
	} else if parseErr := account.Person().SetStringID(personID); parseErr != nil {
		return nil, infra.NewInternalSourceErr(parseErr)
	}
	if insertRoleDataErr := insertAccountRoleData(tx, account); insertRoleDataErr != nil {
		return nil, insertRoleDataErr
	}
	accountID, createAccErr := txQueryRowReturningID(
		tx,
		query.Account().Insert(),
		account.Email(),
		encryptedPassword,
		personID,
		account.Role().Code(),
	)
	if createAccErr != nil {
		return nil, createAccErr
	}
	commitErr := tx.Commit()
	if commitErr != nil {
		return nil, infra.NewUnexpectedSourceErr()
	}
	id, convErr := uuid.Parse(accountID)
	if convErr != nil {
		return nil, infra.NewUnexpectedSourceErr()
	}
	go mail.SendNewAccountEmail(account.Person().Email(), account.Password())
	return &id, nil
}

func (instance *accountRepository) UpdateAccountProfile(account account.Account) infra.Error {
	return defaultExecQuery(
		query.Account().Update().Profile(),
		account.Person().Name(),
		account.Person().BirthDate(),
		account.Person().Phone(),
		account.Person().ID().String(),
	)
}

func (instance *accountRepository) UpdateAccountPassword(accountID uuid.UUID, data updatepassword.UpdatePassword) infra.Error {
	rows, err := repository.Queryx(query.Account().Select().PasswordByID(), accountID.String())
	if err != nil {
		return err
	}
	var accountPassword string = ""
	rows.Next()
	rows.Scan(&accountPassword)
	if err := comparePasswords(accountPassword, data.CurrentPassword()); err != nil {
		return err
	}
	encryptedPassword, encryptErr := utils.EncryptPassword(data.NewPassword())
	if encryptErr != nil {
		return infra.NewSourceErr(encryptErr)
	}
	result, err := repository.ExecQuery(query.Account().Update().Password(), encryptedPassword, accountID.String())
	if err != nil {
		return err
	}
	if rowsAffected, resultErr := result.RowsAffected(); resultErr != nil {
		return infra.NewSourceErr(resultErr)
	} else if rowsAffected == 0 {
		return infra.NewUnexpectedSourceErr()
	}
	return nil
}

func insertAccountRoleData(tx *repository.SQLTransaction, account account.Account) infra.Error {
	var result sql.Result
	var err infra.Error
	if account.Role().Code() == role.ProfessionalRoleCode() {
		result, err = tx.ExecQuery(query.Professional().Insert(), account.Person().ID().String())
	}
	if err != nil {
		return err
	}
	if rowsAff, err := result.RowsAffected(); err != nil {
		return infra.NewInternalSourceErr(err)
	} else if rowsAff == 0 {
		return infra.NewUnexpectedSourceErr()
	}
	return nil
}
