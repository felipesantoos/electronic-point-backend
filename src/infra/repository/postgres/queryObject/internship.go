package queryObject

import (
	"eletronic_point/src/core/domain/campus"
	"eletronic_point/src/core/domain/course"
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/institution"
	"eletronic_point/src/core/domain/internship"
	"eletronic_point/src/core/domain/internshipLocation"
	"eletronic_point/src/core/domain/simplifiedStudent"
	"eletronic_point/src/infra/repository/postgres/query"
	"eletronic_point/src/utils"
	"fmt"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type InternshipObjectBuilder interface {
	FromMap(data map[string]interface{}, shouldReturnStudentInfo bool) (internship.Internship, errors.Error)
	FromRows(rows *sqlx.Rows, shouldReturnStudentInfo bool) ([]internship.Internship, errors.Error)
}

type internshipQueryObjectBuilder struct{}

func Internship() InternshipObjectBuilder {
	return &internshipQueryObjectBuilder{}
}

func (i *internshipQueryObjectBuilder) FromMap(data map[string]interface{}, shouldReturnStudentInfo bool) (internship.Internship, errors.Error) {
	internshipID, err := uuid.Parse(string(data[query.InternshipID].([]uint8)))
	if err != nil {
		logger.Error().Msg(err.Error())
		return nil, errors.NewUnexpected()
	}
	layout := "2006-01-02 15:04:05 -0700 -0700"
	internshipStartedInString := fmt.Sprint(data[query.InternshipStartedIn])
	internshipStartedIn, err := time.Parse(layout, internshipStartedInString)
	if err != nil {
		logger.Error().Msg(err.Error())
		return nil, errors.NewUnexpected()
	}
	internshipEndedIn := utils.GetNullableValue[time.Time](data[query.InternshipEndedIn])
	scheduleEntryTime := utils.GetNullableValue[time.Time](data[query.InternshipScheduleEntryTime])
	scheduleExitTime := utils.GetNullableValue[time.Time](data[query.InternshipScheduleExitTime])
	locationID, err := uuid.Parse(string(data[query.InternshipLocationID].([]uint8)))
	if err != nil {
		logger.Error().Msg(err.Error())
		return nil, errors.NewUnexpected()
	}
	locationName := fmt.Sprint(data[query.InternshipLocationName])
	locationNumber := fmt.Sprint(data[query.InternshipLocationNumber])
	locationStreet := fmt.Sprint(data[query.InternshipLocationStreet])
	locationNeighborhood := fmt.Sprint(data[query.InternshipLocationNeighborhood])
	locationCity := fmt.Sprint(data[query.InternshipLocationCity])
	locationZipCode := fmt.Sprint(data[query.InternshipLocationZipCode])
	var locationLat float64
	var locationLong float64
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
		locationLat = aux
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
		locationLong = aux
	}
	studentID, err := uuid.Parse(string(data[query.PersonID].([]uint8)))
	if err != nil {
		logger.Error().Msg(err.Error())
		return nil, errors.NewUnexpected()
	}
	studentName := fmt.Sprint(data[query.PersonName])
	studentProfilePicture := utils.GetNullableValue[string](data[query.StudentProfilePicture])
	studentTotalWorkload := int(data[query.StudentTotalWorkload].(int64))
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
	courseID, err := uuid.Parse(string(data[query.CourseID].([]uint8)))
	if err != nil {
		logger.Error().Msg(err.Error())
		return nil, errors.NewUnexpected()
	}
	courseName := fmt.Sprint(data[query.CourseName])
	_course, validationError := course.NewBuilder().WithID(courseID).WithName(courseName).Build()
	if validationError != nil {
		logger.Error().Msg(validationError.String())
		return nil, validationError
	}
	studentBuilder := simplifiedStudent.NewBuilder().WithID(studentID).WithName(studentName).
		WithProfilePicture(studentProfilePicture).WithInstitution(_institution).WithCampus(_campus).WithCourse(_course).
		WithTotalWorkload(studentTotalWorkload)
	_student, validationError := studentBuilder.Build()
	if validationError != nil {
		logger.Error().Msg(validationError.String())
		return nil, validationError
	}
	location, validationError := internshipLocation.NewBuilder().
		WithID(locationID).
		WithName(locationName).
		WithNumber(locationNumber).
		WithStreet(locationStreet).
		WithNeighborhood(locationNeighborhood).
		WithCity(locationCity).
		WithZipCode(locationZipCode).
		WithLat(locationLat).
		WithLong(locationLong).
		Build()
	if validationError != nil {
		logger.Error().Msg(validationError.String())
		return nil, validationError
	}
	internshipBuilder := internship.NewBuilder().
		WithID(internshipID).
		WithStartedIn(internshipStartedIn).
		WithEndedIn(internshipEndedIn).
		WithScheduleEntryTime(scheduleEntryTime).
		WithScheduleExitTime(scheduleExitTime).
		WithLocation(location)
	if shouldReturnStudentInfo {
		internshipBuilder.WithStudent(_student)
	}
	_internship, validationError := internshipBuilder.Build()
	if validationError != nil {
		logger.Error().Msg(validationError.String())
		return nil, validationError
	}
	return _internship, nil
}

func (s *internshipQueryObjectBuilder) FromRows(rows *sqlx.Rows, shouldReturnStudentInfo bool) ([]internship.Internship, errors.Error) {
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
		_internship, err := s.FromMap(serializedInternship, shouldReturnStudentInfo)
		if err != nil {
			logger.Error().Msg(err.String())
			return nil, err
		}
		internships = append(internships, _internship)
	}
	return internships, nil
}
