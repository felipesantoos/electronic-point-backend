package queryObject

import (
	"eletronic_point/src/core/domain/errors"
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
	dateString := fmt.Sprint(data[query.TimeRecordDate])
	date, err := time.Parse(time.RFC3339, dateString)
	if err != nil {
		logger.Error().Msg(err.Error())
		return nil, errors.NewUnexpected()
	}
	entryTimeString := fmt.Sprint(data[query.TimeRecordEntryTime])
	entryTime, err := time.Parse(time.RFC3339, entryTimeString)
	if err != nil {
		logger.Error().Msg(err.Error())
		return nil, errors.NewUnexpected()
	}
	exitTimeString := fmt.Sprint(data[query.TimeRecordExitTime])
	exitTime, err := time.Parse(time.RFC3339, exitTimeString)
	if err != nil {
		logger.Error().Msg(err.Error())
		return nil, errors.NewUnexpected()
	}
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
	_timeRecord, validationError := timeRecord.NewBuilder().
		WithID(id).
		WithDate(date).
		WithEntryTime(entryTime).
		WithExitTime(&exitTime).
		WithLocation(location).
		WithIsOffSite(isOffSite).
		WithJustification(justification).
		WithStudentID(studentID).
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
