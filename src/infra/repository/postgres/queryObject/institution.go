package queryObject

import (
	"eletronic_point/src/core/domain/institution"
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/infra/repository/postgres/query"
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type InstitutionObjectBuilder interface {
	FromMap(data map[string]interface{}) (institution.Institution, errors.Error)
	FromRows(rows *sqlx.Rows) ([]institution.Institution, errors.Error)
}

type institutionQueryObjectBuilder struct{}

func Institution() InstitutionObjectBuilder {
	return &institutionQueryObjectBuilder{}
}

func (this *institutionQueryObjectBuilder) FromMap(data map[string]interface{}) (institution.Institution, errors.Error) {
	id, err := uuid.Parse(string(data[query.InstitutionID].([]uint8)))
	if err != nil {
		logger.Error().Msg(err.Error())
		return nil, errors.NewUnexpected()
	}
	name := fmt.Sprint(data[query.InstitutionName])
	_institution, validationError := institution.NewBuilder().WithID(id).WithName(name).Build()
	if validationError != nil {
		logger.Error().Msg(validationError.String())
		return nil, validationError
	}
	return _institution, nil
}

func (this *institutionQueryObjectBuilder) FromRows(rows *sqlx.Rows) ([]institution.Institution, errors.Error) {
	if rows == nil {
		err := errors.NewFromString("row value cannot be nil")
		logger.Error().Msg(err.String())
		return nil, err
	}
	defer rows.Close()
	institutions := make([]institution.Institution, 0)
	for rows.Next() {
		var serializedInstitution = map[string]interface{}{}
		nativeError := rows.MapScan(serializedInstitution)
		if nativeError != nil {
			logger.Error().Msg(nativeError.Error())
			return nil, errors.New(nativeError)
		}
		_institution, err := this.FromMap(serializedInstitution)
		if err != nil {
			logger.Error().Msg(err.String())
			return nil, err
		}
		institutions = append(institutions, _institution)
	}
	return institutions, nil
}
