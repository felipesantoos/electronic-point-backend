package postgres

import (
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/institution"
	"eletronic_point/src/core/interfaces/secondary"
	"eletronic_point/src/core/services/filters"
	"eletronic_point/src/infra/repository"
	"eletronic_point/src/infra/repository/postgres/query"
	"eletronic_point/src/infra/repository/postgres/queryObject"

	"github.com/google/uuid"
)

type institutionRepository struct{}

func NewInstitutionRepository() secondary.InstitutionPort {
	return &institutionRepository{}
}

func (this institutionRepository) List(_filters filters.InstitutionFilters) ([]institution.Institution, errors.Error) {
	rows, err := repository.Queryx(query.Institution().Select().All(), _filters.Name)
	if err != nil {
		logger.Error().Msg(err.String())
		return nil, err
	}
	defer rows.Close()
	institutions, err := queryObject.Institution().FromRows(rows)
	if err != nil {
		logger.Error().Msg(err.String())
		return nil, err
	}
	return institutions, nil
}

func (this institutionRepository) Get(id uuid.UUID) (institution.Institution, errors.Error) {
	rows, err := repository.Queryx(query.Institution().Select().ByID(), id)
	if err != nil {
		logger.Error().Msg(err.String())
		return nil, err
	}
	defer rows.Close()
	institutions, err := queryObject.Institution().FromRows(rows)
	if err != nil {
		logger.Error().Msg(err.String())
		return nil, err
	}
	if len(institutions) == 0 {
		return nil, errors.NewFromString("institution not found")
	}
	return institutions[0], nil
}

func (this institutionRepository) Create(data institution.Institution) (*uuid.UUID, errors.Error) {
	return execQueryReturningID(query.Institution().Insert(), data.Name())
}

func (this institutionRepository) Update(data institution.Institution) errors.Error {
	return defaultExecQuery(query.Institution().Update(), data.ID(), data.Name())
}

func (this institutionRepository) Delete(id uuid.UUID) errors.Error {
	return defaultExecQuery(query.Institution().Delete(), id)
}
