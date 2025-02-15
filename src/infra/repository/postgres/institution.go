package postgres

import (
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/institution"
	"eletronic_point/src/core/interfaces/secondary"
	"eletronic_point/src/core/services/filters"
	"eletronic_point/src/infra/repository"
	"eletronic_point/src/infra/repository/postgres/query"
	"eletronic_point/src/infra/repository/postgres/queryObject"
)

type institutionRepository struct{}

func NewInstitutionRepository() secondary.InstitutionPort {
	return &institutionRepository{}
}

func (this institutionRepository) List(_filters filters.InstitutionFilters) ([]institution.Institution, errors.Error) {
	rows, err := repository.Queryx(query.Institution().Select().All(), _filters.Name)
	if err != nil {
		logger.Error().Msg(err.String())
		return nil, errors.NewUnexpected()
	}
	defer rows.Close()
	institutions, err := queryObject.Institution().FromRows(rows)
	if err != nil {
		logger.Error().Msg(err.String())
		return nil, errors.NewUnexpected()
	}
	return institutions, nil
}
