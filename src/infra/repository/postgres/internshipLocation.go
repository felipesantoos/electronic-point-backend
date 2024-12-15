package postgres

import (
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/internshipLocation"
	"eletronic_point/src/core/interfaces/secondary"
	"eletronic_point/src/core/messages"
	"eletronic_point/src/core/services/filters"
	"eletronic_point/src/infra/repository"
	"eletronic_point/src/infra/repository/postgres/query"
	"eletronic_point/src/infra/repository/postgres/queryObject"

	"github.com/google/uuid"
)

type internshipLocationRepository struct{}

func NewInternshipLocationRepository() secondary.InternshipLocationPort {
	return &internshipLocationRepository{}
}

func (this internshipLocationRepository) Create(_internshipLocation internshipLocation.InternshipLocation) (*uuid.UUID, errors.Error) {
	var id uuid.UUID
	rows, err := repository.Queryx(query.InternshipLocation().Insert(), _internshipLocation.Name(),
		_internshipLocation.Address(), _internshipLocation.City(), _internshipLocation.Lat(),
		_internshipLocation.Long())
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

func (this internshipLocationRepository) Update(_internshipLocation internshipLocation.InternshipLocation) errors.Error {
	_, err := repository.ExecQuery(query.InternshipLocation().Update(), _internshipLocation.ID(),
		_internshipLocation.Name(), _internshipLocation.Address(), _internshipLocation.City(),
		_internshipLocation.Lat(), _internshipLocation.Long())
	if err != nil {
		logger.Error().Msg(err.String())
		return errors.NewUnexpected()
	}
	return nil
}

func (this internshipLocationRepository) Delete(id uuid.UUID) errors.Error {
	_, err := repository.ExecQuery(query.InternshipLocation().Delete(), id)
	if err != nil {
		logger.Error().Msg(err.String())
		return errors.NewUnexpected()
	}
	return nil
}

func (this internshipLocationRepository) List(_filters filters.InternshipLocationFilters) ([]internshipLocation.InternshipLocation, errors.Error) {
	rows, err := repository.Queryx(query.InternshipLocation().Select().All(), _filters.StudentID)
	if err != nil {
		logger.Error().Msg(err.String())
		return nil, errors.NewUnexpected()
	}
	internshipLocations, err := queryObject.InternshipLocation().FromRows(rows)
	if err != nil {
		logger.Error().Msg(err.String())
		return nil, errors.NewUnexpected()
	}
	return internshipLocations, nil
}

func (this internshipLocationRepository) Get(id uuid.UUID) (internshipLocation.InternshipLocation, errors.Error) {
	rows, err := repository.Queryx(query.InternshipLocation().Select().ByID(), id)
	if err != nil {
		logger.Error().Msg(err.String())
		return nil, errors.NewUnexpected()
	}
	defer rows.Close()
	if !rows.Next() {
		return nil, errors.NewFromString(messages.InternshipLocationNotFoundErrorMessage)
	}
	var serializedInternshipLocation = map[string]interface{}{}
	nativeError := rows.MapScan(serializedInternshipLocation)
	if nativeError != nil {
		logger.Error().Msg(nativeError.Error())
		return nil, errors.NewUnexpected()
	}
	_internshipLocation, err := queryObject.InternshipLocation().FromMap(serializedInternshipLocation)
	if err != nil {
		logger.Error().Msg(err.String())
		return nil, errors.NewUnexpected()
	}
	return _internshipLocation, nil
}
