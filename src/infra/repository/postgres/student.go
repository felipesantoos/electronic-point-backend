package postgres

import (
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/student"
	"eletronic_point/src/core/interfaces/secondary"
	"eletronic_point/src/core/messages"
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
	var id uuid.UUID
	rows, err := repository.Queryx(query.Student().Insert(), _student.Name(), _student.Registration(),
		_student.ProfilePicture(), _student.Institution(), _student.Course(), _student.InternshipLocationName(),
		_student.InternshipAddress(), _student.InternshipLocation(), _student.TotalWorkload())
	if err != nil {
		logger.Error().Msg(err.String())
		if strings.Contains(err.String(), constraints.StudentRegistrationUK) {
			return nil, errors.NewConflictFromString(messages.StudentRegistrationIsAlreadyInUseErrorMessage)
		}
		return nil, errors.NewUnexpected()
	}
	if !rows.Next() {
		return nil, errors.NewUnexpected()
	}
	scanError := rows.Scan(&id)
	if scanError != nil {
		logger.Error().Msg(scanError.Error())
		return nil, errors.NewUnexpected()
	}
	return &id, nil
}

func (this studentRepository) Update(s student.Student) errors.Error {
	_, err := repository.ExecQuery(query.Student().Update(),
		s.ID,
		s.Name,
		s.Registration,
		s.ProfilePicture,
		s.Institution,
		s.Course,
		s.InternshipLocationName,
		s.InternshipAddress,
		s.InternshipLocation,
		s.TotalWorkload,
	)
	if err != nil {
		logger.Error().Msg(err.String())
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

func (this studentRepository) List() ([]student.Student, errors.Error) {
	rows, err := repository.Queryx(query.Student().Select().All())
	if err != nil {
		logger.Error().Msg(err.String())
		return nil, errors.NewUnexpected()
	}
	s, err := queryObject.Student().FromRows(rows)
	if err != nil {
		logger.Error().Msg(err.String())
		return nil, errors.NewUnexpected()
	}
	return s, nil
}

func (this studentRepository) Get(id uuid.UUID) (student.Student, errors.Error) {
	rows, err := repository.Queryx(query.Student().Select().ByID(), id)
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
	student, err := queryObject.Student().FromMap(serializedStudent)
	if err != nil {
		logger.Error().Msg(err.String())
		return nil, errors.NewUnexpected()
	}
	return student, nil
}
