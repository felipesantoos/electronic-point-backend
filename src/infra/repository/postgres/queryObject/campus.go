package queryObject

import (
	"eletronic_point/src/core/domain/campus"
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/infra/repository/postgres/query"
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type CampusObjectBuilder interface {
	FromMap(data map[string]interface{}) (campus.Campus, errors.Error)
	FromRows(rows *sqlx.Rows) ([]campus.Campus, errors.Error)
}

type campusQueryObjectBuilder struct{}

func Campus() CampusObjectBuilder {
	return &campusQueryObjectBuilder{}
}

func (this *campusQueryObjectBuilder) FromMap(data map[string]interface{}) (campus.Campus, errors.Error) {
	id, err := uuid.Parse(string(data[query.CampusID].([]uint8)))
	if err != nil {
		logger.Error().Msg(err.Error())
		return nil, errors.NewUnexpected()
	}
	name := fmt.Sprint(data[query.CampusName])
	institutionID, err := uuid.Parse(string(data[query.CampusInstitutionID].([]uint8)))
	if err != nil {
		logger.Error().Msg(err.Error())
		return nil, errors.NewUnexpected()
	}
	_campus, validationError := campus.NewBuilder().WithID(id).WithName(name).WithInstitutionID(institutionID).Build()
	if validationError != nil {
		logger.Error().Msg(validationError.String())
		return nil, validationError
	}
	return _campus, nil
}

func (this *campusQueryObjectBuilder) FromRows(rows *sqlx.Rows) ([]campus.Campus, errors.Error) {
	if rows == nil {
		err := errors.NewFromString("row value cannot be nil")
		logger.Error().Msg(err.String())
		return nil, err
	}
	defer rows.Close()
	campuss := make([]campus.Campus, 0)
	for rows.Next() {
		var serializedCampus = map[string]interface{}{}
		nativeError := rows.MapScan(serializedCampus)
		if nativeError != nil {
			logger.Error().Msg(nativeError.Error())
			return nil, errors.New(nativeError)
		}
		_campus, err := this.FromMap(serializedCampus)
		if err != nil {
			logger.Error().Msg(err.String())
			return nil, err
		}
		campuss = append(campuss, _campus)
	}
	return campuss, nil
}
