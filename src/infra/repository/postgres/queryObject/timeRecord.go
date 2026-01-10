package queryObject

import (
	"eletronic_point/src/core/domain/campus"
	"eletronic_point/src/core/domain/course"
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/institution"
	"eletronic_point/src/core/domain/internship"
	"eletronic_point/src/core/domain/internshipLocation"
	"eletronic_point/src/core/domain/simplifiedStudent"
	"eletronic_point/src/core/domain/timeRecord"
	"eletronic_point/src/infra/repository/postgres/query"
	"eletronic_point/src/utils"
	"fmt"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type TimeRecordObjectBuilder interface {
	FromMap(data map[string]interface{}) (timeRecord.TimeRecord, errors.Error)
	FromRows(rows *sqlx.Rows) ([]timeRecord.TimeRecord, errors.Error)
}

type timeRecordQueryObjectBuilder struct{}

func TimeRecord() TimeRecordObjectBuilder {
	return &timeRecordQueryObjectBuilder{}
}

func (t *timeRecordQueryObjectBuilder) FromMap(data map[string]interface{}) (timeRecord.TimeRecord, errors.Error) {
	id, err := uuid.Parse(string(data[query.TimeRecordID].([]uint8)))
	if err != nil {
		logger.Error().Msg(err.Error())
		return nil, errors.NewUnexpected()
	}
	layout := "2006-01-02 15:04:05 -0700 -0700"
	dateString := fmt.Sprint(data[query.TimeRecordDate])
	date, err := time.Parse(layout, dateString)
	if err != nil {
		logger.Error().Msg(err.Error())
		return nil, errors.NewUnexpected()
	}
	entryTimeString := fmt.Sprint(data[query.TimeRecordEntryTime])
	entryTime, err := time.Parse(layout, entryTimeString)
	if err != nil {
		logger.Error().Msg(err.Error())
		return nil, errors.NewUnexpected()
	}
	exitTime := utils.GetNullableValue[time.Time](data[query.TimeRecordExitTime])
	location := fmt.Sprint(data[query.TimeRecordLocation])
	isOffSite, err := strconv.ParseBool(fmt.Sprint(data[query.TimeRecordIsOffSite]))
	if err != nil {
		logger.Error().Msg(err.Error())
		return nil, errors.NewUnexpected()
	}
	justification := utils.GetNullableValue[string](data[query.TimeRecordJustification])
	studentID, err := uuid.Parse(string(data[query.TimeRecordStudentID].([]uint8)))
	if err != nil {
		logger.Error().Msg(err.Error())
		return nil, errors.NewUnexpected()
	}
	internshipID, err := uuid.Parse(string(data[query.TimeRecordInternshipID].([]uint8)))
	if err != nil {
		logger.Error().Msg(err.Error())
		return nil, errors.NewUnexpected()
	}
	_timeRecordStatus, validationError := TimeRecordStatus().FromMap(data)
	if validationError != nil {
		logger.Error().Msg(validationError.String())
		return nil, validationError
	}

	// Build SimplifiedStudent
	studentName := fmt.Sprint(data[query.PersonName])
	studentProfilePicture := utils.GetNullableValue[string](data[query.StudentProfilePicture])
	institutionID, _ := uuid.Parse(string(data[query.InstitutionID].([]uint8)))
	institutionName := fmt.Sprint(data[query.InstitutionName])
	_institution, _ := institution.NewBuilder().WithID(institutionID).WithName(institutionName).Build()
	campusID, _ := uuid.Parse(string(data[query.CampusID].([]uint8)))
	campusName := fmt.Sprint(data[query.CampusName])
	_campus, _ := campus.NewBuilder().WithID(campusID).WithName(campusName).Build()
	courseID, _ := uuid.Parse(string(data[query.CourseID].([]uint8)))
	courseName := fmt.Sprint(data[query.CourseName])
	_course, _ := course.NewBuilder().WithID(courseID).WithName(courseName).Build()
	_student, _ := simplifiedStudent.NewBuilder().WithID(studentID).WithName(studentName).
		WithProfilePicture(studentProfilePicture).WithInstitution(_institution).WithCampus(_campus).WithCourse(_course).Build()

	// Build Internship
	var _internship internship.Internship
	if data[query.InternshipID] != nil {
		internshipID, _ := uuid.Parse(string(data[query.InternshipID].([]uint8)))
		layout := "2006-01-02 15:04:05 -0700 -0700"
		startedInString := fmt.Sprint(data[query.InternshipStartedIn])
		startedIn, _ := time.Parse(layout, startedInString)
		endedIn := utils.GetNullableValue[time.Time](data[query.InternshipEndedIn])
		schedEntry := utils.GetNullableValue[time.Time](data[query.InternshipScheduleEntryTime])
		schedExit := utils.GetNullableValue[time.Time](data[query.InternshipScheduleExitTime])
		locID, _ := uuid.Parse(string(data[query.InternshipLocationID].([]uint8)))
		locName := fmt.Sprint(data[query.InternshipLocationName])
		location, _ := internshipLocation.NewBuilder().WithID(locID).WithName(locName).Build()
		_internship, _ = internship.NewBuilder().WithID(internshipID).WithStartedIn(startedIn).WithEndedIn(endedIn).
			WithScheduleEntryTime(schedEntry).WithScheduleExitTime(schedExit).WithLocation(location).Build()
	}

	_timeRecord, validationError := timeRecord.NewBuilder().
		WithID(id).
		WithDate(date).
		WithEntryTime(entryTime).
		WithExitTime(exitTime).
		WithLocation(location).
		WithIsOffSite(isOffSite).
		WithJustification(justification).
		WithStudentID(studentID).
		WithInternshipID(internshipID).
		WithStudent(_student).
		WithInternship(_internship).
		WithTimeRecord(_timeRecordStatus).
		Build()
	if validationError != nil {
		logger.Error().Msg(validationError.String())
		return nil, validationError
	}
	return _timeRecord, nil
}

func (t *timeRecordQueryObjectBuilder) FromRows(rows *sqlx.Rows) ([]timeRecord.TimeRecord, errors.Error) {
	if rows == nil {
		err := errors.NewFromString("row value cannot be nil")
		logger.Error().Msg(err.String())
		return nil, err
	}
	defer rows.Close()
	timeRecords := make([]timeRecord.TimeRecord, 0)
	for rows.Next() {
		var serializedTimeRecord = map[string]interface{}{}
		nativeError := rows.MapScan(serializedTimeRecord)
		if nativeError != nil {
			logger.Error().Msg(nativeError.Error())
			return nil, errors.New(nativeError)
		}
		_timeRecord, err := t.FromMap(serializedTimeRecord)
		if err != nil {
			logger.Error().Msg(err.String())
			return nil, err
		}
		timeRecords = append(timeRecords, _timeRecord)
	}
	return timeRecords, nil
}
