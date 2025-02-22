package postgres

import (
	"eletronic_point/src/core/domain/account"
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/internship"
	"eletronic_point/src/core/domain/person"
	"eletronic_point/src/core/domain/role"
	"eletronic_point/src/core/domain/student"
	"eletronic_point/src/core/domain/timeRecord"
	"eletronic_point/src/core/interfaces/secondary"
	"eletronic_point/src/core/messages"
	"eletronic_point/src/core/services/filters"
	"eletronic_point/src/infra/repository"
	"eletronic_point/src/infra/repository/postgres/constraints"
	"eletronic_point/src/infra/repository/postgres/query"
	"eletronic_point/src/infra/repository/postgres/queryObject"
	"strings"

	"github.com/google/uuid"
)

type studentRepository struct{}

func NewStudentRepository() secondary.StudentPort {
	return &studentRepository{}
}

func (this studentRepository) Create(_student student.Student) (*uuid.UUID, errors.Error) {
	transaction, transactionError := repository.BeginTransaction()
	if transactionError != nil {
		logger.Error().Msg(transactionError.String())
		return nil, transactionError
	}
	defer transaction.CloseConn()
	_person, validationError := person.NewBuilder().
		WithName(_student.Name()).
		WithBirthDate(_student.BirthDate()).
		WithEmail(_student.Email()).
		WithCPF(_student.CPF()).
		WithPhone(_student.Phone()).
		Build()
	if validationError != nil {
		logger.Error().Msg(validationError.String())
		return nil, validationError
	}
	_role, validationError := role.NewBuilder().WithCode(role.STUDENT_ROLE_CODE).Build()
	if validationError != nil {
		logger.Error().Msg(validationError.String())
		return nil, validationError
	}
	_account, err := account.NewBuilder().
		WithEmail(_student.Email()).WithPerson(_person).WithRole(_role).Build()
	if err != nil {
		logger.Error().Msg(err.String())
		return nil, err
	}
	accountRepository := accountRepository{}
	_, personID, err := accountRepository.createPassingTrasaction(transaction, _account)
	if err != nil {
		logger.Error().Msg(err.String())
		return nil, err
	}
	err = defaultTxExecQuery(transaction, query.Student().Insert(), personID, _student.Registration(),
		_student.ProfilePicture(), _student.Campus().ID(), _student.Course().ID(), _student.TotalWorkload())
	if err != nil {
		logger.Error().Msg(err.String())
		if strings.Contains(err.String(), constraints.StudentRegistrationUK) {
			return nil, errors.NewConflictFromString(messages.StudentRegistrationIsAlreadyInUseErrorMessage)
		}
		return nil, errors.NewUnexpected()
	}
	err = defaultTxExecQuery(transaction, query.StudentLinkedToTeacher().Insert(), personID,
		_student.ResponsibleTeacherID())
	if err != nil {
		logger.Error().Msg(err.String())
		return nil, errors.NewUnexpected()
	}
	commitError := transaction.Commit()
	if commitError != nil {
		logger.Error().Msg(commitError.String())
		return nil, errors.NewUnexpected()
	}
	return personID, nil
}

func (this studentRepository) Update(_student student.Student) errors.Error {
	transaction, transactionError := repository.BeginTransaction()
	if transactionError != nil {
		logger.Error().Msg(transactionError.String())
		return transactionError
	}
	defer transaction.CloseConn()
	_person, validationError := person.NewBuilder().
		WithID(*_student.ID()).
		WithName(_student.Name()).
		WithBirthDate(_student.BirthDate()).
		WithEmail(_student.Email()).
		WithCPF(_student.CPF()).
		WithPhone(_student.Phone()).
		Build()
	if validationError != nil {
		logger.Error().Msg(validationError.String())
		return validationError
	}
	_role, validationError := role.NewBuilder().WithCode(role.STUDENT_ROLE_CODE).Build()
	if validationError != nil {
		logger.Error().Msg(validationError.String())
		return validationError
	}
	_account, err := account.NewBuilder().
		WithEmail(_student.Email()).WithPerson(_person).WithRole(_role).Build()
	if err != nil {
		logger.Error().Msg(err.String())
		return err
	}
	accountRepository := accountRepository{}
	err = accountRepository.updatePassingTrasaction(transaction, _account)
	if err != nil {
		logger.Error().Msg(err.String())
		return err
	}
	err = defaultTxExecQuery(transaction, query.Student().Update(), _student.ID(),
		_student.Registration(), _student.ProfilePicture(), _student.Campus().ID(),
		_student.Course().ID(), _student.TotalWorkload())
	if err != nil {
		logger.Error().Msg(err.String())
		if strings.Contains(err.String(), constraints.StudentRegistrationUK) {
			return errors.NewConflictFromString(messages.StudentRegistrationIsAlreadyInUseErrorMessage)
		}
		return errors.NewUnexpected()
	}
	commitError := transaction.Commit()
	if commitError != nil {
		logger.Error().Msg(commitError.String())
		return errors.NewUnexpected()
	}
	return nil
}

