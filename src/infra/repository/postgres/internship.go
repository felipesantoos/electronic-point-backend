package postgres

import (
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/internship"
	"eletronic_point/src/core/interfaces/secondary"
	"eletronic_point/src/core/messages"
	"eletronic_point/src/core/services/filters"
	"eletronic_point/src/infra/repository"
	"eletronic_point/src/infra/repository/postgres/query"
	"eletronic_point/src/infra/repository/postgres/queryObject"

	"github.com/google/uuid"
)

type internshipRepository struct{}

func NewInternshipRepository() secondary.InternshipPort {
	return &internshipRepository{}
}

func (this internshipRepository) Create(_internship internship.Internship) (*uuid.UUID, errors.Error) {
	var id uuid.UUID
	rows, err := repository.Queryx(query.Internship().Insert(), _internship.Student().ID(),
		_internship.Location().ID(), _internship.StartedIn(), _internship.EndedIn())
	if err != nil {
		logger.Error().Msg(err.String())
		return nil, errors.NewUnexpected()
	}
	defer rows.Close()
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

func (this internshipRepository) Update(_internship internship.Internship) errors.Error {
	_, err := repository.ExecQuery(query.Internship().Update(), _internship.ID(),
		_internship.Student().ID(), _internship.Location().ID(), _internship.StartedIn(),
		_internship.EndedIn())
	if err != nil {
		logger.Error().Msg(err.String())
		return errors.NewUnexpected()
	}
	return nil
}

func (this internshipRepository) Delete(id uuid.UUID) errors.Error {
	_, err := repository.ExecQuery(query.Internship().Delete(), id)
	if err != nil {
		logger.Error().Msg(err.String())
		return errors.NewUnexpected()
	}
	return nil
}

func (this internshipRepository) List(_filters filters.InternshipFilters) ([]internship.Internship, errors.Error) {
	rows, err := repository.Queryx(query.Internship().Select().All(), _filters.StudentID)
	if err != nil {
		logger.Error().Msg(err.String())
		return nil, errors.NewUnexpected()
	}
	defer rows.Close()
	internships, err := queryObject.Internship().FromRows(rows, true)
	if err != nil {
		logger.Error().Msg(err.String())
		return nil, errors.NewUnexpected()
	}
	return internships, nil
}

func (this internshipRepository) Get(id uuid.UUID) (internship.Internship, errors.Error) {
	rows, err := repository.Queryx(query.Internship().Select().ByID(), id)
	if err != nil {
		logger.Error().Msg(err.String())
		return nil, errors.NewUnexpected()
	}
	defer rows.Close()
	if !rows.Next() {
		return nil, errors.NewFromString(messages.InternshipNotFoundErrorMessage)
	}
	var serializedInternship = map[string]interface{}{}
	nativeError := rows.MapScan(serializedInternship)
	if nativeError != nil {
		logger.Error().Msg(nativeError.Error())
		return nil, errors.NewUnexpected()
	}
	_internship, err := queryObject.Internship().FromMap(serializedInternship, true)
	if err != nil {
		logger.Error().Msg(err.String())
		return nil, errors.NewUnexpected()
	}
	return _internship, nil
}
