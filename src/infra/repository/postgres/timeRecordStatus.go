package postgres

import (
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/timeRecordStatus"
	"eletronic_point/src/core/interfaces/secondary"
	"eletronic_point/src/core/messages"
	"eletronic_point/src/infra/repository"
	"eletronic_point/src/infra/repository/postgres/query"
	"eletronic_point/src/infra/repository/postgres/queryObject"

	"github.com/google/uuid"
)

type timeRecordStatusRepository struct{}

func NewTimeRecordStatusRepository() secondary.TimeRecordStatusPort {
	return &timeRecordStatusRepository{}
}

func (this timeRecordStatusRepository) List() ([]timeRecordStatus.TimeRecordStatus, errors.Error) {
	rows, err := repository.Queryx(query.TimeRecordStatus().Select().All())
	if err != nil {
		logger.Error().Msg(err.String())
		return nil, errors.NewUnexpected()
	}
	defer rows.Close()
	timeRecordStatuses, err := queryObject.TimeRecordStatus().FromRows(rows)
	if err != nil {
		logger.Error().Msg(err.String())
		return nil, errors.NewUnexpected()
	}
	return timeRecordStatuses, nil
}

func (this timeRecordStatusRepository) Get(id uuid.UUID) (timeRecordStatus.TimeRecordStatus, errors.Error) {
	rows, err := repository.Queryx(query.TimeRecordStatus().Select().ByID(), id)
	if err != nil {
		logger.Error().Msg(err.String())
		return nil, errors.NewUnexpected()
	}
	defer rows.Close()
	if !rows.Next() {
		return nil, errors.NewFromString(messages.TimeRecordStatusNotFoundErrorMessage)
	}
	var serializedTimeRecordStatus = map[string]interface{}{}
	nativeError := rows.MapScan(serializedTimeRecordStatus)
	if nativeError != nil {
		logger.Error().Msg(nativeError.Error())
		return nil, errors.NewUnexpected()
	}
	_timeRecordStatus, err := queryObject.TimeRecordStatus().FromMap(serializedTimeRecordStatus)
	if err != nil {
		logger.Error().Msg(err.String())
		return nil, errors.NewUnexpected()
	}
	return _timeRecordStatus, nil
}

func (this timeRecordStatusRepository) Create(data timeRecordStatus.TimeRecordStatus) (*uuid.UUID, errors.Error) {
	return execQueryReturningID(query.TimeRecordStatus().Insert(), data.Name())
}

func (this timeRecordStatusRepository) Update(data timeRecordStatus.TimeRecordStatus) errors.Error {
	return defaultExecQuery(query.TimeRecordStatus().Update(), data.ID(), data.Name())
}

func (this timeRecordStatusRepository) Delete(id uuid.UUID) errors.Error {
	return defaultExecQuery(query.TimeRecordStatus().Delete(), id)
}
