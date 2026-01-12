package postgres

import (
	"database/sql"
	"eletronic_point/src/core/domain"
	"eletronic_point/src/core/domain/account"
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/person"
	"eletronic_point/src/core/domain/professional"
	"eletronic_point/src/core/domain/role"
	updatepassword "eletronic_point/src/core/domain/updatePassword"
	"eletronic_point/src/core/interfaces/secondary"
	"eletronic_point/src/core/services/filters"
	mail "eletronic_point/src/infra/mail"
	"eletronic_point/src/infra/repository"
	"eletronic_point/src/infra/repository/postgres/query"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/thanhpk/randstr"
	"golang.org/x/crypto/bcrypt"
)

type accountRepository struct {
}

func NewAccountRepository() secondary.AccountPort {
	return &accountRepository{}
}

func (r *accountRepository) List(_filters filters.AccountFilters) ([]account.Account, errors.Error) {
	searchParam := ""
	if _filters.Search != nil {
		searchParam = *_filters.Search
	}
	rows, err := repository.Queryx(query.Account().Select().All(), _filters.RoleID, searchParam)
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

func (r *accountRepository) FindByID(uID *uuid.UUID) (account.Account, errors.Error) {
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
	if strings.EqualFold(account.Role().Code(), role.STUDENT_ROLE_CODE) {
		studentID := account.Person().ID()
		_student, err := NewStudentRepository().Get(*studentID, filters.StudentFilters{})
		if err != nil {
			logger.Error().Msg(err.String())
			return nil, err
		}
		account.SetStudent(_student)
	}
	return account, nil
}

func (r *accountRepository) Create(account account.Account) (*uuid.UUID, errors.Error) {
	transaction, err := repository.BeginTransaction()
	if err != nil {
		logger.Error().Msg(err.String())
		return nil, err
	}
	defer transaction.CloseConn()
	id, _, err := r.createPassingTrasaction(transaction, account)
	if err != nil {
		logger.Error().Msg(err.String())
		return nil, errors.NewUnexpected()
	}
	err = transaction.Commit()
	if err != nil {
		logger.Error().Msg(err.String())
		return nil, errors.NewUnexpected()
	}
	return id, nil
}

func (r *accountRepository) Update(account account.Account) errors.Error {
	transaction, err := repository.BeginTransaction()
	if err != nil {
		logger.Error().Msg(err.String())
		return err
	}
	defer transaction.CloseConn()
	err = r.updatePassingTrasaction(transaction, account)
	if err != nil {
		logger.Error().Msg(err.String())
		return err
	}
	err = transaction.Commit()
	if err != nil {
		logger.Error().Msg(err.String())
		return errors.NewUnexpected()
	}
	return nil
}

func (r *accountRepository) Delete(id uuid.UUID) errors.Error {
	_, err := repository.ExecQuery(query.Account().Delete(), id.String())
	if err != nil {
		return err
	}
	return nil
}

func (r *accountRepository) createPassingTrasaction(tx *repository.SQLTransaction, account account.Account) (*uuid.UUID, *uuid.UUID, errors.Error) {
	account.SetPassword(randstr.Hex(8))
	encryptedPassword, err := encryptPassword(account.Password())
	if err != nil {
		return nil, nil, errors.NewUnexpected()
	}
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
		return nil, nil, createPersonErr
	} else if parseErr := account.Person().SetStringID(personID); parseErr != nil {
		return nil, nil, errors.NewInternal(parseErr)
	}
	if insertRoleDataErr := insertAccountRoleData(tx, account); insertRoleDataErr != nil {
		return nil, nil, insertRoleDataErr
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
		return nil, nil, createAccErr
	}
	accountUUID, convErr := uuid.Parse(accountID)
	if convErr != nil {
		return nil, nil, errors.NewUnexpected()
	}
	personUUID, convErr := uuid.Parse(personID)
	if convErr != nil {
		return nil, nil, errors.NewUnexpected()
	}
	go mail.SendNewAccountEmail(account.Person().Email(), account.Password())
	return &accountUUID, &personUUID, nil
}

func (r *accountRepository) updatePassingTrasaction(tx *repository.SQLTransaction, account account.Account) errors.Error {
	err := defaultTxExecQuery(
		tx,
		query.Person().Update(),
		account.Person().ID(),
		account.Person().Name(),
		account.Person().BirthDate(),
		account.Person().Email(),
		account.Person().CPF(),
		account.Person().Phone(),
	)
	if err != nil {
		logger.Error().Msg(err.String())
		return err
	}
	err = defaultTxExecQuery(
		tx,
		query.Account().Update().EmailByPersonID(),
		account.Email(),
		account.Person().ID(),
	)
	if err != nil {
		logger.Error().Msg(err.String())
		return err
	}
	if account.ID() != nil {
		err = defaultTxExecQuery(
			tx,
			query.Account().Update().RoleByAccountID(),
			account.Role().Code(),
			account.ID(),
		)
	} else {
		err = defaultTxExecQuery(
			tx,
			query.Account().Update().RoleByPersonID(),
			account.Role().Code(),
			account.Person().ID(),
		)
	}
	if err != nil {
		logger.Error().Msg(err.String())
		return err
	}
	return nil
}

func (r *accountRepository) UpdateAccountProfile(person person.Person) errors.Error {
	return defaultExecQuery(
		query.Account().Update().Profile(),
		person.Name(),
		person.BirthDate(),
		person.Phone(),
		person.ID().String(),
	)
}

func (r *accountRepository) UpdateAccountPassword(accountID *uuid.UUID, data updatepassword.UpdatePassword) errors.Error {
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
	encryptedPassword, encryptErr := encryptPassword(data.NewPassword())
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
	if account.Role().IsAdmin() || account.Role().IsStudent() || account.Role().IsTeacher() {
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
	account, err := account.New(&id, email, "", role, person, nil, nil)
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
	return person.New(&id, name, email, birthDate, cpf, phone, createdAt, updatedAt)
}

func fillAccountWithProfessionalRoleEntry(r account.Account, data map[string]interface{}) errors.Error {
	var roleCode = fmt.Sprint(data["role_code"])
	if strings.ToLower(roleCode) == role.PROFESSIONAL_ROLE_CODE {
		if professionalData := domain.BuildMapWithoutParentName(data, role.PROFESSIONAL_ROLE_CODE); len(professionalData) == 0 {
			return errors.NewFromString("you must provide the professional r properties")
		} else {
			professional, err := newProfessionalFromMapRows(professionalData)
			professional.SetPersonID(r.Person().ID())
			if err != nil {
				return err
			}
			r.SetProfessional(professional)
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

func encryptPassword(password string) (string, error) {
	passwordBytes := []byte(password)
	hashedPassword, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(hashedPassword), nil
}