func (this studentRepository) Delete(id uuid.UUID) errors.Error {
	_, err := repository.ExecQuery(query.Student().Delete(), id)
	if err != nil {
		logger.Error().Msg(err.String())
		return errors.NewUnexpected()
	}
	return nil
}

func (this studentRepository) List(_filters filters.StudentFilters) ([]student.Student, errors.Error) {
	rows, err := repository.Queryx(query.Student().Select().All(), _filters.TeacherID, _filters.InstitutionID,
		_filters.CampusID, _filters.StudentID)
	if err != nil {
		logger.Error().Msg(err.String())
		return nil, errors.NewUnexpected()
	}
	students, err := queryObject.Student().FromRows(rows)
	if err != nil {
		logger.Error().Msg(err.String())
		return nil, errors.NewUnexpected()
	}
	for i := range students {
		timeRecordRepository := NewTimeRecordRepository()
		studentID := students[i].ID()
		_filters := filters.TimeRecordFilters{StudentID: studentID}
		timeRecords, err := timeRecordRepository.List(_filters)
		if err != nil {
			logger.Error().Msg(err.String())
			return nil, err
		}
		err = students[i].SetFrequencyHistory(timeRecords)
		if err != nil {
			logger.Error().Msg(err.String())
			return nil, err
		}
		workloadCompleted := calculateWorkloadCompleted(timeRecords)
		students[i].SetWorkloadCompleted(workloadCompleted)
		students[i].SetPendingWorkload(calculatePendingWorkload(students[i].TotalWorkload(), workloadCompleted))
		students[i].SetFrequencyHistory(nil)
	}
	return students, nil
}

func (this studentRepository) Get(id uuid.UUID, _filters filters.StudentFilters) (student.Student, errors.Error) {
	rows, err := repository.Queryx(query.Student().Select().ByID(), id, _filters.TeacherID)
	if err != nil {
		logger.Error().Msg(err.String())
		return nil, errors.NewUnexpected()
	}
	defer rows.Close()
	if !rows.Next() {
		return nil, errors.NewFromString(messages.StudentNotFoundErrorMessage)
	}
	var serializedStudent = map[string]interface{}{}
	nativeError := rows.MapScan(serializedStudent)
	if nativeError != nil {
		logger.Error().Msg(nativeError.Error())
		return nil, errors.NewUnexpected()
	}
	_student, err := queryObject.Student().FromMap(serializedStudent)
	if err != nil {
		logger.Error().Msg(err.String())
		return nil, errors.NewUnexpected()
	}
	timeRecordRepository := NewTimeRecordRepository()
	studentID := _student.ID()
	timeRecordFilters := filters.TimeRecordFilters{StudentID: studentID}
	timeRecords, err := timeRecordRepository.List(timeRecordFilters)
	if err != nil {
		logger.Error().Msg(err.String())
		return nil, err
	}
	err = _student.SetFrequencyHistory(timeRecords)
	if err != nil {
		logger.Error().Msg(err.String())
		return nil, err
	}
	workloadCompleted := calculateWorkloadCompleted(timeRecords)
	_student.SetWorkloadCompleted(workloadCompleted)
	_student.SetPendingWorkload(calculatePendingWorkload(_student.TotalWorkload(), workloadCompleted))
	internships, err := this.listInternshipsByStudent(*_student.ID())
	if err != nil {
		logger.Error().Msg(err.String())
		return nil, errors.NewUnexpected()
	}
	err = _student.SetInternshipHistory(internships)
	if err != nil {
		logger.Error().Msg(err.String())
		return nil, errors.NewUnexpected()
	}
	return _student, nil
}

func (this studentRepository) listInternshipsByStudent(studentID uuid.UUID) ([]internship.Internship, errors.Error) {
	rows, err := repository.Queryx(query.Internship().Select().ByStudentID(), studentID)
	if err != nil {
		logger.Error().Msg(err.String())
		return nil, errors.NewUnexpected()
	}
	internships, err := queryObject.Internship().FromRows(rows, false)
	if err != nil {
		logger.Error().Msg(err.String())
		return nil, errors.NewUnexpected()
	}
	return internships, nil
}

func calculateWorkloadCompleted(timeRecords []timeRecord.TimeRecord) int {
	sum := 0
	for _, _timeRecord := range timeRecords {
		if _timeRecord.ExitTime() != nil {
			duration := _timeRecord.ExitTime().Sub(_timeRecord.EntryTime())
			hours := duration.Hours()
			fullHours := int(hours)
			sum += fullHours
		}
	}
	return sum
}

func calculatePendingWorkload(totalWorkload, workloadCompleted int) int {
	return totalWorkload - workloadCompleted
}
