package postgres

import (
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/timeRecord"
	"eletronic_point/src/core/domain/timeRecordStatus"
	"eletronic_point/src/core/interfaces/secondary"
	"eletronic_point/src/core/messages"
	"eletronic_point/src/core/services/filters"
	"eletronic_point/src/infra/repository"
	"eletronic_point/src/infra/repository/postgres/query"
	"eletronic_point/src/infra/repository/postgres/queryObject"

	"github.com/google/uuid"
)

type timeRecordRepository struct{}

func NewTimeRecordRepository() secondary.TimeRecordPort {
	return &timeRecordRepository{}
}

func (this timeRecordRepository) Create(_timeRecord timeRecord.TimeRecord) (*uuid.UUID, errors.Error) {
	transaction, err := repository.BeginTransaction()
	if err != nil {
		logger.Error().Msg(err.String())
		return nil, err
	}
	defer transaction.CloseConn()
	id, err := txQueryRowReturningID(transaction, query.TimeRecord().Insert(),
		_timeRecord.Date(),
		_timeRecord.EntryTime(),
		_timeRecord.ExitTime(),
		_timeRecord.Location(),
		_timeRecord.IsOffSite(),
		_timeRecord.Justification(),
		_timeRecord.StudentID(),
		_timeRecord.InternshipID(),
	)
	if err != nil {
		logger.Error().Msg(err.String())
		return nil, err
	}
	timeRecordID, convErr := uuid.Parse(id)
	if convErr != nil {
		logger.Error().Msg(convErr.Error())
		return nil, errors.NewUnexpected()
	}
	_, movementErr := txQueryRowReturningID(transaction, query.TimeRecordStatusMovement().Insert(),
		timeRecordID,
		timeRecordStatus.Pending.ID(),
		_timeRecord.StudentID(),
		nil,
	)
	if movementErr != nil {
		logger.Error().Msg(movementErr.String())
		return nil, movementErr
	}
	err = transaction.Commit()
	if err != nil {
		logger.Error().Msg(err.String())
		return nil, err
	}
	return &timeRecordID, nil
}

func (this timeRecordRepository) Update(_timeRecord timeRecord.TimeRecord) errors.Error {
	_, err := repository.ExecQuery(query.TimeRecord().Update(), _timeRecord.ID(),
		_timeRecord.Date(), _timeRecord.EntryTime(), _timeRecord.ExitTime(), _timeRecord.Location(),
		_timeRecord.IsOffSite(), _timeRecord.Justification(), _timeRecord.StudentID(), _timeRecord.InternshipID(),
	)
	if err != nil {
		logger.Error().Msg(err.String())
		return err
	}
	return nil
}

func (this timeRecordRepository) Delete(id uuid.UUID) errors.Error {
	_, err := repository.ExecQuery(query.TimeRecord().Delete(), id)
	if err != nil {
		logger.Error().Msg(err.String())
		return err
	}
	return nil
}

func (this timeRecordRepository) List(_filters filters.TimeRecordFilters) ([]timeRecord.TimeRecord, errors.Error) {
	search := ""
	if _filters.Search != nil {
		search = *_filters.Search
	}

	rows, err := repository.Queryx(query.TimeRecord().Select().All(), _filters.StudentID,
		_filters.StartDate, _filters.EndDate, _filters.TeacherID, _filters.StatusID, search)
	if err != nil {
		logger.Error().Msg(err.String())
		return nil, err
	}
	defer rows.Close()
	timeRecords, err := queryObject.TimeRecord().FromRows(rows)
	if err != nil {
		logger.Error().Msg(err.String())
		return nil, err
	}
	return timeRecords, nil
}

func (this timeRecordRepository) Get(id uuid.UUID, _filters filters.TimeRecordFilters) (timeRecord.TimeRecord, errors.Error) {
	rows, err := repository.Queryx(query.TimeRecord().Select().ByID(), id, _filters.StudentID, _filters.TeacherID)
	if err != nil {
		logger.Error().Msg(err.String())
		return nil, err
	}
	defer rows.Close()
	if !rows.Next() {
		return nil, errors.NewFromString(messages.TimeRecordNotFoundErrorMessage)
	}
	var serializedTimeRecord = map[string]interface{}{}
	nativeError := rows.MapScan(serializedTimeRecord)
	if nativeError != nil {
		logger.Error().Msg(nativeError.Error())
		return nil, errors.NewInternal(nativeError)
	}
	_timeRecord, err := queryObject.TimeRecord().FromMap(serializedTimeRecord)
	if err != nil {
		logger.Error().Msg(err.String())
		return nil, err
	}
	return _timeRecord, nil
}

func (this timeRecordRepository) UpdateStatus(timeRecordID uuid.UUID, updatedBy uuid.UUID, statusID uuid.UUID) errors.Error {
	transaction, err := repository.BeginTransaction()
	if err != nil {
		logger.Error().Msg(err.String())
		return err
	}
	defer transaction.CloseConn()
	terminateErr := defaultTxExecQuery(transaction, query.TimeRecordStatusMovement().Terminate(), timeRecordID)
	if terminateErr != nil {
		logger.Error().Msg(terminateErr.String())
		return terminateErr
	}
	_, insertErr := txQueryRowReturningID(transaction, query.TimeRecordStatusMovement().Insert(),
		timeRecordID,
		statusID,
		updatedBy,
		nil,
	)
	if insertErr != nil {
		logger.Error().Msg(insertErr.String())
		return insertErr
	}
	err = transaction.Commit()
	if err != nil {
		logger.Error().Msg(err.String())
		return err
	}
	return nil
}
