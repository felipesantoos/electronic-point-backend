package queryObject

import (
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/internshipLocation"
	"eletronic_point/src/infra/repository/postgres/query"
	"eletronic_point/src/utils"
	"fmt"
	"strconv"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type InternshipLocationObjectBuilder interface {
	FromMap(data map[string]interface{}) (internshipLocation.InternshipLocation, errors.Error)
	FromRows(rows *sqlx.Rows) ([]internshipLocation.InternshipLocation, errors.Error)
}

type internshipLocationQueryObjectBuilder struct{}

func InternshipLocation() InternshipLocationObjectBuilder {
	return &internshipLocationQueryObjectBuilder{}
}

func (i *internshipLocationQueryObjectBuilder) FromMap(data map[string]interface{}) (internshipLocation.InternshipLocation, errors.Error) {
	id, err := uuid.Parse(string(data[query.InternshipLocationID].([]uint8)))
	if err != nil {
		logger.Error().Msg(err.Error())
		return nil, errors.NewUnexpected()
	}
	name := fmt.Sprint(data[query.InternshipLocationName])
	address := fmt.Sprint(data[query.InternshipLocationAddress])
	city := fmt.Sprint(data[query.InternshipLocationCity])
	nullableLat := utils.GetNullableValue[[]uint8](data[query.InternshipLocationLat])
	nullableLong := utils.GetNullableValue[[]uint8](data[query.InternshipLocationLong])
	var latString string
	var longString string
	var lat *float64
	var long *float64
	if nullableLat != nil {
		for _, item := range *nullableLat {
			latString += string(item)
		}
		aux, err := strconv.ParseFloat(latString, 64)
		if err != nil {
			logger.Error().Msg(err.Error())
			return nil, errors.NewUnexpected()
		}
		lat = &aux
	}
	if nullableLong != nil {
		for _, item := range *nullableLong {
			longString += string(item)
		}
		aux, err := strconv.ParseFloat(longString, 64)
		if err != nil {
			logger.Error().Msg(err.Error())
			return nil, errors.NewUnexpected()
		}
		long = &aux
	}
	_internshipLocation, validationError := internshipLocation.NewBuilder().
		WithID(id).
		WithName(name).
		WithAddress(address).
		WithCity(city).
		WithLat(lat).
		WithLong(long).
		Build()
	if validationError != nil {
		logger.Error().Msg(validationError.String())
		return nil, validationError
	}
	return _internshipLocation, nil
}

func (i *internshipLocationQueryObjectBuilder) FromRows(rows *sqlx.Rows) ([]internshipLocation.InternshipLocation, errors.Error) {
	if rows == nil {
		err := errors.NewFromString("row value cannot be nil")
		logger.Error().Msg(err.String())
		return nil, err
	}
	defer rows.Close()
	internshipLocations := make([]internshipLocation.InternshipLocation, 0)
	for rows.Next() {
		var serializedInternshipLocation = map[string]interface{}{}
		nativeError := rows.MapScan(serializedInternshipLocation)
		if nativeError != nil {
			logger.Error().Msg(nativeError.Error())
			return nil, errors.New(nativeError)
		}
		_internshipLocation, err := i.FromMap(serializedInternshipLocation)
		if err != nil {
			logger.Error().Msg(err.String())
			return nil, err
		}
		internshipLocations = append(internshipLocations, _internshipLocation)
	}
	return internshipLocations, nil
}
