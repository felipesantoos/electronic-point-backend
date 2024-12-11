package postgres

import (
	"backend_template/src/core/domain/errors"
	"backend_template/src/core/domain/student"
	"backend_template/src/core/interfaces/adapters"
	"backend_template/src/core/messages"
	"backend_template/src/infra/repository"
	"backend_template/src/infra/repository/postgres/query"
	"backend_template/src/infra/repository/postgres/queryObject"

	"github.com/google/uuid"
)

type studentRepository struct{}

func NewStudentRepository() adapters.StudentAdapter {
	return &studentRepository{}
}

func (r studentRepository) Create(s student.Student) (*uuid.UUID, errors.Error) {
	var studentID uuid.UUID
	rows, err := repository.Queryx(query.Student().Insert(),
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
		return nil, errors.NewUnexpected()
	}
	rows.Scan(&studentID)
	if err != nil {
		logger.Error().Msg(err.String())
		return nil, errors.NewUnexpected()
	}
	return &studentID, nil
}

func (r studentRepository) Update(s student.Student) errors.Error {
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

func (r studentRepository) Delete(id uuid.UUID) errors.Error {
	_, err := repository.ExecQuery(query.Student().Delete(), id)
	if err != nil {
		logger.Error().Msg(err.String())
		return errors.NewUnexpected()
	}
	return nil
}

func (r studentRepository) List() ([]student.Student, errors.Error) {
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

func (r studentRepository) Get(id uuid.UUID) (student.Student, errors.Error) {
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
