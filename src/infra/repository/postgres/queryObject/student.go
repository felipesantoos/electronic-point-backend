package queryObject

import (
	"eletronic_point/src/core/domain/campus"
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/institution"
	"eletronic_point/src/core/domain/internship"
	"eletronic_point/src/core/domain/internshipLocation"
	"eletronic_point/src/core/domain/person"
	"eletronic_point/src/core/domain/student"
	"eletronic_point/src/infra/repository/postgres/query"
	"eletronic_point/src/utils"
	"fmt"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type StudentObjectBuilder interface {
	FromMap(data map[string]interface{}) (student.Student, errors.Error)
	FromRows(rows *sqlx.Rows) ([]student.Student, errors.Error)
}

type studentQueryObjectBuilder struct{}

func Student() StudentObjectBuilder {
	return &studentQueryObjectBuilder{}
}

func (s *studentQueryObjectBuilder) FromMap(data map[string]interface{}) (student.Student, errors.Error) {
	id, err := uuid.Parse(string(data[query.PersonID].([]uint8)))
	if err != nil {
		logger.Error().Msg(err.Error())
		return nil, errors.NewUnexpected()
	}
	name := fmt.Sprint(data[query.PersonName])
	birthDate := fmt.Sprint(data[query.PersonBirthDate])[:10]
	email := fmt.Sprint(data[query.PersonEmail])
	cpf := fmt.Sprint(data[query.PersonCPF])
	phone := fmt.Sprint(data[query.PersonPhone])
	registration := fmt.Sprint(data[query.StudentRegistration])
	profilePicture := utils.GetNullableValue[string](data[query.StudentProfilePicture])
	campusID, err := uuid.Parse(string(data[query.CampusID].([]uint8)))
	if err != nil {
		logger.Error().Msg(err.Error())
		return nil, errors.NewUnexpected()
	}
	campusName := fmt.Sprint(data[query.CampusName])
	_campus, validationError := campus.NewBuilder().WithID(campusID).WithName(campusName).Build()
	if validationError != nil {
		logger.Error().Msg(validationError.String())
		return nil, validationError
	}
	institutionID, err := uuid.Parse(string(data[query.InstitutionID].([]uint8)))
	if err != nil {
		logger.Error().Msg(err.Error())
		return nil, errors.NewUnexpected()
	}
	institutionName := fmt.Sprint(data[query.InstitutionName])
	_institution, validationError := institution.NewBuilder().WithID(institutionID).WithName(institutionName).Build()
	if validationError != nil {
		logger.Error().Msg(validationError.String())
		return nil, validationError
	}
	course := fmt.Sprint(data[query.StudentCourse])
	totalWorkload, err := strconv.Atoi(fmt.Sprint(data[query.StudentTotalWorkload]))
	if err != nil {
		logger.Error().Msg(err.Error())
		return nil, errors.NewUnexpected()
	}
	var locationID *uuid.UUID
	nullableLocationID := utils.GetNullableValue[[]uint8](data[query.InternshipLocationID])
	if nullableLocationID != nil {
		aux, err := uuid.ParseBytes(*nullableLocationID)
		if err != nil {
			logger.Error().Msg(err.Error())
			return nil, errors.NewUnexpected()
		}
		locationID = &aux
	}
	locationName := utils.GetNullableValue[string](data[query.InternshipLocationName])
	locationNumber := utils.GetNullableValue[string](data[query.InternshipLocationNumber])
	locationStreet := utils.GetNullableValue[string](data[query.InternshipLocationStreet])
	locationNeighborhood := utils.GetNullableValue[string](data[query.InternshipLocationNeighborhood])
	locationCity := utils.GetNullableValue[string](data[query.InternshipLocationCity])
	locationZipCode := utils.GetNullableValue[string](data[query.InternshipLocationZipCode])
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
	var internshipID *uuid.UUID
	nullableInternshipID := utils.GetNullableValue[[]uint8](data[query.InternshipID])
	if nullableInternshipID != nil {
		aux, err := uuid.ParseBytes(*nullableInternshipID)
		if err != nil {
			logger.Error().Msg(err.Error())
			return nil, errors.NewUnexpected()
		}
		internshipID = &aux
	}
	internshipStartedIn := utils.GetNullableValue[time.Time](data[query.InternshipStartedIn])
	internshipEndedIn := utils.GetNullableValue[time.Time](data[query.InternshipEndedIn])
	_person, validationError := person.NewBuilder().WithID(id).WithName(name).
		WithBirthDate(birthDate).WithEmail(email).WithCPF(cpf).WithPhone(phone).Build()
	if validationError != nil {
		logger.Error().Msg(validationError.String())
		return nil, validationError
	}
	internshipLocationBuilder := internshipLocation.NewBuilder()
	if locationID != nil {
		internshipLocationBuilder.WithID(*locationID)
	}
	if locationName != nil {
		internshipLocationBuilder.WithName(*locationName)
	}
	if locationNumber != nil {
		internshipLocationBuilder.WithNumber(*locationNumber)
	}
	if locationStreet != nil {
		internshipLocationBuilder.WithStreet(*locationStreet)
	}
	if locationNeighborhood != nil {
		internshipLocationBuilder.WithNeighborhood(*locationNeighborhood)
	}
	if locationCity != nil {
		internshipLocationBuilder.WithCity(*locationCity)
	}
	if locationZipCode != nil {
		internshipLocationBuilder.WithZipCode(*locationZipCode)
	}
	if locationLat != nil {
		internshipLocationBuilder.WithLat(*locationLat)
	}
	if locationLong != nil {
		internshipLocationBuilder.WithLong(*locationLong)
	}
	location, validationError := internshipLocationBuilder.Build()
	if validationError != nil {
		logger.Error().Msg(validationError.String())
		return nil, validationError
	}
	currentInternshipBuilder := internship.NewBuilder()
	if internshipID != nil {
		currentInternshipBuilder.WithID(*internshipID)
		currentInternshipBuilder.WithStudentID(*_person.ID())
	}
	if internshipStartedIn != nil {
		currentInternshipBuilder.WithStartedIn(*internshipStartedIn)
	}
	if internshipEndedIn != nil {
		currentInternshipBuilder.WithEndedIn(internshipEndedIn)
	}
	if locationID != nil {
		currentInternshipBuilder.WithLocation(location)
	}
	currentInternship, validationError := currentInternshipBuilder.Build()
	if validationError != nil {
		logger.Error().Msg(validationError.String())
		return nil, validationError
	}
	studentBuilder := student.NewBuilder()
	studentBuilder.WithPerson(_person).
		WithRegistration(registration).
		WithProfilePicture(profilePicture).
		WithCampus(_campus).
		WithInstitution(_institution).
		WithCourse(course).
		WithTotalWorkload(totalWorkload)
	if internshipStartedIn != nil {
		studentBuilder.WithCurrentInternship(currentInternship)
	}
	_student, validationError := studentBuilder.Build()
	if validationError != nil {
		logger.Error().Msg(validationError.String())
		return nil, validationError
	}
	return _student, nil
}

func (s *studentQueryObjectBuilder) FromRows(rows *sqlx.Rows) ([]student.Student, errors.Error) {
	if rows == nil {
		err := errors.NewFromString("row value cannot be nil")
		logger.Error().Msg(err.String())
		return nil, err
	}
	defer rows.Close()
	students := make([]student.Student, 0)
	for rows.Next() {
		var serializedStudent = map[string]interface{}{}
		nativeError := rows.MapScan(serializedStudent)
		if nativeError != nil {
			logger.Error().Msg(nativeError.Error())
			return nil, errors.New(nativeError)
		}
		_student, err := s.FromMap(serializedStudent)
		if err != nil {
			logger.Error().Msg(err.String())
			return nil, err
		}
		students = append(students, _student)
	}
	return students, nil
}
