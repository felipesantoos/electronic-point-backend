package postgres

import (
	"backend_template/src/core/domain"
	"backend_template/src/core/domain/account"
	"backend_template/src/core/domain/errors"
	"backend_template/src/core/domain/person"
	"backend_template/src/core/domain/professional"
	"backend_template/src/core/domain/role"
	updatepassword "backend_template/src/core/domain/updatePassword"
	"backend_template/src/core/interfaces/adapters"
	"backend_template/src/core/utils"
	mail "backend_template/src/infra/mail"
	"backend_template/src/infra/repository"
	"backend_template/src/infra/repository/postgres/query"
	"database/sql"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/thanhpk/randstr"
)

type accountRepository struct {
}

func NewAccountRepository() adapters.AccountAdapter {
	return &accountRepository{}
}

func (instance *accountRepository) List() ([]account.Account, errors.Error) {
	rows, err := repository.Queryx(query.Account().Select().All())
	if err != nil {
		return nil, err
	}
	accounts := []account.Account{}
	for rows.Next() {
		var serializedAccount = map[string]interface{}{}
		rows.MapScan(serializedAccount)
		account, err := newAccountFromMapRows(serializedAccount)
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}
	return accounts, nil
}

func (instance *accountRepository) FindByID(uID uuid.UUID) (account.Account, errors.Error) {
	rows, err := repository.Queryx(query.Account().Select().ByID(), uID.String())
	if err != nil {
		return nil, err
	}
	if !rows.Next() {
		return nil, errors.NewFromString("account not found")
	}
	var serializedAccount = map[string]interface{}{}
	rows.MapScan(serializedAccount)
	account, err := newAccountFromMapRows(serializedAccount)
	if err != nil {
		return nil, err
	}
	return account, nil
}

func (instance *accountRepository) Create(account account.Account) (*uuid.UUID, errors.Error) {
	account.SetPassword(randstr.Hex(8))
	encryptedPassword, err := utils.EncryptPassword(account.Password())
	if err != nil {
		return nil, errors.NewUnexpected()
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
		return nil, errors.NewInternal(parseErr)
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
		return nil, errors.NewUnexpected()
	}
	id, convErr := uuid.Parse(accountID)
	if convErr != nil {
		return nil, errors.NewUnexpected()
	}
	go mail.SendNewAccountEmail(account.Person().Email(), account.Password())
	return &id, nil
}

func (instance *accountRepository) UpdateAccountProfile(person person.Person) errors.Error {
	return defaultExecQuery(
		query.Account().Update().Profile(),
		person.Name(),
		person.BirthDate(),
		person.Phone(),
		person.ID().String(),
	)
}

func (instance *accountRepository) UpdateAccountPassword(accountID uuid.UUID, data updatepassword.UpdatePassword) errors.Error {
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
		return errors.New(encryptErr)
	}
	result, err := repository.ExecQuery(query.Account().Update().Password(), encryptedPassword, accountID.String())
	if err != nil {
		return err
	}
	if rowsAffected, resultErr := result.RowsAffected(); resultErr != nil {
		return errors.New(resultErr)
	} else if rowsAffected == 0 {
		return errors.NewUnexpected()
	}
	return nil
}

func insertAccountRoleData(tx *repository.SQLTransaction, account account.Account) errors.Error {
	if account.Role().IsAdmin() {
		return nil
	}
	var result sql.Result
	var err errors.Error
	if account.Role().IsProfessional() {
		result, err = tx.ExecQuery(query.Professional().Insert(), account.Person().ID().String())
	}
	if err != nil {
		return err
	}
	if rowsAff, err := result.RowsAffected(); err != nil {
		return errors.NewUnexpected()
	} else if rowsAff == 0 {
		return errors.NewUnexpected()
	}
	return nil
}

func newAccountFromMapRows(data map[string]interface{}) (account.Account, errors.Error) {
	var id uuid.UUID
	var email = fmt.Sprint(data["email"])
	if parsedID, err := uuid.Parse(string(data["id"].([]uint8))); err != nil {
		return nil, errors.NewUnexpected()
	} else {
		id = parsedID
	}
	var roleData = domain.BuildMapWithoutParentName(data, "role")
	role, err := newRoleFromMapRows(roleData)
	if err != nil {
		return nil, err
	}
	var personData = domain.BuildMapWithoutParentName(data, "person")
	personData["account_email"] = data["email"]
	person, err := newPersonFromMapRows(personData)
	if err != nil {
		return nil, err
	}
	account, err := account.New(&id, email, "", role, person, nil)
	if err != nil {
		return nil, err
	}
	err = fillAccountWithProfessionalRoleEntry(account, data)
	if err != nil {
		return nil, err
	}
	return account, nil
}

func newPersonFromMapRows(data map[string]interface{}) (person.Person, errors.Error) {
	var id uuid.UUID
	var name = fmt.Sprint(data["name"])
	var birthDate = domain.ParseUTCTimestampToDate(fmt.Sprint(data["birth_date"]))
	var cpf = fmt.Sprint(data["cpf"])
	var email = fmt.Sprint(data["account_email"])
	var phone = fmt.Sprint(data["phone"])
	var createdAt = domain.ParseUTCTimestampToRFCNano(fmt.Sprint(data["created_at"]))
	var updatedAt = domain.ParseUTCTimestampToRFCNano(fmt.Sprint(data["updated_at"]))
	if parsedID, err := uuid.Parse(string(data["id"].([]uint8))); err != nil {
		return nil, errors.NewUnexpected()
	} else {
		id = parsedID
	}
	return person.New(&id, name, birthDate, email, cpf, phone, createdAt, updatedAt)
}

func fillAccountWithProfessionalRoleEntry(instance account.Account, data map[string]interface{}) errors.Error {
	var roleCode = fmt.Sprint(data["role_code"])
	if strings.ToLower(roleCode) == role.PROFESSIONAL_ROLE_CODE {
		if professionalData := domain.BuildMapWithoutParentName(data, role.PROFESSIONAL_ROLE_CODE); len(professionalData) == 0 {
			return errors.NewFromString("you must provide the professional instance properties")
		} else {
			professional, err := newProfessionalFromMapRows(professionalData)
			professional.SetPersonID(instance.Person().ID())
			if err != nil {
				return err
			}
			instance.SetProfessional(professional)
		}
	}
	return nil
}

func newProfessionalFromMapRows(data map[string]interface{}) (professional.Professional, errors.Error) {
	var id uuid.UUID
	if parsedID, err := uuid.Parse(string(data["id"].([]uint8))); err != nil {
		return nil, errors.NewUnexpected()
	} else {
		id = parsedID
	}
	return professional.New(&id, nil)
}
