package postgres

import (
	"eletronic_point/src/core/domain/campus"
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/interfaces/secondary"
	"eletronic_point/src/core/services/filters"
	"eletronic_point/src/infra/repository"
	"eletronic_point/src/infra/repository/postgres/query"
	"eletronic_point/src/infra/repository/postgres/queryObject"
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
