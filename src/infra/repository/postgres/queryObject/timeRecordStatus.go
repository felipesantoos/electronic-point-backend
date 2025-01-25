package queryObject

import (
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/timeRecordStatus"
	"eletronic_point/src/infra/repository/postgres/query"
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type TimeRecordStatusObjectBuilder interface {
	FromMap(data map[string]interface{}) (timeRecordStatus.TimeRecordStatus, errors.Error)
	FromRows(rows *sqlx.Rows) ([]timeRecordStatus.TimeRecordStatus, errors.Error)
}

type timeRecordStatusQueryObjectBuilder struct{}

func TimeRecordStatus() TimeRecordStatusObjectBuilder {
	return &timeRecordStatusQueryObjectBuilder{}
}

func (t *timeRecordStatusQueryObjectBuilder) FromMap(data map[string]interface{}) (timeRecordStatus.TimeRecordStatus, errors.Error) {
	id, err := uuid.Parse(string(data[query.TimeRecordStatusID].([]uint8)))
	if err != nil {
		logger.Error().Msg(err.Error())
		return nil, errors.NewUnexpected()
	}

	name := fmt.Sprint(data[query.TimeRecordStatusName])
	_timeRecordStatus, validationError := timeRecordStatus.NewBuilder().WithID(id).WithName(name).Build()
	if validationError != nil {
		logger.Error().Msg(validationError.String())
		return nil, validationError
	}
	return _timeRecordStatus, nil
}

func (t *timeRecordStatusQueryObjectBuilder) FromRows(rows *sqlx.Rows) ([]timeRecordStatus.TimeRecordStatus, errors.Error) {
	if rows == nil {
		err := errors.NewFromString("row value cannot be nil")
		logger.Error().Msg(err.String())
		return nil, err
	}
	defer rows.Close()
	timeRecordStatuses := make([]timeRecordStatus.TimeRecordStatus, 0)
	for rows.Next() {
		var serializedTimeRecordStatus = map[string]interface{}{}
		nativeError := rows.MapScan(serializedTimeRecordStatus)
		if nativeError != nil {
			logger.Error().Msg(nativeError.Error())
			return nil, errors.New(nativeError)
		}
		_timeRecordStatus, err := t.FromMap(serializedTimeRecordStatus)
		if err != nil {
			logger.Error().Msg(err.String())
			return nil, err
		}
		timeRecordStatuses = append(timeRecordStatuses, _timeRecordStatus)
	}
	return timeRecordStatuses, nil
}
