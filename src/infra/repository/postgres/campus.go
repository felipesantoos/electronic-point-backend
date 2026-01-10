package postgres

import (
	"eletronic_point/src/core/domain/campus"
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/interfaces/secondary"
	"eletronic_point/src/core/services/filters"
	"eletronic_point/src/infra/repository"
	"eletronic_point/src/infra/repository/postgres/query"
	"eletronic_point/src/infra/repository/postgres/queryObject"

	"github.com/google/uuid"
)

type campusRepository struct{}

func NewCampusRepository() secondary.CampusPort {
	return &campusRepository{}
}

func (this campusRepository) List(_filters filters.CampusFilters) ([]campus.Campus, errors.Error) {
	rows, err := repository.Queryx(query.Campus().Select().All(), _filters.Name)
	if err != nil {
		logger.Error().Msg(err.String())
		return nil, errors.NewUnexpected()
	}
	defer rows.Close()
	campuss, err := queryObject.Campus().FromRows(rows)
	if err != nil {
		logger.Error().Msg(err.String())
		return nil, errors.NewUnexpected()
	}
	return campuss, nil
}

func (this campusRepository) Get(id uuid.UUID) (campus.Campus, errors.Error) {
	rows, err := repository.Queryx(query.Campus().Select().ByID(), id)
	if err != nil {
		logger.Error().Msg(err.String())
		return nil, errors.NewUnexpected()
	}
	defer rows.Close()
	campuses, err := queryObject.Campus().FromRows(rows)
	if err != nil {
		logger.Error().Msg(err.String())
		return nil, errors.NewUnexpected()
	}
	if len(campuses) == 0 {
		return nil, errors.NewFromString("campus not found")
	}
	return campuses[0], nil
}

func (this campusRepository) Create(data campus.Campus) (*uuid.UUID, errors.Error) {
	return execQueryReturningID(query.Campus().Insert(), data.Name(), data.InstitutionID())
}

func (this campusRepository) Update(data campus.Campus) errors.Error {
	return defaultExecQuery(query.Campus().Update(), data.ID(), data.Name(), data.InstitutionID())
}

func (this campusRepository) Delete(id uuid.UUID) errors.Error {
	return defaultExecQuery(query.Campus().Delete(), id)
}
