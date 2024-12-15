package queryObject

import (
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/internship"
	"eletronic_point/src/core/domain/internshipLocation"
	"eletronic_point/src/infra/repository/postgres/query"
	"eletronic_point/src/utils"
	"fmt"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type InternshipObjectBuilder interface {
	FromMap(data map[string]interface{}) (internship.Internship, errors.Error)
	FromRows(rows *sqlx.Rows) ([]internship.Internship, errors.Error)
}

type internshipQueryObjectBuilder struct{}

func Internship() InternshipObjectBuilder {
	return &internshipQueryObjectBuilder{}
}

func (i *internshipQueryObjectBuilder) FromMap(data map[string]interface{}) (internship.Internship, errors.Error) {
	locationID, err := uuid.Parse(string(data[query.InternshipLocationID].([]uint8)))
	if err != nil {
		logger.Error().Msg(err.Error())
		return nil, errors.NewUnexpected()
	}
	locationName := fmt.Sprint(data[query.InternshipLocationName])
	locationAddress := fmt.Sprint(data[query.InternshipLocationAddress])
	locationCity := fmt.Sprint(data[query.InternshipLocationCity])
	var locationLat *float64
	var locationLong *float64
	nullableLocationLat := utils.GetNullableValue[[]uint8](data[query.InternshipLocationLat])
	nullableLocationLong := utils.GetNullableValue[[]uint8](data[query.InternshipLocationLong])
	var locationLatString string
	var locationLongString string
	if nullableLocationLat != nil {
		for _, item := range *nullableLocationLat {
			locationLatString += string(item)
		}
		aux, err := strconv.ParseFloat(locationLatString, 64)
		if err != nil {
			logger.Error().Msg(err.Error())
			return nil, errors.NewUnexpected()
		}
		locationLat = &aux
	}
	if nullableLocationLong != nil {
		for _, item := range *nullableLocationLong {
			locationLongString += string(item)
		}
		aux, err := strconv.ParseFloat(locationLongString, 64)
		if err != nil {
			logger.Error().Msg(err.Error())
			return nil, errors.NewUnexpected()
		}
		locationLong = &aux
	}
	layout := "2006-01-02 15:04:05 -0700 -0700"
	internshipStartedInString := fmt.Sprint(data[query.StudentWorksAtInternshipLocationStartedIn])
	internshipStartedIn, err := time.Parse(layout, internshipStartedInString)
	if err != nil {
		logger.Error().Msg(err.Error())
		return nil, errors.NewUnexpected()
	}
	var internshipEndedIn *time.Time
	nullableInternshipEndedInString := utils.GetNullableValue[time.Time](data[query.StudentWorksAtInternshipLocationEndedIn])
	if nullableInternshipEndedInString != nil {
		internshipEndedIn = nullableInternshipEndedInString
	}
	location, validationError := internshipLocation.NewBuilder().
		WithID(locationID).
		WithName(locationName).
		WithAddress(locationAddress).
		WithCity(locationCity).
		WithLat(locationLat).
		WithLong(locationLong).
		Build()
	if validationError != nil {
		logger.Error().Msg(validationError.String())
		return nil, validationError
	}
	_internship, validationError := internship.NewBuilder().
		WithStartedIn(internshipStartedIn).
		WithEndedIn(internshipEndedIn).
		WithLocation(location).
		Build()
	if validationError != nil {
		logger.Error().Msg(validationError.String())
		return nil, validationError
	}
	return _internship, nil
}

func (s *internshipQueryObjectBuilder) FromRows(rows *sqlx.Rows) ([]internship.Internship, errors.Error) {
	if rows == nil {
		err := errors.NewFromString("row value cannot be nil")
		logger.Error().Msg(err.String())
		return nil, err
	}
	defer rows.Close()
	internships := make([]internship.Internship, 0)
	for rows.Next() {
		var serializedInternship = map[string]interface{}{}
		nativeError := rows.MapScan(serializedInternship)
		if nativeError != nil {
			logger.Error().Msg(nativeError.Error())
			return nil, errors.New(nativeError)
		}
		_internship, err := s.FromMap(serializedInternship)
		if err != nil {
			logger.Error().Msg(err.String())
			return nil, err
		}
		internships = append(internships, _internship)
	}
	return internships, nil
}
